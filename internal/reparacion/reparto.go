package reparacion

import "time"

type Reparto struct {
	llamadas map[int]Llamada
	citas    map[time.Time]Cita
}

type SinHuecoError struct {
	Llamada Llamada
}

func (e SinHuecoError) Error() string {
	return "no hay hueco libre para dar cita a la llamada"
}
