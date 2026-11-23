package reparacion

import "time"

type Reparto struct {
	llamadas map[int]Llamada
	citas    map[time.Time]Cita
}
