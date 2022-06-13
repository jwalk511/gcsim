package zhongli

import (
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/construct"
	"github.com/genshinsim/gcsim/pkg/core/glog"
)

func (c *char) newStele(dur int, max int) {
	//deal damage when created
	ai := combat.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "Stone Stele (Initial)",
		AttackTag:  combat.AttackTagElementalArt,
		ICDTag:     combat.ICDTagElementalArt,
		ICDGroup:   combat.ICDGroupDefault,
		StrikeType: combat.StrikeTypeBlunt,
		Element:    attributes.Geo,
		Durability: 50,
		Mult:       skill[c.TalentLvlSkill()],
		FlatDmg:    0.019 * c.MaxHP(),
	}
	c.Core.QueueAttack(ai, combat.NewDefCircHit(2, false, combat.TargettableEnemy), 0, 0)

	//create a construct
	con := &stoneStele{
		src:    c.Core.F,
		expiry: c.Core.F + dur,
		c:      c,
	}

	num := c.Core.Constructs.CountByType(construct.GeoConstructZhongliSkill)

	c.Core.Constructs.New(con, num == c.maxStele)

	c.steleCount = c.Core.Constructs.CountByType(construct.GeoConstructZhongliSkill)

	c.Core.Log.NewEvent(
		"Stele added",
		glog.LogCharacterEvent,
		c.Index,
		"orig_count", num,
		"cur_count", c.steleCount,
		"max_hit", max,
		"next_tick", c.Core.F+120,
	)
	// Snapshot buffs for resonance ticks
	aiSnap := combat.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "Stone Stele (Tick)",
		AttackTag:  combat.AttackTagElementalArt,
		ICDTag:     combat.ICDTagElementalArt,
		ICDGroup:   combat.ICDGroupDefault,
		StrikeType: combat.StrikeTypeBlunt,
		Element:    attributes.Geo,
		Durability: 25,
		Mult:       skillTick[c.TalentLvlSkill()],
		FlatDmg:    0.019 * c.MaxHP(),
	}
	snap := c.Snapshot(&aiSnap)
	c.steleSnapshot = combat.AttackEvent{
		Info:        aiSnap,
		Snapshot:    snap,
		Pattern:     combat.NewDefCircHit(1, false, combat.TargettableEnemy),
		SourceFrame: c.Core.F,
	}

	c.Core.Tasks.Add(c.resonance(c.Core.F, max), 120)
}

func (c *char) resonance(src, max int) func() {
	return func() {
		c.Core.Log.NewEvent("Stele checking for tick", glog.LogCharacterEvent, c.Index, "src", src, "char", c.Index)
		if !c.Core.Constructs.Has(src) {
			return
		}
		c.Core.Log.NewEvent("Stele ticked", glog.LogCharacterEvent, c.Index, "next expected", c.Core.F+120, "src", src, "char", c.Index)

		// Use snapshot for damage
		ae := c.steleSnapshot

		//check how many times to hit
		count := c.Core.Constructs.Count() - c.Core.Constructs.CountByType(construct.GeoConstructZhongliSkill) + 1
		if count > max {
			count = max
		}
		orb := false
		for i := 0; i < count; i++ {
			c.Core.QueueAttackEvent(&ae, 0)
			if c.energyICD < c.Core.F && !orb && c.Core.Rand.Float64() < .5 {
				orb = true
			}
		}
		if orb {
			c.energyICD = c.Core.F + 90
			c.Core.QueueParticle("zhongli", 1, attributes.Geo, 120)
		}
		c.Core.Tasks.Add(c.resonance(src, max), 120)
	}
}
