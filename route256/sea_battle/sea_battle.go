package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		var str string
		fmt.Fscanln(in, &str)
		strs := strings.Split(str, " ")
		ary := make([]int, len(strs))
		for i := range ary {
			ary[i], _ = strconv.Atoi(strs[i])
		}

		fmt.Println(ary)
	}
}
