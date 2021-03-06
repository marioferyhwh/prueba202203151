package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/labstack/echo"
	"github.com/marioferyhwh/prueba202203151/pkg/routes"
)

func main() {
	Port := 8080
	hostName, err := os.Hostname()
	if err != nil {
		log.Println("error leyendo el host name")
		return
	}
	addrs, err := net.LookupHost(hostName)
	if err != nil {
		fmt.Printf("error leyendo las direcciones ip del hostname: %v\n", err)
		return
	}
	for _, a := range addrs {
		log.Println("direccion:", a)
	}

	e := echo.New()
	routes.InitRoutes(e)

	err = e.Start(fmt.Sprintf(":%d", Port))
	if err != nil {
		log.Println("servicion detenido ¿eso es lo que queria?")
		log.Println(err)
	}

}
