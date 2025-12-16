package main

import (
	"fmt"
	"log"
	"net/http"
)
func main()  {

http.HandleFunc("/", hello)
fmt.Println("server running port:3000")
 log.Fatal(http.ListenAndServe(":3000" , nil))
	
}



func hello( w http.ResponseWriter , req *http.Request){
	fmt.Fprintf(w , "wansing sadfsf");
}