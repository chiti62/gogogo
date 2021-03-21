package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/chiti62/gogogo/ch2/global"
	"github.com/chiti62/gogogo/ch2/internal/routers"
	"github.com/chiti62/gogogo/ch2/pkg/setting"
	"github.com/gin-gonic/gin"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err) //os.Exit, no stacktrace
	}
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	fmt.Printf("%v, %v, %v", global.ServerSetting, global.AppSetting, global.DatabaseSetting)

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}
