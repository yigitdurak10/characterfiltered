package main

import (
	"fmt"
)

func cencored(s string) any {
	sum := 0
	cencoredList := "@#&%*"
	for i := 0; i <= len(s)-1; i++ {
		for x := 0; x <= len(cencoredList)-1; x++ {
			if string(cencoredList[x]) == string(s[i]) {
				sum = 1
				break
			}
		}

	}
	if sum > 0 {
		fmt.Println("Yasaklı karakter algılandı")
	} else {
		fmt.Println("Başarılı")
	}
	return 0
}

func main() {
	cencored("kyd")
}
