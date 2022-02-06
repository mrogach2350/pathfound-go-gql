package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mrogach2350/pathfound_go/helpers"
	"github.com/mrogach2350/pathfound_go/models"
	"net/http"
)

type FifthEdHandler struct {
	MagicVariantsJson models.MagicVariantJson
	PsionicsJson      models.PsionicJson
	RecipesJson       models.RecipeJson
	ItemsJson         models.ItemJson
}

var _ IHandler = &FifthEdHandler{}

func (h *FifthEdHandler) BindData(_ *bool) {
	helpers.BindLocalData(models.LocalFifthEditionUrl+models.MagicVariantUri, &h.MagicVariantsJson)
	helpers.BindLocalData(models.LocalFifthEditionUrl+models.PsionicUri, &h.PsionicsJson)
	helpers.BindLocalData(models.LocalFifthEditionUrl+models.RecipeUri, &h.RecipesJson)
	helpers.BindLocalData(models.LocalFifthEditionUrl+models.ItemsUri, &h.ItemsJson)
}

func (h *FifthEdHandler) GetAllRecordsHandler(c *gin.Context) {
	start, end := helpers.GetStartEndPagination(c)
	entityType, exists := c.Params.Get("type")
	if exists {
		switch entityType {
		case "magic-variants":
			c.JSON(http.StatusOK, gin.H{"magic-variants": h.MagicVariantsJson.MagicVariants[start:end]})
		case "psionics":
			c.JSON(http.StatusOK, gin.H{"psionics": h.PsionicsJson.Psionics[start:end]})
		case "recipes":
			c.JSON(http.StatusOK, gin.H{"recipes": h.RecipesJson.Recipes[start:end]})
		case "items":
			c.JSON(http.StatusOK, gin.H{"recipes": h.ItemsJson.Items[start:end]})
		case "item-groups":
			c.JSON(http.StatusOK, gin.H{"recipes": h.ItemsJson.ItemGroups[start:end]})
		default:
			fmt.Printf("no matching case for %s\n", entityType)
			c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("can not find data for %s", entityType)})
		}
	}
}
