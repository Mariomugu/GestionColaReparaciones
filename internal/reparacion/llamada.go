package reparacion

import "time"

type Llamada struct {
	urgencia       Urgencia
	momento        time.Time
	materialDesde  time.Time
	duracion       time.Duration
	disponibilidad []DisponibilidadDiaria
}
