package main

import "fmt"

func main() {
	arr := [...]string{"Hamed", "Maghsoud", "Mina", "Saeed"}
	fmt.Println(arr)
	fmt.Println("--------------------------------")
	for i, _ := range arr {
		fmt.Println(arr[i])
	}
	fmt.Println("--------------------------------")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	fmt.Println("--------------------------------")
	var i int
	for {
		if i == len(arr) {
			break
		}
		fmt.Println(arr[i])
		i++
	}
	fmt.Println("-------------MAPPPPPPPPPP-------------------")

	m := map[string]string{"NumberOne": "Hamed", "NumberTwo": "Maghsoud", "NumberThree": "Mina", "NumberFour": "Saeed", "EmptyVal": ""}
	// var m map[string]string
	// m = make(map[string]string, 4)
	// m["NumberOne"] = "Hamed"
	// m["NumberTwo"] = "Maghsoud"
	// m["NumberThree"] = "Mina"
	// m["NumberFour"] = "Saeed"
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("-------------MAPPPPPPPPPP-------------------")

	v := m["NumberOne"]
	fmt.Println(v)

	v, ok := m["NumberOneNadarim"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Not Found")
	}

	delete(m, "NumberOne")

	for k, _ := range m {
		fmt.Println(k, m[k])
	}

	fmt.Println("--------------SET----------------------")

	var s map[string]struct{}
	s = make(map[string]struct{}, 0)
	s["NumberOne"] = struct{}{}
	s["NumberTwo"] = struct{}{}
	s["NumberThree"] = struct{}{}
	s["NumberFour"] = struct{}{}
	s["NumberFive"] = struct{}{}

	for k, _ := range s {
		fmt.Println(k)
	}
	delete(s, "NumberOne")
	fmt.Println("------------------------------------")
	for k, _ := range s {
		fmt.Println(k)
	}
	fmt.Println("Finiiiiiish")
}
