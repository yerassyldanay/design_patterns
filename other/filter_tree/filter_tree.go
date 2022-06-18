package main

import "fmt"

// PersonFilter all conditional statements must follow the same structure
type PersonFilter interface {
	// Is gets a list of people
	// returns only those who went through filtering
	Is(p []Person) []Person
}

// MustBeAdult checks whether age >= 18
type MustBeAdult struct{}

func (m MustBeAdult) Is(p []Person) []Person {
	var result []Person
	for _, person := range p {
		if person.Age >= 18 {
			result = append(result, person)
		}
	}
	return result
}

// MustBeChild age < 18
type MustBeChild struct{}

func (m MustBeChild) Is(p []Person) []Person {
	var result []Person
	for _, person := range p {
		if person.Age < 18 {
			result = append(result, person)
		}
	}
	return result
}

type MustRideCar struct{}

func (m MustRideCar) Is(p []Person) []Person {
	var result []Person
	for _, person := range p {
		if person.RidingCar {
			result = append(result, person)
		}
	}
	return result
}

type MustHaveToy struct{}

func (m MustHaveToy) Is(p []Person) []Person {
	var result []Person
	for _, person := range p {
		if person.HasToy {
			result = append(result, person)
		}
	}
	return result
}

// AndStatement checks whether both statements are true
type AndStatement struct {
	FirstFilter  PersonFilter
	SecondFilter PersonFilter
}

func (a *AndStatement) Is(p []Person) []Person {
	var firsList = a.FirstFilter.Is(p)
	var firstListMap = make(map[int64]bool)
	for _, person := range firsList {
		firstListMap[person.Id] = true
	}

	// result
	var result []Person

	// after filtering with second filter
	secondList := a.SecondFilter.Is(p)
	for _, person := range secondList {
		// must be found also in the first filter
		if _, ok := firstListMap[person.Id]; ok {
			result = append(result, person)
		}
	}

	// ok
	return result
}

// OrStatement enables to put two conditions, where one is true
type OrStatement struct {
	FirstFilter  PersonFilter
	SecondFilter PersonFilter
}

func (o *OrStatement) Is(p []Person) []Person {
	// result
	var result []Person

	// first filter
	var firsList = o.FirstFilter.Is(p)
	var firstListMap = make(map[int64]*Person)

	// append to the list and save on map to avoid duplicating the person
	// inside the list
	for i := range firsList {
		result = append(result, firsList[i])
		firstListMap[firsList[i].Id] = &firsList[i]
	}

	// after filtering with the second filter
	secondList := o.SecondFilter.Is(p)
	for _, person := range secondList {
		if _, ok := firstListMap[person.Id]; ok {
			// already included
			continue
		}
		result = append(result, person)
	}

	// ok
	return result
}

func And(first, second PersonFilter) PersonFilter {
	return &AndStatement{
		FirstFilter:  first,
		SecondFilter: second,
	}
}

func Or(first, second PersonFilter) PersonFilter {
	return &OrStatement{
		FirstFilter:  first,
		SecondFilter: second,
	}
}

//
func main() {
	// lets say that we have accounts of people,
	// we wnat to filter out normal people (real)
	// person is normal / real if:
	// * adult and rides a car
	// * child and has a toy
	// fake:
	// * adult and has a toy
	// * child and rides a car

	// (adult && car) || (child && toy)
	newFilter := Or(And(MustBeAdult{}, MustRideCar{}), And(MustBeChild{}, MustHaveToy{}))
	for _, person := range newFilter.Is(People) {
		fmt.Printf("%#v is a real person \n", person)
	}
}
