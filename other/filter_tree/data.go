package main

type Person struct {
	Id       int64 `json:"id" db:"primary_key"`
	Name     string
	Age      int
	RidesCar bool
	HasToy   bool
}

var People = []Person{
	{
		Id:       1,
		Name:     "a",
		Age:      12,
		RidesCar: true,
		HasToy:   false,
	},
	{
		Id:       2,
		Name:     "b",
		Age:      8,
		RidesCar: false,
		HasToy:   true,
	},
	{
		Id:       3,
		Name:     "c",
		Age:      25,
		RidesCar: true,
		HasToy:   false,
	},
	{
		Id:       4,
		Name:     "d",
		Age:      20,
		RidesCar: false,
		HasToy:   true,
	},
	{
		Id:       5,
		Name:     "e",
		Age:      30,
		RidesCar: true,
		HasToy:   false,
	},
	{
		Id:       6,
		Name:     "f",
		Age:      18,
		RidesCar: true,
		HasToy:   true,
	},
	{
		Id:       7,
		Name:     "g",
		Age:      18,
		RidesCar: true,
		HasToy:   true,
	},
}
