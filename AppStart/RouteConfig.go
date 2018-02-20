package AppStart

import (
	//"../Controllers"
	"github.com/gin-gonic/gin"
	//go-mssqldb
	//_ "github.com/denisenkom/go-mssqldb"
	"cloudsql/Controllers"
	//mysql
	_ "github.com/go-sql-driver/mysql"
)

// ========== server

//Config struct
type Config struct {
	Port         string
	StaticFolder string
	IndexFile    string
}

//SetDefault Sever data
func (config *Config) SetDefault() {
	config.Port = ":8080"
	config.StaticFolder = "../dist"
	config.IndexFile = "../index.html"
}

////////////////////

// RouteConfig blablaba
func RouteConfig() {
	// set config
	config := Config{}
	config.SetDefault()

	// Creates a default gin router
	router := gin.Default() // Grouping routes

	//group： url //首頁
	url := router.Group("/")
	{
		url.GET("/", Controllers.Hello)

	}

	//group： api //api
	api := router.Group("/api")
	{
		api.GET("/get", Controllers.Student)
		api.POST("/post", Controllers.Insert)
	}

	router.Run(config.Port) // listen and serve on 0.0.0.0:8000

}
