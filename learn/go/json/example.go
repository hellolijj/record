package main

import (
	"fmt"
	"encoding/json"
	"github.com/gpmgo/gopm/modules/log"
)

// 这里的json后面不能有空格
type Movie struct {
	Title  string
	Year   int `json:"released"`
	Color  bool `json:"color, omitempty"`
	Actors []string
}


var movies = []Movie{
	{Title: "calsdf", Year: 1942, Color: false, Actors: []string{"this", "is"}},
	{Title: "calsdfsdf", Year: 192, Color: true, Actors: []string{"thissaf", "iafas"}},
}

func main()  {

	if data, err := json.MarshalIndent(movies, "", "    "); err != nil {
		log.Fatal("json marshaling fail %s", err)
	} else {
		// data is a json string
		fmt.Printf("%s\n", data)
	}

	if data, err := json.Marshal(movies); err != nil {
		log.Fatal("json marshaling fial %s", err)
	} else {
		fmt.Printf("%T", data)
	}

	var s string
	s = "hello world"
	fmt.Printf("%T", s)
}

/*
json.MarshalIndent(data, "", "	")  // struct => format json string
json.Marshal(data)    // struct => json string
json.Unmarshall(data, &title)  // json => struct

*/