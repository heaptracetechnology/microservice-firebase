package messaging

import (
	"encoding/json"
	"github.com/NaySoftware/go-fcm"
	"github.com/heaptracetechnology/microservice-firebase/result"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ArgsData struct {
	Token           string      `json:"token"`
	Topic           string      `json:"topic"`
	Title           string      `json:"title"`
	Body            string      `json:"body"`
	Icon            string      `json:"icon"`
	Data            interface{} `json:"data"`
	RegistrationIds []string    `json:"registration_ids"`
}

//Send Message By Token
func SendMessageByToken(w http.ResponseWriter, r *http.Request) {

	var serverKey = os.Getenv("SERVER_KEY")

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		result.WriteErrorResponse(w, err)
	}

	var argsdata ArgsData
	err = json.Unmarshal(body, &argsdata)
	if err != nil {
		result.WriteErrorResponse(w, err)
	}

	client := fcm.NewFcmClient(serverKey)

	notification := &fcm.NotificationPayload{
		Title: argsdata.Title,
		Body:  argsdata.Body,
		Icon:  argsdata.Icon,
	}

	message := &fcm.FcmMsg{
		Data:            argsdata.Data,
		To:              argsdata.Token,
		Notification:    *notification,
		RegistrationIds: argsdata.RegistrationIds,
	}

	client.Message = *message

	response, err := client.Send()
	if err != nil {
		result.WriteErrorResponse(w, err)
	}
	bytes, _ := json.Marshal(response)
	result.WriteJsonResponse(w, bytes, http.StatusOK)
}

//Send Message By Topic
func SendMessageByTopic(w http.ResponseWriter, r *http.Request) {

	var serverKey = os.Getenv("SERVER_KEY")
	client := fcm.NewFcmClient(serverKey)
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		result.WriteErrorResponse(w, err)
	}

	var argsdata ArgsData
	err = json.Unmarshal(body, &argsdata)
	if err != nil {
		result.WriteErrorResponse(w, err)
	}

	notification := &fcm.NotificationPayload{
		Title: argsdata.Title,
		Body:  argsdata.Body,
		Icon:  argsdata.Icon,
	}
	message := &fcm.FcmMsg{
		Data:         argsdata.Data,
		To:           argsdata.Token,
		Notification: *notification,
	}

	token := argsdata.Token
	topic := argsdata.Topic
	to := token + "/" + topic

	client.NewFcmMsgTo(to, message)
	response, err := client.Send()
	if err != nil {
		log.Fatalln(err)
	}
	bytes, _ := json.Marshal(response)
	result.WriteJsonResponse(w, bytes, http.StatusOK)
}
