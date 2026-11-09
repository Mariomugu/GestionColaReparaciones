package reparacion

import "errors"

var ErrCitasSolapadas = errors.New("Haquin no puede estar en dos trabajos a la vez")
var ErrUrgenciaInexistente = errors.New("la urgencia no es urgente, normal ni puede esperar")
