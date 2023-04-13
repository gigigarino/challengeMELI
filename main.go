//resolver problema del go
//go mod init 'yourmodulename' 
//go get 'yourpckagename'
package main

import (
	"log"
	"net/http"
	
	"github.com/gin-gonic/gin"
)

//llamar las funciones del controller

//crear estructura

//func newbookcontroler () bookcontroller {
//todos metodos vamos a ir pasandole los metodos de las funciones
//convierten la funciones publicas a metodos publicos
//}

/*------declaramos el puerto que escucha -------*/
const port = ":9000"

func main() {
	r := gin.Default()

	//va el repo/usecase/controller


	//van los endpoint
	/*------GETS -------*/
	//r.GET("/", Index)
	//r.GET("v1/listaInicial", GetListaInicial)
	r.GET("v1/items/:id", GetItemById)
	r.GET("v1/items", GetAllItems)

	/*------POST -------*/
	r.POST("v1/items", AddItem)

	/*------PUT-------*/
	//r.PUT("v1/items/:id", UpdateItem)

	/*------DELETE-------*/
	//r.DELETE("v1/items/:id", DeleteItem)


	/*------MSJ ESCUCHANDO PUERTO -------*/
	r.Run(port)
	log.Println("server listening to the port:", port)

	/*------MSJ DE ERROR -------*/
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalln(err)
	}
}