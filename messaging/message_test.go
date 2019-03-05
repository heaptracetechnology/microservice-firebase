package messaging

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"os"
)

type TestArgsData struct {
	Token string      `json:"token"`
	Topic string      `json:"topic"`
	Title string      `json:"title"`
	Body  string      `json:"body"`
	Icon  string      `json:"icon"`
	Data  interface{} `json:"data"`
}

var _ = Describe("Firebase cloud messaging", func() {

	testmessage := TestArgsData{
		Token: "eNstKTcV5Dg:APA91bEGvEHaP6-UdcLBfgaib1lOPUZgrP1QYDAOUoZc_ZQNNlGO1afiR5lGYqbuJTc4YQ0yn3Xogjuj1GeryvvgkcutItfu0kjMwCTIN2CNdp9oiQBPm2394FxHjWMyW8ZgsL1p4xHo",
		Title: "Test cases",
		Body:  "Hello body from cli"}

	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(testmessage)

	os.Setenv("SECRET_KEY", "AAAAY4LfSh4:APA91bEwy_Gn8glVVPfKmPiYKPx5nQVSW4XErmz0YBIpUqr9GrP2x6Zo-iHh-kMJn_v3mdGE9u2DB2HwaSiPX91zNcdgSlQke5Peti3AAFqt4DrPZ2fn4qJgCiHXH-OoZcPuxdNC-W8h")
	req, err := http.NewRequest("POST", "/send-message-by-token", reqbody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessageByToken)
	handler.ServeHTTP(recorder, req)

	Describe("Send message by token", func() {
		Context("SendMessageByToken", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Firebase cloud messaging", func() {

	testmessage := TestArgsData{
		Token: "eNstKTcV5Dg:APA91bEGvEHaP6-UdcLBfgaib1lOPUZgrP1QYDAOUoZc_ZQNNlGO1afiR5lGYqbuJTc4YQ0yn3Xogjuj1GeryvvgkcutItfu0kjMwCTIN2CNdp9oiQBPm2394FxHjWMyW8ZgsL1p4xHo",
		Topic: "news",
		Title: "Test cases",
		Body:  "Hello body from cli"}
	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(testmessage)

	req, err := http.NewRequest("POST", "/send-message-by-topic", reqbody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessageByTopic)
	handler.ServeHTTP(recorder, req)

	Describe("Send message by topic", func() {
		Context("SendMessageByTopic", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})
