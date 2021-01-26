package main

func main() {
	cards := createDeck()

	hand, _ := deal(cards, 3)
	hand.saveToFile("myObject")
}
