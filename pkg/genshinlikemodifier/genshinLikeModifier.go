package genshinlikemodifier

import (
	"github.com/genshinsim/gcsim/pkg/core/glog"
)

type Mod interface {
	Key() string
	Expiry() int
	Event() glog.Event
	SetEvent(glog.Event)
	AffectedByHitlag() bool
	Extend(string, glog.Logger, int, int)
}

// CONSTANTS //

// Matches genshin mod types. May not be exhaustive, taken from discord dm
// Thanks to @Shizuka
// for any given key, call it blah:
// 
// Refresh - only 1, can be refreshed (basically what gcsim mods is)
// Unique - 1 one at a time, does nothing on add/refresh (dur is icd)
// Multiple - can stack, can't refresh when at max (oldest needs to expire before it can stack again)
// MultipleRefresh - can stack, refresh oldest when at max
// MultipleRefreshNoRemove - same as MultipleRefresh but doesn't trigger onAdded/onRemoved (we don't really care about this)
// MultipleAllRefresh - can stack, always refresh all
type ModType int
const (
	Refresh = ModType(iota) ; 
	Unique
	Multiple
	MultipleRefresh
	MultipleRefreshNoRemove
	MultipleAllRefresh
)

// Stack timer types
const (
	timerIndependent = iota	// Each stack has its own independent timer
	timerShared	// All stacks share a timer
)

// Stack behavior on expire types
const (
	expAll = iota // Expire mod when stack expires
	expSingle	// Reduce payload by 1 when stack expires. TODO: maybe some stack reduces by more than 1? To be validated
)

// Stack refresh types
const (
	refreshOldest = iota  // Can refresh the oldest available stack, if max stacks reached
	refreshBlockable	// Need to wait on open stack before refresh
)

// TODO: Don't use a map here, use an array for performance purposes. 
// With compiler opts, maybe doesn't matter?
var modClassMapper = map[ModType][...]ModType {
	Refresh: [...]ModType{timerShared, expAll, refreshOldest}, 
	Unique: [...]ModType{timerShared, expAll, refreshBlockable},
	Multiple: [...]ModType{timerIndependent, expAll, refreshBlockable},
	MultipleRefresh: [...]ModType{timerIndependent, expAll, refreshOldest},
	MultipleRefreshNoRemove: [...]ModType{timerIndependent, expAll, refreshOldest},
	MultipleAllRefresh: [...]ModType{timerShared, expAll, refreshOldest},
}

// DATA TYPES //

type stackHolder struct {
	stackExpiry int
	event glog.Event
}

//Use int types when possible, floats otherwise
type Payload interface {
	int | float
}
type StackedBase struct {
	// Initialized by caller
	ModKey    string
	Dur       int
	Hitlag    bool
	maxStacks int
	modType ModType
	icd int
	
	// Set in init
	timerType ModType
	expireType ModType
	refreshType ModType

	// Changes dynamically
	ModExpiry int // Refers to when all stacks will fall off and entire mod expires
	extension int //Todo: assume extension applies to all stacks equally; might need to be done on a per-stack basis.
	event     glog.Event
	stacks []stackHolder
	stackIdx int // At any given point in time, refer to the oldest stack.
	payload Payload //todo: may have to be float. Unsure if furina fanfare keeps track of incremental hp changes or float hp changes
	icdExpiry int
}

//// METHODS ////
func (t *StackedBase) Key() string             { return t.ModKey }
func (t *StackedBase) Expiry() int             { 
	exp_frame_adjustor := 0

	//If only one stack is dropped on mod expiry, then we had duration * stacks remaining after expiry to 
	//To find the expected expiry of the whole mod
	if t.expireType == expSingle {exp_frame_adjustor = (t.payload - 1) * t.Dur}
	if exp_frame_adjustor < 0 {exp_frame_adjustor = 0}
	return t.ModExpiry + t.extension + exp_frame_adjustor
}
func (t *StackedBase) Event() glog.Event       { return t.event }
func (t *StackedBase) SetEvent(evt glog.Event) { t.event = evt }
func (t *StackedBase) AffectedByHitlag() bool  { return t.Hitlag }

func (t *StackedBase) Extend(key string, logger glog.Logger, index, amt int) {
	t.extension += amt
	if t.extension < 0 {
		t.extension = 0
	}
	t.event.SetEnded(t.Expiry())
	logger.NewEvent("mod extended", glog.LogStatusEvent, index).
		Write("key", key).
		Write("amt", amt).
		Write("expiry", t.Expiry()).
		Write("ext", t.extension)
}

func (t *StackedBase) SetExpiry(f int) {
	if t.Dur == -1 {
		t.ModExpiry = -1
	} else {
		t.ModExpiry = f + t.Dur
		if t.timerType == timerIndependent  {
			// TODO: Figure out how to throw an error in go
			// Should not directly set expiry for an indepndently-timed stack
			// Or maybe it's ok? Maybe it can be a way to force a stack to expire immediately?
		}
	}
}

//// INITIALIZERS ////
// Creates a new mod for a character
// key: Unique identifier for a type of mod
// dur: How many frames a mod should last after init, assuming no hitlag
// hitlag: Whether or not mod is affected by hitlag. True: Affected by hitlag. False: Ignores hitlag
// modType: One of set of possible ModTypes. Logic internally calculates timer, expire, and refresh types
// maxStacks: Max count of stacks. Size of stack array if independent duration. Max size of payload if shared duration
// icd: Delay in frames before a stack can be refreshed. 
//   If refreshType is refreshBlockable, is the delay from expiry to when the next stack can be added
//   If refreshType is refreshOldest, is the delay from when the previous stack was refreshed to the next
// Defaults if using overloaded methods: hitlag=false, modType=Refresh, maxStacks=0, icd=0
func NewStackedBase(
	key string, 
	dur int,  

	hitlag bool,
	modType ModType,
	maxStacks int, 
	icd int,
	) StackedBase {
		return StackedBase{
			// Initialized by caller
			ModKey: key,
			Dur:    dur,
			Hitlag: hitlag,
			maxStacks: maxStacks,
			modType: modType,
			maxStacks: maxStacks,
			icd: icd,

			// Calculated
			timerType: modClassMapper[modType][0],
			expireType: modClassMapper[modType][1],
			refreshType: modClassMapper[modType][2],
		}
}

// Same signitures as old mod for compatibility
func NewBase(key string, dur int) StackedBase {
	return NewBase(key, dur, false, Refresh, 0, 0)
}

func NewBaseWithHitlag(key string, dur int) StackedBase {
	return NewBase(key, dur, true, Refresh, 0, 0)
}


//// Helpful mod operator functions ////
// Delete removes a modifier. Returns the last detected mod of a given key if at least one mod deleted, or the nil Mod otherwise
func Delete[K Mod](slice *[]K, key string) Mod {
	n := 0
	var m Mod
	for i, v := range *slice {
		if v.Key() == key {
			m = (*slice)[i]
		} else {
			(*slice)[n] = v
			n++
		}
	}
	*slice = (*slice)[:n]
	return m
}

// Add adds a modifier. Returns true if overwritten and the original evt (if overwritten)
// TODO: consider adding a map here to track the index to assist with faster lookups
func Add[K Mod](slice *[]K, mod K, f int) (bool, glog.Event) {
	ind := Find(slice, mod.Key())
	overwrote := false
	var evt glog.Event

	// if does not exist, make new and add
	if ind == -1 {
		*slice = append(*slice, mod)
		return overwrote, evt
	}

	//TODO: come back to this. Why doesn't this fail if an extend was added to the mod?
	// otherwise check not expired
	if (*slice)[ind].Expiry() > f || (*slice)[ind].Expiry() == -1 {
		overwrote = true
		evt = (*slice)[ind].Event()
	}

	if (*slice)[ind]
	(*slice)[ind] = mod

	return overwrote, evt
}

func Find[K Mod](slice *[]K, key string) int {
	ind := -1
	for i, v := range *slice {
		if v.Key() == key {
			ind = i
		}
	}
	return ind
}

func FindCheckExpiry[K Mod](slice *[]K, key string, f int) (int, bool) {
	ind := Find(slice, key)
	if ind == -1 {
		return ind, false
	}
	if (*slice)[ind].Expiry() < f && (*slice)[ind].Expiry() > -1 {
		return ind, false
	}
	return ind, true
}

// LogAdd is a helper that logs mod add events
func LogAdd[K Mod](prefix string, index int, mod K, logger glog.Logger, overwrote bool, oldEvt glog.Event) {
	var evt glog.Event
	if overwrote {
		logger.NewEventBuildMsg(
			glog.LogStatusEvent, index,
			prefix, " mod refreshed",
		).Write(
			"overwrite", true,
		).Write(
			"key", mod.Key(),
		).Write(
			"expiry", mod.Expiry(),
		)
		evt = oldEvt
	} else {
		evt = logger.NewEventBuildMsg(
			glog.LogStatusEvent, index,
			prefix, " mod added",
		).Write(
			"overwrite", false,
		).Write(
			"key", mod.Key(),
		).Write(
			"expiry", mod.Expiry(),
		)
	}
	evt.SetEnded(mod.Expiry())
	mod.SetEvent(evt)
}
