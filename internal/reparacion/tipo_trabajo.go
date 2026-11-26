package reparacion

type TipoTrabajo int

const (
	AiresAcondicionados TipoTrabajo = iota + 1
	Fontaneria
	Electricidad
	CamarasFrigorificas
	Pintura
	Comunidades
	Maquinas
	Otro
)
