package main

import "testing"

func TestPlayerIsWinnerIfPlayerHas21(t *testing.T) {
	player:=Player{score1:21,score2:12}
	dealer:=Player{score1:0,score2:0}
	result := decide(player,dealer)
	if result != playerwins {
		t.Errorf("Player should win but %s won", result)
	}
}

func TestPlayerIsWinnerIfDealerHasGreaterThan21(t *testing.T) {
	player:=Player{score1:15,score2:15}
	dealer:=Player{score1:22,score2:22}
	result := decide(player,dealer)
	if result != playerwins {
		t.Errorf("Player should win but %s won", result)
	}
}


func TestDealerIsWinnerIfPlayerHasGreaterThan21(t *testing.T) {
	player:=Player{score1:22,score2:22}
	dealer:=Player{score1:0,score2:0}
	result := decide(player,dealer)
	if result != dealerwins {
		t.Errorf("Dealer should win but %s won", result)
	}
}

func TestPlayerIsWinnerIfPlayerHasLessThan21AndDealerHasGreaterThan17AndLesserThanPlayer(t *testing.T) {
	player:=Player{score1:19,score2:19}
	dealer:=Player{score1:17,score2:17}
	result := decide(player,dealer)
	if result != playerwins {
		t.Errorf("Player should win but %s won", result)
	}
}



func TestDealerIsWinnerIfDealerHasGreaterThanPlayerButLesserThanOrEqualTo21(t *testing.T) {
	player:=Player{score1:19,score2:19}
	dealer:=Player{score1:20,score2:20}
	result := decide(player,dealer)
	if result != dealerwins {
		t.Errorf("Dealer should win but %s won", result)
	}
}


