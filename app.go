package main

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"github.com/spf13/viper"
)

// Struct permit to have DB access on method
type Ressource struct {
	db *gorm.DB
}

// Init DB
func NewRessource(db *gorm.DB) Ressource {
	return Ressource{
		db: db,
	}
}

func startApp(db *gorm.DB) {
	log := logging.MustGetLogger("log")

	if viper.GetString("logtype") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	g := gin.Default()
	//r := NewRessource(db)

	g.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	g.Static("/", "./static")

	/*v1 := g.Group("api/v1")
	{
		v1.GET("/temperatures", r.GetTemperatures)
		v1.POST("/temperature", r.PostTemperature)
	}*/
	log.Debug("Port: %d", viper.GetInt("server.port"))
	g.Run(":" + strconv.Itoa(viper.GetInt("server.port")))
}

func main() {
	confPath := "cfg"
	confFilename := "server"
	logFilename := "error.log"

	fd := initLogging(&logFilename)
	defer fd.Close()

	loadConfig(&confPath, &confFilename)
	dbmap := Initdb()

	startApp(dbmap)
}
