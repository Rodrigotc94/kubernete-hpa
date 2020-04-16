package main

import (
	"fmt"
	"math"
	"net/http"
)

func main() {
	calcular()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, greeting("Code.education Rocks!"))
		// fmt.Fprint(w, calcular())
	})
	fmt.Println(http.ListenAndServe(":8000", nil))
}

func greeting(msg string) string {
	return "<b>" + msg + "</b>"
}

func calcular() float64 {
	x := 0.0001
	for i := 0; i < 1000000; i++ {
		x += math.Sqrt(x)
	}
	return x
}
