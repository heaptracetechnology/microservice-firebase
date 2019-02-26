package Messaging

import (
	"encoding/json"
	"github.com/NaySoftware/go-fcm"
	"io/ioutil"
	"net/http"
	"os"
)

type ArgsData struct {
	Token string      `json:"token"`
	Title string      `json:"title"`
	Body  string      `json:"body"`
	Icon  string      `json:"icon"`
	Data  interface{} `json:"data"`
}

//Send Message By Token
func SendMessageByToken(w http.ResponseWriter, r *http.Request) {

	var serverKey = os.Getenv("SERVER_KEY")

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var argsdata ArgsData
	err = json.Unmarshal(b, &argsdata)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	client := fcm.NewFcmClient(serverKey)

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

	client.Message = *message

	response, err := client.Send()
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}
	bytes, _ := json.Marshal(response)
	WriteJsonResponse(w, bytes, http.StatusCreated)
}

func WriteErrorResponse(w http.ResponseWriter, err error) {
	msgbytes, _ := json.Marshal(err)
	WriteJsonResponse(w, msgbytes, http.StatusBadRequest)
}

func WriteJsonResponse(w http.ResponseWriter, bytes []byte, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}
