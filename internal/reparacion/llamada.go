package reparacion

import "time"

type Llamada struct {
	tipoTrabajo    TipoTrabajo
	urgencia       Urgencia
	momento        time.Time
	materialDesde  time.Time
	duracion       time.Duration
	disponibilidad []DisponibilidadDiaria
}
