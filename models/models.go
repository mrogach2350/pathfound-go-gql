package models

const (
	BasePathfinderUrl  = "https://www.mrogach.com/pathfinder-data/"
	LocalPathfinderUrl = "./data/pathfinder/"
	WeaponsUri         = "baseWeapons.json"
	ArmorUri           = "baseArmor.json"
	SpellsUri          = "spell.json"
	CasterTypesUri     = "casterType.json"
	FeatsUri           = "feats.json"
	MagicItemsUri      = "magicItems.json"
	MonstersUri        = "monsters.json"
	RacesUri           = "races.json"
)

const (
	BaseFifthEditionUrl  = ""
	LocalFifthEditionUrl = "./data/5e-data/"
	MagicVariantUri      = "magicvariants.json"
	VariantKey           = "variant"
	PsionicUri           = "psionics.json"
	PsionicKey           = "psionic"
	RecipeUri            = "recipes.json"
	RecipeKey            = "recipe"
)

var PathfinderTypes = map[string]bool{
	"*[]models.Weapon":     true,
	"*[]models.Armor":      true,
	"*[]models.Spell":      true,
	"*[]models.CasterType": true,
	"*[]models.Feat":       true,
	"*[]models.MagicItem":  true,
	"*[]models.Monster":    true,
	"*[]models.Race":       true,
}

var FifthEdTypes = map[string]bool{
	"*models.MagicVariantJson": true,
	"*models.PsionicJson":      true,
	"*models.RecipeJson":       true,
}
