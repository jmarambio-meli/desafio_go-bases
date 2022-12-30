package main

import (
	"errors"
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	//total, err := tickets.GetTotalTickets("Brazil")
	defer func() {
		fmt.Println("Menu finalizado")

		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()
	tickets.CargarTickets()
	var opcion string
	var pais string
	var tiempo string
	var loop bool = true
	for loop {
		println("Ingrese '1' para hallar cuantas personas viajaron a un pais en especifico")
		println("Ingrese '2' para hallar cuantas personas viajaron en un periodo en especifico")
		println("Ingrese '3' para hallar el promedio de personas que viajaron a un pais en especifico")
		println("Ingrese S Para salir del Menu")
		fmt.Printf("\n")
		fmt.Scanln(&opcion)
		switch opcion {
		case "1":
			fmt.Println("Porfavor ingrese un pais (Recuerde que la primeta letra es con Mayuscula)")
			fmt.Scan(&pais)
			total, err := tickets.GetTotalTickets(pais)
			if err != nil {
				panic(err)
			}
			fmt.Printf("La cantidad de personas que viajan a %s es de %d\n", pais, total)
		case "2":
			println("Ingrese '1' para madrugada")
			println("Ingrese '2' para mañana")
			println("Ingrese '3' para tarde")
			println("Ingrese '4' para noche")
			fmt.Scan(&tiempo)
			switch tiempo {
			case "1":
				total, _ := tickets.GetMornings("madrugada")
				fmt.Println("La cantidad de personas que viajaron por la madrugada son: ", total)
			case "2":
				total, _ := tickets.GetMornings("mañana")
				fmt.Println("La cantidad de personas que viajaron por la mañana son: ", total)
			case "3":
				total, _ := tickets.GetMornings("tarde")
				fmt.Println("La cantidad de personas que viajaron por la tarde son: ", total)
			case "4":
				total, _ := tickets.GetMornings("noche")
				fmt.Println("La cantidad de personas que viajaron por la noche son: ", total)
			default:
				fmt.Println(errors.New("opcion incorrecta del periodo"))
			}
		case "3":
			fmt.Println("Porfavor ingrese un pais (Recuerde que la primeta letra es con Mayuscula)")
			fmt.Scan(&pais)
			average, err := tickets.AverageDestination(pais)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("El promedio de personas es: ", average)
		case "S":
			loop = false
		default:
			fmt.Println(errors.New("opcion incorrecta, ingrese nuevamente"))
		}
	}
}
