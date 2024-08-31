package simulation

import (
	// stats collectors
	_ "github.com/genshinsim/gcsim/pkg/stats/action"
	_ "github.com/genshinsim/gcsim/pkg/stats/damage"
	_ "github.com/genshinsim/gcsim/pkg/stats/endstats"
	_ "github.com/genshinsim/gcsim/pkg/stats/energy"
	_ "github.com/genshinsim/gcsim/pkg/stats/heal"
	_ "github.com/genshinsim/gcsim/pkg/stats/position"
	_ "github.com/genshinsim/gcsim/pkg/stats/reaction"
	_ "github.com/genshinsim/gcsim/pkg/stats/shield"
	_ "github.com/genshinsim/gcsim/pkg/stats/status"
	_ "github.com/genshinsim/gcsim/pkg/stats/swap"

	// artifacts
	_ "github.com/genshinsim/gcsim/internal/artifacts/archaic"
	_ "github.com/genshinsim/gcsim/internal/artifacts/berserker"
	_ "github.com/genshinsim/gcsim/internal/artifacts/blizzard"
	_ "github.com/genshinsim/gcsim/internal/artifacts/bloodstained"
	_ "github.com/genshinsim/gcsim/internal/artifacts/bolide"
	_ "github.com/genshinsim/gcsim/internal/artifacts/braveheart"
	_ "github.com/genshinsim/gcsim/internal/artifacts/crimson"
	_ "github.com/genshinsim/gcsim/internal/artifacts/deepwood"
	_ "github.com/genshinsim/gcsim/internal/artifacts/defenderswill"
	_ "github.com/genshinsim/gcsim/internal/artifacts/desertpavilionchronicle"
	_ "github.com/genshinsim/gcsim/internal/artifacts/echoes"
	_ "github.com/genshinsim/gcsim/internal/artifacts/emblem"
	_ "github.com/genshinsim/gcsim/internal/artifacts/exile"
	_ "github.com/genshinsim/gcsim/internal/artifacts/flowerofparadiselost"
	_ "github.com/genshinsim/gcsim/internal/artifacts/fragmentofharmonicwhimsy"
	_ "github.com/genshinsim/gcsim/internal/artifacts/gambler"
	_ "github.com/genshinsim/gcsim/internal/artifacts/gildeddreams"
	_ "github.com/genshinsim/gcsim/internal/artifacts/gladiator"
	_ "github.com/genshinsim/gcsim/internal/artifacts/goldentroupe"
	_ "github.com/genshinsim/gcsim/internal/artifacts/heartofdepth"
	_ "github.com/genshinsim/gcsim/internal/artifacts/huskofopulentdreams"
	_ "github.com/genshinsim/gcsim/internal/artifacts/instructor"
	_ "github.com/genshinsim/gcsim/internal/artifacts/lavawalker"
	_ "github.com/genshinsim/gcsim/internal/artifacts/maiden"
	_ "github.com/genshinsim/gcsim/internal/artifacts/marechausseehunter"
	_ "github.com/genshinsim/gcsim/internal/artifacts/martialartist"
	_ "github.com/genshinsim/gcsim/internal/artifacts/nighttimewhispersintheechoingwoods"
	_ "github.com/genshinsim/gcsim/internal/artifacts/noblesse"
	_ "github.com/genshinsim/gcsim/internal/artifacts/nymphsdream"
	_ "github.com/genshinsim/gcsim/internal/artifacts/oceanhuedclam"
	_ "github.com/genshinsim/gcsim/internal/artifacts/paleflame"
	_ "github.com/genshinsim/gcsim/internal/artifacts/reminiscence"
	_ "github.com/genshinsim/gcsim/internal/artifacts/scholar"
	_ "github.com/genshinsim/gcsim/internal/artifacts/sojourner"
	_ "github.com/genshinsim/gcsim/internal/artifacts/songofdayspast"
	_ "github.com/genshinsim/gcsim/internal/artifacts/tenacity"
	_ "github.com/genshinsim/gcsim/internal/artifacts/thunderingfury"
	_ "github.com/genshinsim/gcsim/internal/artifacts/thundersoother"
	_ "github.com/genshinsim/gcsim/internal/artifacts/unfinishedreverie"
	_ "github.com/genshinsim/gcsim/internal/artifacts/vermillion"
	_ "github.com/genshinsim/gcsim/internal/artifacts/viridescent"
	_ "github.com/genshinsim/gcsim/internal/artifacts/vourukashasglow"
	_ "github.com/genshinsim/gcsim/internal/artifacts/wanderer"

	// weapons
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/alley"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/amos"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/aqua"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/blackcliff"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/cloudforged"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/compound"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/elegy"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/endoftheline"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/favonius"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/firstgreatmagic"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/hamayumi"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/heartstrings"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/huntersbow"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/hunterspath"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/ibispiercer"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/kingssquire"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/messenger"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/mitternachtswaltz"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/mouunsmoon"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/polarstar"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/predator"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/prototype"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/rangegauge"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/raven"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/recurve"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/royal"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/rust"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/sacrificial"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/scionoftheblazingsun"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/seasonedhuntersbow"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/sharpshooter"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/skyward"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/slingshot"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/songofstillness"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/stringless"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/thundering"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/twilight"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/viridescent"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/windblume"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/apprentice"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/athousandfloatingdreams"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/balladoftheboundlessblue"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/blackcliff"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/cashflow"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/cranesechoingcall"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/dodoco"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/emeraldorb"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/eternalflow"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/favonius"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/flowingpurity"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/frostbearer"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/fruitoffulfillment"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/hakushin"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/jadefallssplendor"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/kagura"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/magicguide"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/mappa"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/memory"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/moonglow"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/oathsworneye"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/otherworldly"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/perception"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/pocket"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/prayer"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/prototype"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/ringofyaxche"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/royal"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/sacrifical"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/sacrificialjade"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/skyward"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/solar"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/thrilling"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/tulaytullahsremembrance"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/twin"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/wanderingevenstar"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/widsith"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/wine"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/akuoumaru"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/aquamarine"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/beacon"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/bell"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/blackcliff"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/bloodtainted"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/debateclub"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/earthshaker"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/favonius"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/ferrousshadow"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/forestregalia"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/lithic"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/mailedflower"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/nagamasa"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/oldmercspal"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/pines"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/powersaw"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/prototype"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/rainslasher"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/redhorn"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/royal"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/sacrifical"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/sealord"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/skyrider"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/skyward"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/spine"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/starsilver"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/talkingstick"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/tidalshadow"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/ultimateoverlordsmegamagicsword"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/unforged"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/verdict"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/waster"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/whiteblind"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/whiteiron"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/wolf"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/balladofthefjords"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/beginners"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/blackcliff"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/blacktassel"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/calamity"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/catch"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/crescent"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/crimsonmoonssemblance"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/deathmatch"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/dialoguesofthedesertsages"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/dragonbane"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/dragonspine"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/engulfing"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/favonius"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/footprint"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/halberd"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/homa"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/ironpoint"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/kitain"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/lithic"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/lumidouceelegy"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/missive"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/moonpiercer"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/primordial"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/prospectorsdrill"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/prototype"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/rightfulreward"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/royal"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/scarletsands"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/skyward"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/vortex"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/wavebreaker"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/whitetassel"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/absolution"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/alley"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/amenoma"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/aquila"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/blackcliff"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/blacksword"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/cinnabar"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/coolsteel"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/darkironsword"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/dockhand"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/dullblade"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/favonius"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/festering"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/filletblade"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/finaleofthedeep"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/fleuvecendreferryman"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/flute"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/foliar"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/freedom"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/haran"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/harbinger"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/ironsting"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/kagotsurubeisshin"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/keyofkhajnisut"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/lion"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/mistsplitter"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/primordial"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/prototype"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/royal"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/sacrifical"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/sapwoodblade"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/silversword"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/skyrider"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/skyward"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/splendoroftranquilwaters"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/summit"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/swordofdescension"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/swordofnarzissenkreuz"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/toukaboushigure"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/travelershandysword"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/urakumisugiri"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/wolffang"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/xiphos"
)
