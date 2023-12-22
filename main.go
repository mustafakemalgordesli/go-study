package main

// go mod init github.com/mustafakemalgordesli/hocourse
// go get github.com/twpayne/go-geom
// go mod tidy   // kullanılmayan dependency leri siler eklenmemiş olanları ekler

import (
	"fmt"

	"github.com/mustafakemalgordesli/hocourse/helpers"
	"github.com/mustafakemalgordesli/hocourse/helpers/difference"
)

func main() {
	fmt.Println(helpers.Sum(10, 20))
	fmt.Println(difference.Difference(20, 10))
}
