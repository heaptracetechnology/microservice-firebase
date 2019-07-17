package messaging

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/NaySoftware/go-fcm"
	"github.com/heaptracetechnology/microservice-firebase/result"
)

type ArgsData struct {
	Token           string      `json:"token"`
	Topic           string      `json:"topic"`
	Title           string      `json:"title"`
	Body            string      `json:"body"`
	Icon            string      `json:"icon"`
	Data            interface{} `json:"data"`
	RegistrationIds []string    `json:"registrationIds"`
}

//Send Message By Token
func SendMessageByToken(responseWriter http.ResponseWriter, request *http.Request) {

	var serverKey = os.Getenv("SERVER_KEY")

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}
	defer request.Body.Close()

	var argsdata ArgsData
	er := json.Unmarshal(body, &argsdata)
	if er != nil {
		result.WriteErrorResponse(responseWriter, er)
		return
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

	response, sendErr := client.Send()
	if sendErr != nil || response.StatusCode == 401 {
		result.WriteErrorResponse(responseWriter, sendErr)
		return
	}
	bytes, _ := json.Marshal(response)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//Send Message By Topic
func SendMessageByTopic(responseWriter http.ResponseWriter, request *http.Request) {

	var serverKey = os.Getenv("SERVER_KEY")

	client := fcm.NewFcmClient(serverKey)
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	defer request.Body.Close()

	var argsdata ArgsData
	er := json.Unmarshal(body, &argsdata)
	if er != nil {
		result.WriteErrorResponse(responseWriter, er)
		return
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
	response, sendErr := client.Send()
	if sendErr != nil || response.StatusCode == 401 {
		result.WriteErrorResponse(responseWriter, sendErr)
		return
	}
	bytes, _ := json.Marshal(response)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
