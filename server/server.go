package alertsystem

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Aleena48/Alert-System/developer"
	"github.com/Aleena48/Alert-System/message"
	"github.com/Aleena48/Alert-System/model"
	"github.com/Aleena48/Alert-System/teams"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func AlertSystemService() {

	// create log folder
	err := os.MkdirAll("../logs", 0777)
	if err != nil {
		log.Fatalln(err)
	}

	// open file to write log, create if does not exsist, truncte if file exsist
	f, err := os.OpenFile("../logs/alertsystem.log", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	// logger to log the  log date,time and line number
	model.Logger = log.New(f, "alertsystem :: ", log.Ldate|log.Ltime|log.Lshortfile)

	// creating db alertsystem with default port
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "mysecretpassword", "alertsystem")

	// opening db connection
	model.DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		model.Logger.Fatalln(err)
	}

	err = model.DB.Ping()
	if err != nil {
		model.Logger.Fatalln(err)
	}
	model.Logger.Println("connection was successfull")
	defer model.DB.Close()

	// open file to write gin log, create if does not exsist, truncte if file exsist
	ginFile, err := os.OpenFile("../logs/restApi.log", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		model.Logger.Fatalf("error opening file: %v", err)
	}
	defer ginFile.Close()
	// write gin log into file
	gin.DefaultWriter = ginFile

	// http router for gin framework
	router := gin.Default()

	router.POST("/teams", teams.CreateTeam)
	router.GET("/teams", teams.ListTeam)
	router.GET("/teams/:id", teams.GetTeam)
	router.DELETE("/teams/:id", teams.DeleteTeam)
	router.PUT("/teams/:id", teams.UpdateTeam)

	router.POST("/developer", developer.CreateDeveloper)
	router.GET("/developer", developer.ListDeveloper)
	router.GET("/developer/:id", developer.GetDeveloper)
	router.DELETE("/developer/:id", developer.DeleteDeveloper)
	router.PUT("/developer/:id", developer.UpdateDeveloper)

	router.POST("/trigger_notification", message.CreateNotification)

	model.Logger.Println("strating http service on 8080")
	// setting default port as 8080 for local routing
	router.Run()
}
