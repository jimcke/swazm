
//There is no clear explication how the luck percentage is affecting the gameplay  so let's make an assumption on how this should work.
//Definition of being lucky means that a random number generated should be close to 0. If the player has zero luck than "gameLuck" is generated between 0-100 if player is all the time lucky than number generated is 0.
//If a player is 40% luck than the random number will be generated between (100 - 40).

package main

import (
	"fmt"
	"math/rand"
)

type Ability interface {
	attack(attacker *Warrior, defender *Warrior) float32
	defend(attacker *Warrior, defender *Warrior) float32
}

//--------- Critical Strike ---------
type CriticalStrike struct {}

func (a CriticalStrike) attack(attacker *Warrior, defender *Warrior) float32{
	fmt.Print("Is using strike with ")
	damage := attacker.strength - defender.defence
	chance := rand.Float32() * (100 - attacker.luck)
	if chance < 1 {
		fmt.Println("...triple damage")
		return damage * 3
	}else if chance < 10 {
		fmt.Println("...double damage")
		return damage * 2
	}
	fmt.Println("...simple damage")
	return damage
}

func (a CriticalStrike) defend(attacker *Warrior, defender *Warrior) float32{
	//only an offensive ability
	return 0
}

//--------- Resilience ---------
type Resilience struct {
	available bool
}

func (a Resilience) attack(attacker *Warrior, defender *Warrior) float32{
	return 0
}

func (a Resilience) defend(attacker *Warrior, defender *Warrior) float32{
	if a.available {
		chance := rand.Float32() * (100 - defender.luck)
		defendLevel := (attacker.strength - defender.health)/2
		if defendLevel > 0 && chance < 20{
			fmt.Println(defender.name + " got lucky and shield up by ", defendLevel)
			a.available = !a.available
			return defendLevel
		}
	}else {
		a.available = !a.available
	}
	return 0
}


//--------- Normal Strike ---------
type NormalStrike struct {}

func (a NormalStrike) attack(attacker *Warrior, defender *Warrior) float32{
	fmt.Println("Is using normal strike ")
	damage := attacker.strength - defender.defence
	return damage
}

func (a NormalStrike) defend(attacker *Warrior, defender *Warrior) float32{
	//only an offensive ability
	return 0
}

