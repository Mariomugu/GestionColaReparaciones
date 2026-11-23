package reparacion

type Urgencia int

const (
	PuedeEsperar Urgencia = iota + 1
	Normal
	Urgente
)
