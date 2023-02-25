package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	name string
	age  uint
}

func TestFindIndex(t *testing.T) {
	str := New[string]()

	str.AddFirst("Sadik")
	str.AddFirst("Omar")
	str.AddFirst("Faruk")
	str.RemoveAt(1)

	assert.Equal(t, 1, str.FindIndex("Sadik"))
	assert.Equal(t, -1, str.FindIndex("Ahmad"))
	assert.Equal(t, 0, str.FindIndex("Faruk"))
}

func TestAddFirst(t *testing.T) {
	list := New[int]()
	list.AddFirst(10)
	list.AddFirst(20)
	list.AddFirst(30)
	f, _ := list.First()

	assert.Equal(t, 30, f)

	str := New[string]()

	_, err := str.First()
	assert.NotNil(t, err)

	str.AddFirst("Mango")
	str.AddFirst("Banana")
	str.AddFirst("Coconut")
	sf, _ := str.First()
	assert.Equal(t, "Coconut", sf)

	person := New[Person]()
	_, err = person.First()

	assert.NotNil(t, err)
	person.AddFirst(Person{name: "Omar Faruk", age: 20})
	person.AddFirst(Person{name: "Tanvir Raj", age: 25})
	p1, _ := person.First()
	assert.Equal(t, Person{name: "Tanvir Raj", age: 25}, p1)
	assert.Equal(t, uint(25), p1.age)
	assert.Equal(t, "Tanvir Raj", p1.name)
}

func TestAddLast(t *testing.T) {
	list := New[int]()
	_, err := list.Last()

	assert.NotNil(t, err)

	list.AddLast(30)
	list.AddLast(20)
	list.AddLast(10)

	l, err := list.Last()
	assert.Nil(t, err)
	assert.Equal(t, 10, l)

	str := New[string]()
	_, err = str.Last()
	assert.NotNil(t, err)

	str.AddLast("Coconut")
	str.AddLast("Mango")
	str.AddLast("Banana")

	ls, err := str.Last()
	assert.Nil(t, err)
	assert.Equal(t, "Banana", ls)

	person := New[Person]()
	_, err = person.Last()
	assert.Panics(t, func() { panic(err) })

	person.AddLast(Person{name: "Tanvir Raj", age: 25})
	person.AddLast(Person{name: "Omar Faruk", age: 20})

	p1, err := person.Last()
	assert.Nil(t, err)
	assert.Equal(t, Person{name: "Omar Faruk", age: 20}, p1)
	assert.Equal(t, uint(20), p1.age)
	assert.Equal(t, "Omar Faruk", p1.name)
}

func TestInsertAt(t *testing.T) {
	list := New[int]()
	list.InsertAt(30, 0)
	list.InsertAt(20, 1)
	list.InsertAt(10, 2)
	assert.Equal(t, 1, list.FindIndex(20))

	str := New[string]()
	str.InsertAt("Coconut", 0)
	str.InsertAt("Banana", 1)
	str.InsertAt("Mango", 1)
	str.InsertAt("Lichi", 3)
	str.InsertAt("Orange", str.Size())

	last, err := str.Last()
	assert.Nil(t, err)
	assert.Equal(t, "Orange", last)
	assert.Equal(t, 2, str.FindIndex("Banana"))
	assert.Panics(t, func() {
		str.InsertAt("Orange", 6)
	})

	person := New[Person]()
	person.InsertAt(Person{name: "Tanvir Raj", age: 25}, 0)
	person.InsertAt(Person{name: "Omar Faruk", age: 21}, 1)
	person.InsertAt(Person{name: "Sadik Ahmad", age: 20}, 1)
	p1, err := person.Last()
	assert.Nil(t, err)
	assert.Equal(t, Person{name: "Omar Faruk", age: 21}, p1)
	assert.Equal(t, uint(21), p1.age)
	assert.Equal(t, "Omar Faruk", p1.name)
}

func TestRemoveFirst(t *testing.T) {
	str := New[string]()

	err := str.RemoveAt(2)
	assert.NotNil(t, err)

	str.AddFirst("Omar")
	str.AddFirst("Faruk")
	str.RemoveFirst()

	s, err := str.First()
	assert.Nil(t, err)

	assert.Equal(t, "Omar", s)
}

func TestRemoveLast(t *testing.T) {
	str := New[string]()
	err := str.RemoveAt(2)
	assert.NotNil(t, err)
	str.AddFirst("Omar")
	str.AddFirst("Faruk")
	str.RemoveLast()
	s, err := str.Last()
	assert.Nil(t, err)

	assert.Equal(t, "Faruk", s)
}

func TestRemoveAt(t *testing.T) {
	str := New[string]()
	err := str.RemoveAt(2)
	assert.NotNil(t, err)
	str.AddFirst("Sadik")
	str.AddFirst("Faruk")
	str.AddFirst("Omar")
	str.RemoveAt(1)
	s, err := str.Last()
	assert.Nil(t, err)

	assert.Equal(t, "Sadik", s)
}
