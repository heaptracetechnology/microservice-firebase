package Messaging

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"testing"
// )

// type TestArgsData struct {
// 	Token string      `json:"token"`
// 	Title string      `json:"title"`
// 	Body  string      `json:"body"`
// 	Icon  string      `json:"icon"`
// 	Data  interface{} `json:"data"`
// }

// func TestSendMessageByToken(t *testing.T) {

// 	testmessage := TestArgsData{Token: "epCGkWwiA2U:APA91bEy_1m6tCeOb-kpeKGCpnU6B3V3l-4jb1L5TY1_0dFdW2Q0XNJibwOyFA07HixW9rAllCegvL3LOii4Pdz0mA3AeF99TL7CcZb1iz-A0pG1qp0Y8qc2Yn1oDm1l-b73UtPVh-D7", Title: "Firebase Notification Message", Body: "Hello, Firebase message from cli", Icon: "firebase-logo.png"}
// 	reqbody := new(bytes.Buffer)
// 	json.NewEncoder(reqbody).Encode(testmessage)

// 	req, err := http.NewRequest("POST", "/send-message-by-token", reqbody)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	os.Setenv("SECRET_KEY", "AAAAY4LfSh4:APA91bEwy_Gn8glVVPfKmPiYKPx5nQVSW4XErmz0YBIpUqr9GrP2x6Zo-iHh-kMJn_v3mdGE9u2DB2HwaSiPX91zNcdgSlQke5Peti3AAFqt4DrPZ2fn4qJgCiHXH-OoZcPuxdNC-W8h")
// 	recorder := httptest.NewRecorder()
// 	handler := http.HandlerFunc(SendMessageByToken)

// 	handler.ServeHTTP(recorder, req)

// 	if status := recorder.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// }
