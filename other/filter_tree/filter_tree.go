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

type mustBeAdult struct{}

func (f mustBeAdult) Is(p *Person) *Person {
	if p == nil || p.Age < 18 {
		return nil
	}
	return p
}

func MustBeAdult() mustBeAdult {
	return mustBeAdult{}
}

type mustBeChild struct{}

func (f mustBeChild) Is(p *Person) *Person {
	if p == nil || p.Age >= 18 {
		return nil
	}
	return p
}

func MustBeChild() mustBeChild {
	return mustBeChild{}
}

type mustRideCar struct{}

func (f mustRideCar) Is(p *Person) *Person {
	if p == nil || !p.RidesCar {
		return nil
	}
	return p
}

func MustRideCar() mustRideCar {
	return mustRideCar{}
}

type mustHaveToy struct{}

func (f mustHaveToy) Is(p *Person) *Person {
	if p == nil || !p.HasToy {
		return nil
	}
	return p
}

func MustHaveToy() mustHaveToy {
	return mustHaveToy{}
}

type and struct {
	filters []Filter
}

func (f and) Is(p *Person) *Person {
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
	return and{filters: filters}
}

type or struct {
	filters []Filter
}

func (f or) Is(p *Person) *Person {
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
	return or{filters: filters}
}

func main() {
	// filtering real people from fake ones
	// to be a real person, person must meet following requirements
	// (adult && car) || (child && toy)
	newFilter := Or(And(MustBeAdult(), MustRideCar()), And(MustBeChild(), MustHaveToy()))
	for _, person := range People {
		resp := newFilter.Is(&person)
		if resp == nil {
			continue
		}
		fmt.Printf("%#v is a real person \n", resp)
	}
}
