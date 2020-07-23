hola!
el objetivo del proyecto es tomar un archivo plano el cual contiene el historial de una aplicacion de geolocalizacion y extraer las coordenadas de ubicacion del usuario, y codificarlas a traves de codigo de polilinea.
los pasos a seguir para la elaboracion de este proyecto son los siguientes:
1. identificar el archivo del cual necesitamos extraer la informacion, el cual es un archivo de texto plano que tiene una estructura determinada y dentro de la cual se encuentra informacion de la aplicacion y resultados arrojados por la misma de diferentes ubicaciones donde estuvo el usuario
2. el cliente necesita que tomemos un archivo y leamos el contenido, para encontrar las coordenadas de ubicacion del usuario.
3.  para ello debemos abrir el archivo dentro de nuestro codigo, y solucionar posibles errores en la apertura del mismo, lo cual hicimos de la siguiente forma:

    file, err := os.Open("example.log")
	if err != nil {
		log.Fatalf("error al abrir archivo:%s", err)
	}

4. leer linea por linea el archivo con el fin de hallar las coordenadas coontenidas en el archivo.
5. una vez encontrada la informacion, extraerla del archivo.
6. almacenar la informacion encontrada.
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

7. codificar la informacion encontrada, para ello utilizaremos codigo polilinea desarrollado por otros programadores, el cual encontraremos en la plataforma GitHub.

    fmt.Println("%s\n", polyline.EncodeCoords(coords))
    
8. mostrar la informacion codificada.
9. enviar la informacion.
en la digitacion del codigo de este proyecto encontraremos los comentarios de los pasos anteriormente mencionados.


