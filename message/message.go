package message

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/Aleena48/Alert-System/model"
	"github.com/gin-gonic/gin"
)

type message struct {
	TeamId  int64  `json:"team_id,omitempty"`
	Content string `json:"content,omitempty"`
	Title   string `json:"title,omitempty"`
}

type sms struct {
	Id      int64     `json:"id,omitempty"`
	Mobile  int64     `json:"mobile,omitempty"`
	Content string    `json:"content,omitempty"`
	MsgTime time.Time `json:"sent_at,omitempty"`
}
type email struct {
	Id      int64     `json:"id,omitempty"`
	Email   string    `json:"email,omitempty"`
	Content string    `json:"content,omitempty"`
	Title   string    `json:"title,omitempty"`
	MsgTime time.Time `json:"sent_at,omitempty"`
}
type alert struct {
	TeamId int64 `json:"team_id,omitempty"`
	Sms    sms   `json:"sms,omitempty"`
	Email  email `json:"email,omitempty"`
}

func CreateNotification(ctx *gin.Context) {
	var msg message
	var alertMsg alert
	payload, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		model.Logger.Println(err)
		return
	}
	model.Logger.Println("Payload=", payload)
	err = json.Unmarshal(payload, &msg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		model.Logger.Println(err)
		return
	}
	model.Logger.Println("Message=", msg)

	model.Logger.Println("Table data insterted")

	alertMsg = alert{
		TeamId: msg.TeamId,
		Sms: sms{
			Id:      rand.Int63(),
			Mobile:  0,
			Content: msg.Content,
			MsgTime: time.Now(),
		},
		Email: email{
			Id:      rand.Int63(),
			Email:   "",
			Content: msg.Content,
			Title:   msg.Title,
			MsgTime: time.Now(),
		},
	}
	ctx.JSON(http.StatusOK, alertMsg)
}
