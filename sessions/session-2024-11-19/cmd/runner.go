package main

import "fmt"

func main() {
	//collections()
	switchs()
}

func switchs() {
	arr := [...]string{"Hamed", "Maghsoud", "Mina", "Saeed"}
	fmt.Println(arr)
	fmt.Println("-------------------------Switch--------------------------------")
	for i, v := range arr {
		switch {
		case i == 0:
			fmt.Printf("Zero Item %d with val %v\n", i, v)
		case i%2 == 0:
			fmt.Printf("Even Item %d with val %v\n", i, v)
		default:
			fmt.Printf("Odd Item %d with val %v\n", i, v)
		}
	}
}

func collections() {
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

type animal interface {
	move()
}
type snake struct {
	movement string
}
type bird struct {
	movement string
}

type mammal struct {
	movement string
}

func (o snake) move() {
	fmt.Println(o.movement)
}

func (o snake) crawl() {
	fmt.Println("crawling")
}
func (o bird) move() {
	fmt.Println(o.movement)
}

func (o bird) fly() {
	fmt.Println("fly")
}
func (o mammal) move() {
	fmt.Println(o.movement)
}

func (o mammal) walk() {
	fmt.Println("walk")
}
func printMyType(v animal) {
	// please instantiate of each animal. snake, mammal and bird
	// using switch case please find the animal type and print the type
	// based on type, please call specific method of the object, == fly(),walk() and crawl()
	// thanks a lot
}
