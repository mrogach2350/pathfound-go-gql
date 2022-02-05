package main

import (
	firebase "firebase.google.com/go/v4"
	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mrogach2350/pathfound_go/handlers"
	"log"
	"net/http"
	"strings"
)

type AuthHeader struct {
	Authorization string `header:"Authorization"`
}

func MyJwtMiddleware(app *firebase.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := AuthHeader{}
		err := c.BindHeader(&h)
		if err != nil {
			log.Fatalf("error getting Auth client: %v\n", err)
		}

		splitToken := strings.Split(h.Authorization, "Bearer ")
		if len(splitToken) != 2 {
			c.JSON(http.StatusForbidden, gin.H{"message": "missing JWT"})
			c.Abort()
		}
		idToken := splitToken[1]

		client, err := app.Auth(c)
		if err != nil {
			log.Fatalf("error getting Auth client: %v\n", err)
		}

		_, err = client.VerifyIDToken(c, idToken)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"message": "invalid JWT"})
			c.Abort()
		}

		c.Next()
	}
}

// run this command every session
// export GOOGLE_APPLICATION_CREDENTIALS=/Users/mrogach/Code/pathfound/pathfound_go/pathfound-go-firebase-adminsdk-uvyop-3f7ba6d110.json
func main() {
	localDataPtr := flag.Bool("localData", true, "load data from ./data or from GH")
	flag.Parse()
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	r.Use(cors.New(config))

	api := r.Group("/api")
	pathfinderEndpoints := api.Group("/pathfinder")
	fifthedEndpoints := api.Group("/fifthed")

	//app, err := firebase.NewApp(context.Background(), nil)
	//if err != nil {
	//	log.Fatalf("error initializing app: %v\n", err)
	//}
	//r.Use(MyJwtMiddleware(app))

	pfHandler := handlers.PathfinderHandler{}
	pfHandler.BindData(localDataPtr)

	feHandler := handlers.FifthEdHandler{}
	feHandler.BindData()

	pathfinderEndpoints.GET("/:type", pfHandler.GetAllRecordsHandler)
	pathfinderEndpoints.GET("/:type/:id", pfHandler.GetRecordByIdHandler)

	fifthedEndpoints.GET("/:type", feHandler.GetAllRecordsHandler)
	r.Run()
}
