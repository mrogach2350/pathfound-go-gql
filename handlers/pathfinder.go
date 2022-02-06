package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mrogach2350/pathfound_go/helpers"
	"github.com/mrogach2350/pathfound_go/models"
	"net/http"
)

type PathfinderHandler struct {
	Weapons     []models.Weapon
	Armor       []models.Armor
	Spells      []models.Spell
	CasterTypes []models.CasterType
	Feats       []models.Feat
	MagicItems  []models.MagicItem
	Monsters    []models.Monster
	Races       []models.Race
}

var _ IHandler = &PathfinderHandler{}

func (h *PathfinderHandler) BindData(localDataPtr *bool) {
	if *localDataPtr {
		helpers.BindLocalData(models.LocalPathfinderUrl+models.WeaponsUri, &h.Weapons)
		helpers.BindLocalData(models.LocalPathfinderUrl+models.ArmorUri, &h.Armor)
		helpers.BindLocalData(models.LocalPathfinderUrl+models.SpellsUri, &h.Spells)
		helpers.BindLocalData(models.LocalPathfinderUrl+models.CasterTypesUri, &h.CasterTypes)
		helpers.BindLocalData(models.LocalPathfinderUrl+models.FeatsUri, &h.Feats)
		helpers.BindLocalData(models.LocalPathfinderUrl+models.MagicItemsUri, &h.MagicItems)
		helpers.BindLocalData(models.LocalPathfinderUrl+models.MonstersUri, &h.Monsters)
		helpers.BindLocalData(models.LocalPathfinderUrl+models.RacesUri, &h.Races)
	} else {
		helpers.BindCloudData(models.WeaponsUri, &h.Weapons)
		helpers.BindCloudData(models.ArmorUri, &h.Armor)
		helpers.BindCloudData(models.SpellsUri, &h.Spells)
		helpers.BindCloudData(models.CasterTypesUri, &h.CasterTypes)
		helpers.BindCloudData(models.FeatsUri, &h.Feats)
		helpers.BindCloudData(models.MagicItemsUri, &h.MagicItems)
		helpers.BindCloudData(models.MonstersUri, &h.Monsters)
		helpers.BindCloudData(models.RacesUri, &h.Races)
	}
}

func (h *PathfinderHandler) GetAllRecordsHandler(c *gin.Context) {
	start, end := helpers.GetStartEndPagination(c)
	entityType, exists := c.Params.Get("type")
	if exists {
		switch entityType {
		case "weapons":
			c.JSON(http.StatusOK, gin.H{"weapons": h.Weapons[start:end]})
		case "armor":
			c.JSON(http.StatusOK, gin.H{"armor": h.Armor[start:end]})
		case "spells":
			c.JSON(http.StatusOK, gin.H{"spells": h.Spells[start:end]})
		case "casterTypes":
			c.JSON(http.StatusOK, gin.H{"casterTypes": h.CasterTypes[start:end]})
		case "races":
			c.JSON(http.StatusOK, gin.H{"races": h.Races[start:end]})
		case "monsters":
			c.JSON(http.StatusOK, gin.H{"monsters": h.Monsters[start:end]})
		case "feats":
			c.JSON(http.StatusOK, gin.H{"feats": h.Feats[start:end]})
		case "magicItems":
			c.JSON(http.StatusOK, gin.H{"magicItems": h.MagicItems[start:end]})
		default:
			fmt.Printf("no matching case for %s\n", entityType)
			c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("can not find data for %s", entityType)})
		}
	}
}

func (h *PathfinderHandler) GetRecordByIdHandler(c *gin.Context) {
	start, end := helpers.GetStartEndPagination(c)
	entityType, exists := c.Params.Get("type")
	if exists {
		switch entityType {
		case "weapons":
			c.JSON(http.StatusOK, gin.H{"weapons": h.Weapons[start:end]})
		case "armor":
			c.JSON(http.StatusOK, gin.H{"armor": h.Armor[start:end]})
		case "spells":
			c.JSON(http.StatusOK, gin.H{"spells": h.Spells[start:end]})
		case "casterTypes":
			c.JSON(http.StatusOK, gin.H{"casterTypes": h.CasterTypes[start:end]})
		case "races":
			c.JSON(http.StatusOK, gin.H{"races": h.Races[start:end]})
		case "monsters":
			c.JSON(http.StatusOK, gin.H{"monsters": h.Monsters[start:end]})
		case "feats":
			c.JSON(http.StatusOK, gin.H{"feats": h.Feats[start:end]})
		case "magicItems":
			c.JSON(http.StatusOK, gin.H{"magicItems": h.MagicItems[start:end]})
		default:
			fmt.Printf("no matching case for %s\n", entityType)
			c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("can not find data for %s", entityType)})
		}
	}
}
