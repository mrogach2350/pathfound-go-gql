package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mrogach2350/pathfound_go/models"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func BindLocalData(fullUrl string, i interface{}) {
	file, err := os.Open(fullUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Successfully Opened Local %s \n", fullUrl)
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, i)
}

func BindCloudData(uri string, i interface{}) {
	resp, err := http.Get(models.BasePathfinderUrl + uri)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Printf("Successfully Opened GH %s \n", uri)
	byteValue, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(byteValue, i)
}

func GetStartEndPagination(c *gin.Context) (int, int) {
	page := c.DefaultQuery("p", "1")
	pageLength := c.DefaultQuery("pl", "25")
	a, _ := strconv.Atoi(page)
	b, _ := strconv.Atoi(pageLength)
	start := (a - 1) * b
	end := a * b
	return start, end
}
