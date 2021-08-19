package main

type Person struct {
	Id int64 `json:"id" db:"primary_key"`
	Name string
	Age int
	RidingCar bool
	HasToy bool
}

var People = []Person{
	{
		Id:        1,
		Name:      "a",
		Age:       12,
		RidingCar: true,
		HasToy:    false,
	},
	{
		Id:        2,
		Name:      "b",
		Age:       8,
		RidingCar: false,
		HasToy:    true,
	},
	{
		Id:        3,
		Name:      "c",
		Age:       25,
		RidingCar: true,
		HasToy:    false,
	},
	{
		Id:        4,
		Name:      "d",
		Age:       20,
		RidingCar: false,
		HasToy:    true,
	},
	{
		Id:        5,
		Name:      "e",
		Age:       30,
		RidingCar: true,
		HasToy:    false,
	},
}
