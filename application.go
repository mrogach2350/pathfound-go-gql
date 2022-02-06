package main

import (
	"flag"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/mrogach2350/pathfound_go/graph"
	"github.com/mrogach2350/pathfound_go/graph/generated"
	"log"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mrogach2350/pathfound_go/handlers"
)

type AuthHeader struct {
	Authorization string `header:"Authorization"`
}

func MyJwtMiddleware(client *auth.Client) gin.HandlerFunc {
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

		_, err = client.VerifyIDToken(c, idToken)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"message": "invalid JWT"})
			c.Abort()
		}

		c.Next()
	}
}

// Defining the Graphql handler
func GraphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
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

	//ctx := context.Background()
	//app, err := firebase.NewApp(ctx, nil)
	//if err != nil {
	//	log.Fatalf("error initializing app: %v\n", err)
	//}
	//authClient, err := app.Auth(ctx)
	//if err != nil {
	//	log.Fatalf("error initializing app: %v\n", err)
	//}
	// if !*localDataPtr {
	//r.Use(MyJwtMiddleware(authClient))
	// }
	//
	//firestoreClient, err := app.Firestore(ctx)
	//if err != nil {
	//	log.Fatalf("error initializing app: %v\n", err)
	//}
	//fbHandler := handlers.FirebaseHandler{C: firestoreClient}
	//fbHandler.Test()

	pfHandler := handlers.PathfinderHandler{}
	pfHandler.BindData(localDataPtr)
	pathfinderEndpoints := api.Group("/pathfinder")
	pathfinderEndpoints.GET("/:type", pfHandler.GetAllRecordsHandler)
	pathfinderEndpoints.GET("/:type/:id", pfHandler.GetRecordByIdHandler)

	feHandler := handlers.FifthEdHandler{}
	feHandler.BindData(localDataPtr)
	fifthedEndpoints := api.Group("/fifthed")
	fifthedEndpoints.GET("/:type", feHandler.GetAllRecordsHandler)

	r.POST("/query", GraphqlHandler())
	r.GET("/", PlaygroundHandler())

	r.Run()
}
