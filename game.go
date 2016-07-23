package main

import (
	"math/rand"
	"fmt"
	"time"
	"bufio"
	"os"
)
var player=Player{score1:0,score2:0}
var dealer=Player{score1:0,score2:0}
var cards= getpack()
type Card struct{
	val int
	symbol string
	suite string
	dealtalready bool
}

func main(){
	var card=getcard();
	fmt.Println("Player dealt: %s",card.symbol);
	player.updatecard(card.val)
	card=getcard();
	fmt.Println("Player dealt: %s",card.symbol);
	player.updatecard(card.val)
	card=getcard();
	fmt.Println("Dealer dealt: %s",card.symbol);
	dealer.updatecard(card.val)
	card=getcard();
	fmt.Println("Dealer dealt: %s",card.symbol);
	dealer.updatecard(card.val)

	for{
		reader := bufio.NewReader(os.Stdin)
	    fmt.Println("Do you want to hit?")
	    text, _ := reader.ReadString('\n');
	    if text == "no\n" {
	    	break
	    }
	    card=getcard();
		fmt.Println("Player dealt: %s",card.symbol);
	    player.updatecard(card.val)
	    if decide(player,dealer)==dealerwins{
	    	break
	    }
	}
	for{
		if decide(player,dealer)!=anyonecanwin{
			break
		}
		card=getcard();
		fmt.Println("Dealer dealt: %s",card.symbol);
		dealer.updatecard(card.val)
		
	}
	fmt.Println("%s wins",decide(player,dealer))
}

func gencards(suite string) []Card{
	cards:=[]Card{}
	var symbol string
	for i := 1; i <= 10; i++ {
		symbol= fmt.Sprintf("%d",i)
		card := Card{i,symbol,suite,false}
	    cards = append(cards,card)
	} 
	cards = append(cards,Card{10,"J",suite,false})
	cards = append(cards,Card{10,"Q",suite,false})
	cards = append(cards,Card{10,"K",suite,false})
	return cards
}
func getpack() []Card{
	cards:=[]Card{}
	cards = Append(cards, gencards("D"))
	cards = Append(cards, gencards("S"))
	cards = Append(cards, gencards("C"))
	cards = Append(cards, gencards("H"))
	return cards
}
func getcard() Card{
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

	for{
		card:=&cards[r1.Intn(52)]
		if(card.dealtalready==false){
			card.dealtalready=true
			return *card
		}
	}
}
func Append(slice []Card, items []Card) []Card {
    for _, item := range items {
        slice = append(slice, item)
    }
    return slice
}

