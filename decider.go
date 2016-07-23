package main

const playerwins="player"
const dealerwins="dealer"
const anyonecanwin="anyone"

func decide(player,dealer Player) string{
	if player.score1==21 || player.score2 ==21 {
		return playerwins
	} else if dealer.score1 > 21 &&  dealer.score2  > 21 {
		return playerwins
	} else if player.score1 > 21 && player.score2 >21 {
		return dealerwins
	} else if isbothgreaterthan17(dealer) && isdealerlessthanplayer(dealer,player){
		return playerwins
	} else if isdealergreaterthanplayer(dealer,player){
		return dealerwins
	}
	return anyonecanwin
}


func isbothgreaterthan17(person Player) bool{
	if person.score1>=17 && person.score2>=17{
		return true
	}
	return false
}

func isonegreaterthan17(person Player) bool{
	if person.score1>=17 && person.score2>=17{
		return true
	}
	return false
}

func isdealerlessthanplayer(dealer,player Player) bool{
	if dealer.score1 < player.score1 && dealer.score2 < player.score2 && dealer.score2< player.score1 && dealer.score2 < player.score2{
		return true
	}
	return false
}

func isdealergreaterthanplayer(dealer,player Player) bool{
	if (dealer.score1>=17 && dealer.score1 > player.score1 && dealer.score1 > player.score2) ||  (dealer.score2>=17 && dealer.score2 > player.score1 && dealer.score2 > player.score2){
		return true
	}
	return false
}