package main

import (
	"errors"
	"fmt"
	"time"
)

type Medicine struct {
	name            string
	manufacturer    string
	manufactureDate time.Time
	shelfLife       int
}

func (m *Medicine) SetName(name string) error {
	if name == "" {
		return errors.New("el nombre no puede estar vacío")
	}
	m.name = name
	return nil
}

func (m *Medicine) GetName() string {
	return m.name
}

func (m *Medicine) SetManufacturer(manufacturer string) error {
	if manufacturer == "" {
		return errors.New("el fabricante no puede estar vacío")
	}
	m.manufacturer = manufacturer
	return nil
}

func (m *Medicine) GetManufacturer() string {
	return m.manufacturer
}

func (m *Medicine) SetManufactureDate(date time.Time) error {
	if date.IsZero() {
		return errors.New("la fecha de fabricación debe ser una fecha válida")
	}
	m.manufactureDate = date
	return nil
}

func (m *Medicine) GetManufactureDate() time.Time {
	return m.manufactureDate
}

func (m *Medicine) SetShelfLife(shelfLife int) error {
	if shelfLife <= 0 {
		return errors.New("la vida útil debe ser un número positivo")
	}
	m.shelfLife = shelfLife
	return nil
}



func (m *Medicine) CalculateExpirationDate() time.Time {
	return m.manufactureDate.AddDate(0, m.shelfLife, 0)
}

type Tablet struct {
	Medicine
	dosePerTablet  float64
