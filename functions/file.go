package ascart

import (
	"fmt"
	"io/ioutil"
)

func ReadWholeFile(style string) []rune {
	contents, err := ioutil.ReadFile(style)
	if err != nil {
		fmt.Println(err.Error())
	}
	if len(string(contents)) == 0 {
		fmt.Println("Error! Empty parametre!")
		return nil
	}
	res := string(contents)
	var arr []rune
	for i := 0; i < len(res); i++ {
		arr = append(arr, rune(res[i]))
	}
	return arr
}
