package main

import "fmt"

type part struct {
	description string
	count int
}

func showInfo(p part) {
	fmt.Println("Description: ", p.description)
	fmt.Println("Count: ", p.count)
}

func minimumOrder(description string) part {
	var p part
	p.description = description
	p.count = 100
	return p
}

func main() {
	p := minimumOrder("Hex bolts")
	fmt.Println(p.description, p.count)
	/*
	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 100
	showInfo(bolts)
	/*
	var s1 subscriber

	fmt.Printf("%#v\n", s1)

	s1.name = "sunny"
	s1.rate = 5000
	s1.active = true

	fmt.Printf("%s\n", s1.name)
	fmt.Println(s1.rate)
	fmt.Println(s1.active)
	/*
	var subscriber struct {
		name string
		rate int
		active bool
	}
	fmt.Printf("%#v\n", subscriber)

	subscriber.name = "sunny"
	subscriber.rate = 5000
	subscriber.active = true

	fmt.Printf("%s\n", subscriber.name)
	fmt.Println(subscriber.rate)
	fmt.Println(subscriber.active)
	/*
	subscriber := map[string]float64{}
	subscriber ["name"] = "Kim"
	subscriber ["rate"] = 5000
	subscriber ["active"] = false

	subscriber := []string{}
	subscriber = append(subscriber, "kim")
	subscriber = append(subscriber, "han")
	subscriber = append(subscriber, "sunny")
	fmt.Println("------test------")
	*/
}
