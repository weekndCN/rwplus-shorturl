package main

import (
    "fmt"
    "net/http"

    "github.com/weekndCN/shorturl/models"
    "github.com/weekndCN/shorturl/pkg/setting"
    "github.com/weekndCN/shorturl/routers"
)

func init() {
    setting.Setup()
    models.Setup()
}

func main() {
	routersInit := routers.InitRouter()

	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf("0.0.0.0:%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

    s := &http.Server{
        Addr:           endPoint,
        Handler:        routersInit,
        ReadTimeout:    readTimeout,
        WriteTimeout:   writeTimeout,
        MaxHeaderBytes: maxHeaderBytes,
    }

    s.ListenAndServe()
}
