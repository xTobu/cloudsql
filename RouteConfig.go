package main

import (
	//"../Controllers"
	"github.com/gin-gonic/gin"
	//go-mssqldb
	//_ "github.com/denisenkom/go-mssqldb"

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
	config.Port = ":8000"
	config.StaticFolder = "../dist"
	config.IndexFile = "../index.html"
}

////////////////////

// Init blablaba
func start() {
	// set config
	config := Config{}
	config.SetDefault()

	// Creates a default gin router
	router := gin.Default() // Grouping routes

	//group： url //首頁
	url := router.Group("/")
	{
		url.GET("/", Hello)

	}
	router.Run(config.Port) // listen and serve on 0.0.0.0:8000

}
