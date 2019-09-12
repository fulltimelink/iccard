package iccard

import (
	"fmt"
	"math"
	"strconv"
)



func Ten2Eight(cardNo string) (string,error){
	cardNoF, err := strconv.ParseFloat(cardNo, 64)
	if nil != err{
		return "",err
	}
	part1 := math.Mod(cardNoF, 256*256)
	part2 := math.Mod(math.Floor(cardNoF/(256*256)), 256)*100000
	return fmt.Sprintf("%08.0f", part1+part2), nil
}
