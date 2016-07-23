package main

type Player struct{
	score1 int
	score2 int
}
func (this *Player) updatecard(val int){
	if val==1{
		this.score1+=val
		this.score2+=11
	}else{
		this.score1+=val
		this.score2+=val
	}
}