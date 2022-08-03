package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

type Die struct {
	Id    int `json:"id"`
	Value int `json:"value"`
}

type DieString struct {
	Id    int    `json:"id"`
	Value string `json:"value"`
}

// Return the value of the die as a string of periods
func (d *Die) AsDots() *DieString {
	return &DieString{
		Id:    d.Id,
		Value: strings.Repeat(".", d.Value),
	}
}

// Return the value of the die as a word
func (d *Die) AsWord() *DieString {
	var v string
	switch d.Value {
	case 1:
		v = "one"
	case 2:
		v = "two"
	case 3:
		v = "three"
	case 4:
		v = "four"
	case 5:
		v = "five"
	case 6:
		v = "six"
	}
	return &DieString{
		Id:    d.Id,
		Value: v,
	}
}

// Return the value of the die as a floating point string
func (d *Die) AsFloat() *DieString {
	return &DieString{
		Id:    d.Id,
		Value: fmt.Sprintf("%.1f", float32(d.Value)),
	}
}

// Randomly choose a new value for the die
func (d *Die) Roll() {
	rand.Seed(time.Now().UnixNano())
	d.Value = rand.Intn(6) + 1
	log.Printf("Die #%d was rolled and it now shows %d\n", d.Id, d.Value)
}
