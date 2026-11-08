package reparacion

import "time"

type Trabajo struct {
	id       int
	urgencia Urgencia
	llamada  time.Time
}
