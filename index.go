package main

import (
	"fmt"
	"math"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler);
	http.ListenAndServe(":80", nil);
}

func Load() (string) {
	x := 0.0001
	for i := 0; i < 1000000; i++ {
	    x += math.Sqrt(x)
	}
	return Message();
}

func Message() (string){
	return "Code.education Rocks!";
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, Load());
}
