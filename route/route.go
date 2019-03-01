package route

import (
    "github.com/gorilla/mux"
    "github.com/user/microservice-firebase/messaging"
    "log"
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "SendMessageByToken",
        "POST",
        "/send-message-by-token",
        messaging.SendMessageByToken,
    },
    Route{
        "SendMessageByTopic",
        "POST",
        "/send-message-by-topic",
        messaging.SendMessageByTopic,
    },
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        log.Println(route.Name)
        handler = route.HandlerFunc

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }
    return router
}
