package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize: 1024*10,
    WriteBufferSize: 1024*10,
}

func handler(w http.ResponseWriter, r* http.Request) {
    // allow connection to all CORS
    upgrader.CheckOrigin = func(r *http.Request) bool { return true }

    connection,err := upgrader.Upgrade(w,r,nil);
    // if upgrader is not build close
    if err != nil {
        fmt.Println(err);
        return
    }

    defer connection.Close();

    for{
        messageType, message, err := connection.ReadMessage();
        if err != nil {
            fmt.Println(err);
            return
        }
        // log the input message recieved
        fmt.Println(string(message));

        // socket cannot understand strings thus creating bytes buffer
        if err:= connection.WriteMessage(messageType,message); err!=nil {
            fmt.Println(err);
            return
        }
    }
}

func main() {
    // what happens when / endpoint is hit (handler function is called)
    http.HandleFunc("/",handler);
    http.ListenAndServe(":8080",nil);

}
