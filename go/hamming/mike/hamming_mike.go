package hamming

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(Distance("zlsjnebg93hbhd", "xlsgnebg93hbhd"))
}

func Distance(a, b string) (int, error) {
	distance := 0

	if len(a) != len(b) {
		return 2, errors.New("Strings not equal")
	}

	for i, v := range b {
		if a[i] != b[i] {
			distance++
			fmt.Println(i, string(v))
		}
	}
	return distance, nil
}
