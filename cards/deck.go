package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
)
// create a new type of 'deck'
// which is a slice of string

type deck []string
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, num int) (deck, deck) {
	return d[:num], d[num:]
}

func (d deck) toString() string {
	return strings.Join(d,  ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	} 
	// convert []byte to string
	s := strings.Split(string(bs),",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(int64(time.Now().UnixNano()))
	r := rand.New(source)
	for i := range(d) {
		np := r.Intn(len(d))
		d[i], d[np] = d[np], d[i]
	}
}