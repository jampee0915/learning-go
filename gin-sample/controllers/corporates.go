package controllers

import "../models"

type Corporates struct {
}

func NewCorporates() Corporates {
	return Corporates{}
}

func (c Corporates) Get(n int64) interface{} {
	repo := models.NewCorporatesRepository()
	corporate := repo.GetByCorporateId(n)
	return corporate
}
