package tickets

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	id           string
	nombre       string
	email        string
	pais_destino string
	hora         string
	precio       string
}

var Tickets []Ticket

var (
	Madrugada = "madrugada"
	Mañana    = "mañana"
	Tarde     = "tarde"
	Noche     = "noche"
)

// ejemplo 1
func GetTotalTickets(destination string) (total int, err error) {
	_, err = existCountry(destination)
	if err != nil {
		return
	}
	for _, ticket := range Tickets {
		if ticket.pais_destino == destination {
			total++
		}
	}
	return
}

func existCountry(destination string) (b bool, err error) {
	for _, v := range Tickets {
		if v.pais_destino == destination {
			return true, nil
		}
	}
	err = errors.New("pais inexistente")
	return
}

// ejemplo 2
func GetMornings(time string) (total int, err error) {
	/*for i, ticket := range Tickets {
		fmt.Println(i, strings.Split(ticket.hora, ":")[0])
	}*/
	switch time {
	case Madrugada:
		for _, ticket := range Tickets {
			//Conversion de string a int
			num, er := strconv.Atoi(strings.Split(ticket.hora, ":")[0])
			if er != nil {
				fmt.Println("Error during conversion")
				err = er
				return
			}
			//Validar Madrugada
			if num >= 0 && num <= 6 {
				total++
			}
		}
		return
	case Mañana:
		for _, ticket := range Tickets {
			//Conversion de string a int
			num, er := strconv.Atoi(strings.Split(ticket.hora, ":")[0])
			if er != nil {
				fmt.Println("Error during conversion")
				err = er
				return
			}
			//Validar Mañana
			if num >= 7 && num <= 12 {
				total++
			}
		}
		return
	case Tarde:
		for _, ticket := range Tickets {
			//Conversion de string a int
			num, er := strconv.Atoi(strings.Split(ticket.hora, ":")[0])
			if er != nil {
				fmt.Println("Error during conversion")
				err = er
				return
			}
			//Validar Tarde
			if num >= 13 && num <= 19 {
				total++
			}
		}
		return
	case Noche:
		for _, ticket := range Tickets {
			//Conversion de string a int
			num, er := strconv.Atoi(strings.Split(ticket.hora, ":")[0])
			if er != nil {
				fmt.Println("Error during conversion")
				err = er
				return
			}
			//Validar Noche
			if num >= 20 && num <= 23 {
				total++
			}
		}
		return
	}
	return
}

// ejemplo 3
func AverageDestination(destination string) (promedio float64, err error) {
	pasajeros, e := GetTotalTickets(destination)
	if e != nil {
		fmt.Println(e)
		return 0, err
	}
	promedio = float64(pasajeros) / float64(len(Tickets)) * 100
	return
}

func CargarTickets() {
	//Open File
	CsvFile, err := os.Open("tickets.csv")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}

	//Parse/Process
	r := csv.NewReader(CsvFile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic("Error")
		}
		ticket := Ticket{record[0], record[1], record[2], record[3], record[4], record[5]}
		Tickets = append(Tickets, ticket)
	}
}
