package main

import "fmt"

type ChainCheck func(p *Person) bool

func MustBeAdult(p *Person) bool {
	if p.Age >= 18 {
		return true
	}
	return false
}

func MustRideCar(p *Person) bool {
	return p.RidingCar
}

func MustNotHaveToy(p *Person) bool {
	return !p.HasToy
}

func Check(people []Person, checks ... ChainCheck) []Person {
	var result []Person
	for _, person := range people {
		var ok bool = true
		for _, check := range checks {
			if !check(&person) {
				ok = false
				break
			}
		}

		if ok {
			result = append(result, person)
		}
	}

	return result
}

func main() {
	// here we set a chain of checks
	// adult = meets age requirement + rides a car + has no toys
	for _, person := range Check(People, MustBeAdult, MustRideCar, MustNotHaveToy) {
		fmt.Printf("%#v is an adult person \n", person)
	}
}