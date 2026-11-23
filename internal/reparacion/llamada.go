package reparacion

import "time"

type Llamada struct {
	id             int
	urgencia       Urgencia
	momento        time.Time
	materialDesde  time.Time
	duracion       time.Duration
	citas          []Cita
	disponibilidad []DisponibilidadDiaria
}
