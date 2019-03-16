package main

import "fmt"

// int 转罗马数字
func intToRoman(num int) string {

	//RomanMap := map[int]string{
	//	1:    "I",
	//	5:    "V",
	//	10:   "X",
	//	50:   "L",
	//	100:  "C",
	//	500:  "D",
	//	1000: "M",
	//}

	target := ""

	for num >= 1 {
		switch true {
		case num >= 1000:
			indexM := num / 1000
			for indexM > 0 {
				target += "M"
				indexM--
			}
			num %= 1000
			break
		case num >= 900:
			target += "C" + "M"
			num %= 100
			break
		case num >= 500:
			indexC := (num - 500) / 100
			target += "D"
			for indexC > 0 {
				target += "C"
				indexC--
			}
			num %= 100
			break
		case num >= 400:
			target += "C" + "D"
			num %= 100
			break
		case num >= 100:
			indexC := num / 100
			for indexC > 0 {
				target += "C"
				indexC--
			}
			num %= 100
		case num >= 90:
			target += "X" + "C"
			num %= 10
			break
		case num >= 50:
			indexX := (num - 50) / 10
			target += "L"
			for indexX > 0 {
				target += "X"
				indexX--
			}
			num %= 10
			break
		case num >= 40:
			target += "X" + "L"
			num %= 10
			break
		case num >= 10:
			indexX := num / 10
			for indexX > 0 {
				target += "X"
				indexX--
			}
			num %= 10
		case num >= 9:
			target += "I" + "X"
			num /= 10
			break
		case num >= 5:
			indexI := num - 5
			target += "V"
			for indexI > 0 {
				target += "I"
				indexI--
			}
			num /= 10
			break
		case num >= 4:
			target += "I" + "V"
			num /= 10
			break
		case num >= 1:
			for num > 0 {
				target += "I"
				num--
			}
		}
	}

	return target
}

func main() {
	fmt.Println(intToRoman(1333))
}
