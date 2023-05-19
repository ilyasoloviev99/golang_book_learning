package main

import "fmt"

type Animal interface {
	GetSpeed() float64
}

type Cat struct {
	meow  string
	speed float64
}

type Dog struct {
	speed float64
	bark  string
}

func (cat Cat) GetSpeed() float64 {
	return cat.speed
}

func (dog Dog) GetSpeed() float64 {
	return dog.speed
}

func averageSpeed(animals []Animal) float64 {
	sum := float64(0)
	for _, animal := range animals {
		sum += animal.GetSpeed()
	}
	return sum / float64(len(animals))
}

func main() {
	dog := Dog{speed: 20, bark: "Гав"}
	cat := Cat{speed: 37, meow: "Мяу"}
	animals := []Animal{cat, dog}

	fmt.Println(averageSpeed(animals))

	newAnimal := animals[1]
	newCat, ok := newAnimal.(Cat)
	for _, animal := range animals {
		switch animal.(type) {
		case Dog:
			fmt.Println("Привет, я собака")
		case Cat:
			fmt.Println("Привет, я кот")
		default:
			fmt.Println("Привет, я пока не определился")
		}
	}

	if !ok {
		fmt.Println("Ты че дурной, я не кот")
		return
	}
	fmt.Println(newCat.meow)
}
