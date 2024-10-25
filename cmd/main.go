package main

import (
	"1hw/pkg/storage"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	console()

}

func console() {
	storage := storage.NewStorage()
	storage.Deseserialization()
	defer storage.Serialization()
	for {

		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		if input == "q" {
			fmt.Println("end")
			break
		}
		data := strings.Split(input, " ")
		var array []int
		if len(data) < 3 {
			fmt.Println(len(storage.GetList(data[1])))
		} else {
			for i := 2; i < len(data); i++ {
				a, err := strconv.Atoi(data[i])
				if err != nil {
					fmt.Println("err in parsing ", err)
					return
				}
				array = append(array, a)
			}
			switch data[0] {
			case "LPOP":
				fmt.Println(storage.LPOP(data[1], array))
			case "RPOP":
				fmt.Println(storage.RPOP(data[1], array))
			case "LPUSH":
				fmt.Println(storage.LPUSH(data[1], array))
			case "RPUSH":
				fmt.Println(storage.RPUSH(data[1], array))
			case "RADDTOSET":
				storage.RADDTOSET(data[1], array)
			case "LSET":
				fmt.Println(storage.LSET(data[1], array[0], array[1]))
			case "LGET":
				s, err := storage.LGET(data[1], array[0])
				if err {
					fmt.Println("index out of range")
				}
				fmt.Println(s)
			default:
				fmt.Println("wrong command")

			}
		}

	}
}
