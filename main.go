package main

import (
	"encoding/json"
	"fmt"
	//"fmt"
	//"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	puerto      = ":8080"
	archivoJson = "./productos.json"
)

//repaso en casa clase asincronica
/*func main(){

	//usando el paquete gin
	router := gin.Default()
	router.GET("/hello-world", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Hello world",
		})

	})

	router.Run()
	//lo pruebo en http://localhost:8080/hello-word

}*/

//Clase sincronica 9
//persona es una estructura que define...
//punto 1
/*type Persona struct{
	Nombre string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad int `json:"edad"`
	Direccion string `json:"direccion"`
	Telefono string `json:"telefono"`
	Activo bool `json:"activo"`

}
func main(){
	//simula que le enviamos un post de tipo persona
	jsonPersona := `{
		"nombre":"juan",
		"apellido":"Perez",
		"edad":25,
		"direccion":"abdk",
		"telefono":"123456",
		"activo":true
	}`

	var persona Persona

	//de json a estructura (cuando recibo) el segundo parametro es donde quiero que me guarde el array  de byte que recibe
	err := json.Unmarshal([]byte(jsonPersona), &persona)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(persona)

	//punto 2
	personaResponse := Persona{
		Nombre: "Pedro",
		Apellido: "Lopez",
		Edad: 45,
		Direccion: "Av x",
		Telefono: "21552",
		Activo: true,

		}
	//instancio default gin
	router := gin.Default()
	//endpoint que siempre haya que tener para probar si
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"data": personaResponse,
		})
	})



	if err := router.Run(puerto); err != nil{
		log.Fatal(err)
	}

	//router.Run(puerto)
}*/

// Clase sincronica 9 - mesa de trabajo
// persona es una estructura que define...
// punto 1
type Producto struct {
	Id              int     `json:"id"`
	Nombre          string  `json:"nombre"`
	Precio          float64 `json:"precio"`
	Stock           int     `json:"stock"`
	Codigo          string  `json:"codigo"`
	Publicado       bool    `json:"publicado"`
	FechaDeCreacion string  `json:"fecha_de_creacion"`
}

func main() {
	//instancio default gin
	router := gin.Default()

	//endpoint que siempre haya que tener para probar si
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"data": "pong",
		})
	})

	router.GET("/readjson", func(c *gin.Context) {
		// Lee el contenido del archivo JSON
		fileContent, err := os.ReadFile(archivoJson)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer el archivo"})
			return
		}

		// Declara una variable para almacenar los datos decodificados
		var produc []Producto

		// Decodifica el contenido del archivo JSON en la estructura de Go
		err = json.Unmarshal([]byte(fileContent), &produc)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al decodificar el JSON"})
			return
		}

		// Devuelve los datos JSON como respuesta
		c.JSON(http.StatusOK, produc)
		fmt.Println(produc)
	})

	router.Run(puerto)

}
