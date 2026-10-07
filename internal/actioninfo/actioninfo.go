package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, v := range dataset {
		err := dp.Parse(v)
		if err != nil {
			log.Fatal(err)
			continue
		}
		result, err := dp.ActionInfo()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}
}
