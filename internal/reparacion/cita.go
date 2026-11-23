package reparacion

import "time"

const DuracionMaximaCita = 3 * time.Hour

type Cita struct {
	llamada int
	inicio  time.Time
	fin     time.Time
}
