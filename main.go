package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)


var HttpSrvHandler *http.Server


func HttpRun(){

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5*time.Second)
		fmt.Fprintf(w, "hello")
	})

	HttpSrvHandler = &http.Server{
		Addr:           "0.0.0.0:8889",
		//Handler:        &router.Router{},
		ReadTimeout:    time.Duration(10) * time.Second,
		WriteTimeout:   time.Duration(10) * time.Second,
	}
	go func() {
		log.Printf(" [INFO] HttpServerRun:%s\n","8889")
		if err := HttpSrvHandler.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n","8889", err)
		}
	}()

}




func HttpStop(){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	log.Printf(" [INFO] HttpServerStop stopped\n")
}




func main() {

	fmt.Println("------------------------服务开启：端口在8889-------------------------------------------")

	HttpRun()

	quit := make(chan os.Signal,1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit


	HttpStop()


}
