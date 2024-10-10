package main

import (
	"1hw/storage"
	"fmt"
)



func main() {
	st := storage.NewStorage()
	st.Set("1", "1")
	st.Set("2", "string")

	fmt.Println(*st.Get("1"))
	if s := st.Get("3"); s != nil {
		fmt.Println(*s)
	} else {
		fmt.Println("not found")
	}

	fmt.Println(st.GetKind("1"))
	fmt.Println(st.GetKind("2"))


}
