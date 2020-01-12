package main

import (
	"fmt"
	"log"

	hw "github.com/AirArto/hw-3"
)

func init() {
	//
}

func main() {
	UnpackedArr := "oNe \\seven\\tWo, tHree, fOur,oNe \\seven\\tWo,   seven  tHree, fOur,//six nine, two, three-foUr/seven/, nine.../seven///eight//'.four'" +
		" nine.nine.nine.nine.nine//nine eight//. eight  seven six five oNe \\\\tWo, tHree, fOur/seven/six nine.,;',nine three'/seven///eight//>/seven/'.four123nine" +
		"//eight////eight//six five five;five//eight//eight//;<>0]1354six//.six['six///eight//five';five]]['six*five//eight,cat,cat,... dog some more random words lorem ipsum"
	data, err := hw.Do(&UnpackedArr)
	if err == nil {
		fmt.Println("Requested data:", data)
	} else {
		log.Println("Data displaying error: \n\t", err)
	}
}
