package main

import "testing"


func TestIfGetPackHasAllCards(t *testing.T) {
	if len(cards)!=52 {
		t.Errorf("There are less than 52 cards")
	}
}

func TestIfNoCardIsReturnedTwice(t *testing.T) {
	card1 := getcard()
	card2 := getcard()
	if card1.val==card2.val{
		t.Errorf("Same card returned twice")
	}
}



