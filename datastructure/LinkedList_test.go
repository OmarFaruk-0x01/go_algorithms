package datastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	name string
	age  uint
}

func TestFindIndex(t *testing.T) {
	str := LinkedList[string]{}

	str.AddFirst("Sadik")
	str.AddFirst("Omar")
	str.AddFirst("Faruk")
	str.RemoveAt(1)

	assert.Equal(t, 1, str.FindIndex("Sadik"))
	assert.Equal(t, -1, str.FindIndex("Ahmad"))
	assert.Equal(t, 0, str.FindIndex("Faruk"))
}

func TestAddFirst(t *testing.T) {
	list := LinkedList[int]{}
	list.AddFirst(10)
	list.AddFirst(20)
	list.AddFirst(30)
	assert.Equal(t, 30, list.First())

	str := LinkedList[string]{}
	str.AddFirst("Mango")
	str.AddFirst("Banana")
	str.AddFirst("Coconut")
	assert.Equal(t, "Coconut", str.First())

	person := LinkedList[Person]{}
	person.AddFirst(Person{name: "Omar Faruk", age: 20})
	person.AddFirst(Person{name: "Tanvir Raj", age: 25})
	assert.Equal(t, Person{name: "Tanvir Raj", age: 25}, person.First())
	assert.Equal(t, uint(25), person.First().age)
	assert.Equal(t, "Tanvir Raj", person.First().name)
}

func TestAddLast(t *testing.T) {
	list := LinkedList[int]{}
	list.AddLast(30)
	list.AddLast(20)
	list.AddLast(10)
	assert.Equal(t, 10, list.Last())

	str := LinkedList[string]{}
	str.AddLast("Coconut")
	str.AddLast("Mango")
	str.AddLast("Banana")
	assert.Equal(t, "Banana", str.Last())

	person := LinkedList[Person]{}
	person.AddLast(Person{name: "Tanvir Raj", age: 25})
	person.AddLast(Person{name: "Omar Faruk", age: 20})
	assert.Equal(t, Person{name: "Omar Faruk", age: 20}, person.Last())
	assert.Equal(t, uint(20), person.Last().age)
	assert.Equal(t, "Omar Faruk", person.Last().name)
}

func TestInsertAt(t *testing.T) {
	list := LinkedList[int]{}
	list.InsertAt(30, 0)
	list.InsertAt(20, 1)
	list.InsertAt(10, 2)
	assert.Equal(t, 1, list.FindIndex(20))

	str := LinkedList[string]{}
	str.InsertAt("Coconut", 0)
	str.InsertAt("Banana", 1)
	str.InsertAt("Mango", 1)
	str.InsertAt("Lichi", 3)
	str.InsertAt("Orange", str.Size())

	assert.Equal(t, "Orange", str.Last())
	assert.Equal(t, 2, str.FindIndex("Banana"))
	assert.Panics(t, func() {
		str.InsertAt("Orange", 6)
	})

	person := LinkedList[Person]{}
	person.InsertAt(Person{name: "Tanvir Raj", age: 25}, 0)
	person.InsertAt(Person{name: "Omar Faruk", age: 21}, 1)
	person.InsertAt(Person{name: "Sadik Ahmad", age: 20}, 1)
	assert.Equal(t, Person{name: "Omar Faruk", age: 21}, person.Last())
	assert.Equal(t, uint(21), person.Last().age)
	assert.Equal(t, "Omar Faruk", person.Last().name)
}

func TestRemoveFirst(t *testing.T) {
	str := LinkedList[string]{}

	assert.Panics(t, func() {
		str.RemoveFirst()
	})

	str.AddFirst("Omar")
	str.AddFirst("Faruk")
	str.RemoveFirst()

	assert.Equal(t, "Omar", str.First())
}

func TestRemoveLast(t *testing.T) {
	str := LinkedList[string]{}

	assert.Panics(t, func() {
		str.RemoveLast()
	})

	str.AddFirst("Omar")
	str.AddFirst("Faruk")
	str.RemoveLast()

	assert.Equal(t, "Faruk", str.First())
}

func TestRemoveAt(t *testing.T) {
	str := LinkedList[string]{}

	assert.Panics(t, func() {
		str.RemoveAt(2)
	})

	str.AddFirst("Sadik")
	str.AddFirst("Faruk")
	str.AddFirst("Omar")
	str.RemoveAt(1)

	assert.Equal(t, "Sadik", str.Last())
}
