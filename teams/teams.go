package teams

import (
	"encoding/json"
	"net/http"

	"github.com/Aleena48/Alert-System/model"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type teams struct {
	ID          int64   `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	DeptName    string  `json:"dept_name,omitempty"`
	DeveloperId []int64 `json:"developer_ids,omitempty"`
}

// func to create new team
func CreateTeam(ctx *gin.Context) {
	var team teams
	payload, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		model.Logger.Println(err)
		return
	}
	model.Logger.Println("Payload=", payload)
	err = json.Unmarshal(payload, &team)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		model.Logger.Println(err)
		return
	}
	model.Logger.Println("Team=", team)

	row := model.DB.QueryRow(
		`insert into teams(name,dept_name, developer_ids) values($1, $2, $3) returning id;`,
		team.Name, team.DeptName, pq.Array(team.DeveloperId),
	)
	err = row.Scan(&team.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		model.Logger.Println(err)
		return
	}
	model.Logger.Println("Table data insterted")
	ctx.JSON(http.StatusOK, team)
}

// func to list all teams
func ListTeam(ctx *gin.Context) {
	var teamList []teams
	rows, err := model.DB.Query(`select * from teams`)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		model.Logger.Println(err)
		return
	}
	for rows.Next() {
		var team teams
		err := rows.Scan(&team.ID, &team.Name, &team.DeptName, pq.Array(&team.DeveloperId))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			model.Logger.Println(err)
			return
		}
		teamList = append(teamList, team)
	}
	model.Logger.Println("Table data list")
	ctx.JSON(http.StatusOK, teamList)
}

// func to retrive individual team data
func GetTeam(ctx *gin.Context) {
	id := ctx.Param("id")
	var team teams
	row := model.DB.QueryRow(`select * from teams where id = $1`, id)
	err := row.Scan(&team.ID, &team.Name, &team.DeptName, pq.Array(&team.DeveloperId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		model.Logger.Println(err)
		return
	}
	model.Logger.Println("Team data on passed ID")
	ctx.JSON(http.StatusOK, team)
}

//func to delete a specified team
func DeleteTeam(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := model.DB.Exec(`delete from teams where id = $1`, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		model.Logger.Println(err)
		return
	}
	model.Logger.Println("Table data list")
	ctx.JSON(http.StatusOK, id)
}

//func to update team data
func UpdateTeam(ctx *gin.Context) {
	id := ctx.Param("id")
	var team teams
	payload, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		model.Logger.Println(err)
		return
	}
	model.Logger.Println("Payload=", payload)
	err = json.Unmarshal(payload, &team)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		model.Logger.Println(err)
		return
	}
	model.Logger.Println("Team=", team)
	_, err = model.DB.Exec(`update teams set name= $1, dept_name=$2 , developer_ids=$3 where id =$4`, team.Name, team.DeptName, pq.Array(team.DeveloperId), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		model.Logger.Println(err)
		return
	}
	model.Logger.Println("Updated teams list")
	ctx.JSON(http.StatusOK, team)
}
