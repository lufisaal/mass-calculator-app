package main

import (
	"fmt"
	"net/http"
	"strconv"
	"os"
	"math"
)
// objects
type Mass struct {
	Density float64
}

type Sphere struct {
	Mass
}

type Cube struct {
	Mass
}

type MassVolume interface {
	density() float64
	volume(dimension float64) float64
}

// density
func (s Sphere) density() float64 {
	return s.Mass.Density
}

func (c Cube) density() float64 {
	return c.Mass.Density
}

// volume calculation
func (s Sphere) volume(radius float64) float64 {
	return (4.0 / 3.0) * math.Pi * math.Pow(radius, 3)
}

func (c Cube) volume(sideLength float64) float64 {
	return math.Pow(sideLength, 3)
}

func Handler(massVolume MassVolume) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if dimension, err := strconv.ParseFloat(r.URL.Query().Get("dimension"), 64); err == nil {
			weight := massVolume.density() * massVolume.volume(dimension)
			w.Write([]byte(fmt.Sprintf("%.2f", math.Round(weight*100)/100)))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
	}
}

func main() {
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	aluminiumSphere := Sphere{Mass{Density: 2.710}}
	ironCube := Cube{Mass{Density: 7.874}}
	http.HandleFunc("/aluminium/sphere", Handler(aluminiumSphere))
	http.HandleFunc("/iron/cube", Handler(ironCube))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}