package main

import (
	"fmt"
	"strconv"
)

// 将某个8进制数转化为a-z排列的字符串，a代表1，z代表26,aa代表27，以此类推，aaa代表703。比如输入32，输出z。
func int8ToWord(input int) string {
	i, err := strconv.ParseInt(strconv.Itoa(input), 8, 0)
	if err != nil {
		panic(err)
	}

	target := ""
	t := int(i)

	for t >= 1 {

		tmp := 0
		if t > 26 {
			tmp = (t % 26) + 96
		} else {
			tmp = t + 96
		}

		t /= 26

		fmt.Println(t)

		target = target + string(rune(tmp))
	}

	return target
}

func main() {
	fmt.Print(int8ToWord(1277))
}
