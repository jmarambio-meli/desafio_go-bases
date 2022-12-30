package tickets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalTickets(t *testing.T) {
	pais := "Inexistente"

	valor_esperado := 0

	respuesta, _ := GetTotalTickets(pais)
	assert.Equal(t, valor_esperado, respuesta, "Obtener Tickets incorrecto")

}

func TestAverageDestination(t *testing.T) {
	pais := "Brazil"

	valor_esperado := float64(0)

	respuesta, _ := AverageDestination(pais)
	assert.Equal(t, valor_esperado, respuesta, "Obtener promedio incorrecto")

}
