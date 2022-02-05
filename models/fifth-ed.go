package models

type MagicVariantJson struct {
	MagicVariants []MagicVariant `json:"variant"`
}
type MagicVariant struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Requires []struct {
		Type string `json:"type"`
	} `json:"requires"`
	Ammo     bool `json:"ammo"`
	Inherits struct {
		NamePrefix  string   `json:"namePrefix"`
		Source      string   `json:"source"`
		Page        int      `json:"page"`
		Srd         bool     `json:"srd"`
		Tier        string   `json:"tier"`
		Rarity      string   `json:"rarity"`
		BonusWeapon string   `json:"bonusWeapon"`
		Entries     []string `json:"entries"`
		LootTables  []string `json:"lootTables"`
	} `json:"inherits"`
}

type PsionicJson struct {
	Psionics []Psionic `json:"psionic"`
}
type Psionic struct {
	Name    string   `json:"name"`
	Source  string   `json:"source"`
	Page    int      `json:"page"`
	Type    string   `json:"type"`
	Order   string   `json:"order"`
	Entries []string `json:"entries"`
	Focus   string   `json:"focus"`
	Modes   []struct {
		Cost struct {
			Min int `json:"min"`
			Max int `json:"max"`
		} `json:"cost"`
		Name          string   `json:"name"`
		Entries       []string `json:"entries"`
		Concentration struct {
			Duration int    `json:"duration"`
			Unit     string `json:"unit"`
		} `json:"concentration,omitempty"`
	} `json:"modes"`
}

type RecipeJson struct {
	Recipes []Recipe `json:"recipe"`
}
type Recipe struct {
	Name      string   `json:"name"`
	Source    string   `json:"source"`
	Page      int      `json:"page"`
	Type      string   `json:"type"`
	DishTypes []string `json:"dishTypes"`
	Diet      string   `json:"diet"`
	Serves    struct {
		Note  string `json:"note"`
		Exact int    `json:"exact"`
	} `json:"serves"`
	Ingredients []struct {
		Type    string `json:"type"`
		Entry   string `json:"entry"`
		Amount1 int    `json:"amount1"`
	} `json:"ingredients"`
	Instructions   []string `json:"instructions"`
	HasFluff       bool     `json:"hasFluff"`
	HasFluffImages bool     `json:"hasFluffImages"`
}
