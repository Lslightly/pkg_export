package mypkg

import "fmt"

func ExportPrint(s string) {
	fmt.Printf("mypkg_%s\n", s)
}
