package main

import "fmt"

func main() {
	b := [7]int{1, 5, 17, 19, 29, 55, 56}
	a := [5]int{2, 4, 16, 28, 30}

	c := [len(a) + len(b)]int{}
	i, j, k := 0, 0, 0
	for {
		if i >= len(a){
			c[k] = b[j]
			for z := j;z<len(b)-1;z++{
				k++
				j++
				c[k] = b[j]

			}
			break
		} else if j >= len(b){
			c[k] = a[i]
			for z := i;z<len(a)-1;z++{
				k++
				i++
				c[k] = a[i]
			}
			break
		}
		if a[i] <= b[j] {
			c[k] = a[i]
			i++
			k++
		} else {
			c[k] = b[j]
			j++
			k++
		}
	}

	fmt.Println(c)

}
