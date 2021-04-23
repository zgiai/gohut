package main

import (
	"github.com/go-ini/ini"
	"github.com/gohut/gin-kit/illuminate/Database/Connection"
	"github.com/gohut/gin-kit/pkg/env"
	//"github.com/go-ini/ini"
	//"log"
	"time"
	//"log"
	//"net/http"

)
var (
	Cfg *ini.File

	RunMode string

	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)
func main() {
	//var err error
	//router := gin.Default()
	env.New()
	Connection.New()
	//RunMode = env.Load("RUN_MODE","")
	//HttpPort := env.Load("HTTP_PORT","")
	//readTimeOut := env.Load("READ_TIMEOUT","")
	//writeTimeOut := env.Load("WRITE_TIMEOUT","")
	//fmt.Println("Cfg",RunMode)
	//Cfg, err := ini.Load(".env")
	//if err != nil {
	//	log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	//}
	//RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
	//RunMode = Cfg.Section("database").Key("TYPE").String()
	//RunMode = Env("RUN_MODE","")
	//fmt.Println("Cfg",env.New())
	//router.GET("/test", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "test",
	//	})
	//})
	////fmt.Println(setting)gin
	//s := &http.Server{
	//	Addr:           fmt.Sprintf(":%d", HttpPort),
	//	Handler:        router,
	//	ReadTimeout:    readTimeOut,
	//	WriteTimeout:   writeTimeOut,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//
	//s.ListenAndServe()
}