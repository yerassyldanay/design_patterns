package main

import (
	"fmt"
)

// and or not
// and(mustBe* ...)
// or(mustBe* ...)

// p Person(name string, age int, hasToy bool, ridesCar bool)
// and( and(cond1, cond2), or(cond2, cond3) )

type Filter interface {
	Is(p *Person) *Person
}

type MustBeAdult struct{}

func (f MustBeAdult) Is(p *Person) *Person {
	if p == nil || p.Age < 18 {
		return nil
	}
	return p
}

type MustBeChild struct{}

func (f MustBeChild) Is(p *Person) *Person {
	if p == nil || p.Age >= 18 {
		return nil
	}
	return p
}

type MustRideCar struct{}

func (f MustRideCar) Is(p *Person) *Person {
	if p == nil || !p.RidesCar {
		return nil
	}
	return p
}

type MustHaveToy struct{}

func (f MustHaveToy) Is(p *Person) *Person {
	if p == nil || !p.HasToy {
		return nil
	}
	return p
}

type AndStatement struct {
	filters []Filter
}

func (f AndStatement) Is(p *Person) *Person {
	if p == nil {
		return nil
	}

	for _, filter := range f.filters {
		if resp := filter.Is(p); resp == nil {
			return nil
		}
	}
	return p
}

func And(filters ...Filter) Filter {
	return AndStatement{filters: filters}
}

type OrStatement struct {
	filters []Filter
}

func (f OrStatement) Is(p *Person) *Person {
	if p == nil {
		return nil
	}

	for _, filter := range f.filters {
		if resp := filter.Is(p); resp != nil {
			return p
		}
	}
	return nil
}

func Or(filters ...Filter) Filter {
	return OrStatement{filters: filters}
}

func main() {
	// filtering real people from fake ones
	// to be a real person, person must meet following requirements
	// (adult && car) || (child && toy)
	newFilter := Or(And(MustBeAdult{}, MustRideCar{}), And(MustBeChild{}, MustHaveToy{}))
	for _, person := range People {
		resp := newFilter.Is(&person)
		if resp == nil {
			continue
		}
		fmt.Printf("%#v is a real person \n", resp)
	}
}
