package internal
import "strconv"


func hexLookUp(digit uint8) rune {


}

func HexConvert(decimal uint8) (hexRepresentation uint16) {
    for pos, digit := range strconv.Itoa(int(decimal)) {
        hexRepresentation += 16^pos * int(digit) 
    }
}
