package developer

import (
	"encoding/json"
	"net/http"

	"github.com/Aleena48/Alert-System/model"
	"github.com/gin-gonic/gin"
)

type developer struct {
	ID       int64  `json:"id,omitempty"`
	FullName string `json:"full_name,omitempty"`
	Email    string `json:"email,omitempty"`
	Mobile   int64  `json:"mobile,omitempty"`
	TeamId   int64  `json:"team_id,omitempty"`
}

// func to create new developer details
func CreateDeveloper(ctx *gin.Context) {
	var dev developer
	payload, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		model.Logger.Println(err)
		return
	}
	model.Logger.Println("Payload=", payload)
	err = json.Unmarshal(payload, &dev)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		model.Logger.Println(err)
		return
	}
	model.Logger.Println("Developer=", dev)

	row := model.DB.QueryRow(
		`insert into developer(full_name,email,mobile) values($1, $2, $3) returning id;`,
		dev.FullName, dev.Email, dev.Mobile,
	)
	err = row.Scan(&dev.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		model.Logger.Println(err)
		return
	}
	model.Logger.Println("Table data insterted")
	ctx.JSON(http.StatusOK, dev)
}

//func to list all developers
func ListDeveloper(ctx *gin.Context) {
	var devList []developer
	rows, err := model.DB.Query(`select id, full_name,email,mobile from developer`)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		model.Logger.Println(err)
		return
	}
	for rows.Next() {
		var dev developer
		err := rows.Scan(&dev.ID, &dev.FullName, &dev.Email, &dev.Mobile)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			model.Logger.Println(err)
			return
		}
		devList = append(devList, dev)
	}
	model.Logger.Println("Developer list")
	ctx.JSON(http.StatusOK, devList)
}

// func to retrive individual developer data
func GetDeveloper(ctx *gin.Context) {
	id := ctx.Param("id")
	var dev developer
	row := model.DB.QueryRow(`select id, full_name,email,mobile from developer where id = $1`, id)
	err := row.Scan(&dev.ID, &dev.FullName, &dev.Email, &dev.Mobile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		model.Logger.Println(err)
		return
	}
	model.Logger.Println("Developer data on passed ID")
	ctx.JSON(http.StatusOK, dev)
}

//func to delete a specified developer data
func DeleteDeveloper(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := model.DB.Exec(`delete from developer where id = $1`, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		model.Logger.Println(err)
		return
	}
	model.Logger.Println("Table data list")
	ctx.JSON(http.StatusOK, id)
}

//func to update developer data
func UpdateDeveloper(ctx *gin.Context) {
	id := ctx.Param("id")
	var dev developer
	payload, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		model.Logger.Println(err)
		return
	}
	model.Logger.Println("Payload=", payload)
	err = json.Unmarshal(payload, &dev)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		model.Logger.Println(err)
		return
	}
	model.Logger.Println("Team=", dev)
	_, err = model.DB.Exec(`update developer set full_name= $1, email=$2, mobile=$3 where id =$4`, dev.FullName, dev.Email, dev.Mobile, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		model.Logger.Println(err)
		return
	}
	model.Logger.Println("Updated Developer list")
	ctx.JSON(http.StatusOK, dev)
}

func MapDevToTeam(developerIDs []int64, teamID interface{}) {

	for _, val := range developerIDs {
		_, err := model.DB.Exec(`update developer set team_id=$1 where id=$2`, teamID, val)
		if err != nil {
			model.Logger.Println(err)
			return
		}
	}
}
