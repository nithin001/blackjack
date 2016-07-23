package main

import "testing"


func TestPlayerDealtCard(t *testing.T) {
	p := new(Player)
	p.updatecard(5)
	if p.score1 != 5 {
		t.Errorf("Player score should be 5 win but was %d", p.score1)
	}
}
func TestPlayerDealtCardWithMultipleValues(t *testing.T) {
	p := new(Player)
	p.updatecard(1)
	if p.score2 != 11 &&  p.score1 != 11{
		t.Errorf("Alteast One Player score should be have been 11")
	}
}