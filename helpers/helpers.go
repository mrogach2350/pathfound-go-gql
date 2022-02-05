package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mrogach2350/pathfound_go/models"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
)

func BindLocalData(uri string, i interface{}) {
	var url string
	typeString := reflect.TypeOf(i).String()
	if models.PathfinderTypes[typeString] {
		url = models.LocalPathfinderUrl
	}
	if models.FifthEdTypes[typeString] {
		url = models.LocalFifthEditionUrl
	}
	file, err := os.Open(url + uri)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Successfully Opened Local %s \n", uri)
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
