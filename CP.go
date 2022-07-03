package main

import (
	"bufio"
	"fmt"
	"os"
)

//Input Output template
var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) {
	fmt.Fprintf(writer, f, a...)
}

func scanf(f string, a ...interface{}) {
	fmt.Fscanf(reader, f, a...)
}

func main() {
	defer writer.Flush()

	var mat [][]int
	var n int
	scanf("%d\n", &n)

	for i := 0; i < n; i++ {
		a := make([]int, n)
		for j := 0; j < n; j++ {
			scanf("%d", &a[j])
		}
		mat = append(mat, a)
		scanf("\n")
	}

	printf("%v\n", mat)
}
