package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"https://github.com/twpayne/go-polyline"
)

func main() {
	//paso 1,2,3 identificamos, traemos y abrimos archivo log
	//y manejamos los posibles errores de apertura
	file, err := os.Open("example.log")
	if err != nil {
		log.Fatalf("error al abrir archivo:%s", err)
	}
	//4.leer linea por linea el archivo y manejar posibles
	//errores de lectura

	fileScanner := bufio.NewScanner(file)
	var coordenada [2][400]float64
	incrementador := 1

	for fileScanner.Scan() {
		string := fileScanner.Text()
		pos := strings.Index(string, "location arrived")
		if pos != -1 {
			x, _ := strconv.ParseFloat(string[78:89], 64)
			y, _ := strconv.ParseFloat(string[90:102], 64)

			fmt.Print(x, y)

			incrementador++

		}
		fmt.Println(coordenada[0][256])
		fmt.Println(incrementador - 1)
		fmt.Println("%s\n", polyline.EncodeCoords(coords))

	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("error al abrir archivo: %s", err)

	}

	file.Close()
	//5.encontrar la informacion solicitada y extraerla
	//6. y 7. almacenar la informacion encontrada y codificarla
	//8.enviar informacion

}
