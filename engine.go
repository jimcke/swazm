package main

import (
	"fmt"
	"os"
)

type Gameplay struct {
	characters []Warrior
}

func newGame(gameplay *Gameplay) {

	fmt.Println("New game has started")
	hero := NewCharacter("Jon Snow", "Stark", true)
	enemy := NewCharacter("Ramsey Bolton", "Bolton", false)
	fmt.Println("")

	gameplay.characters = append(gameplay.characters, hero, enemy)

	fmt.Println("Hero: " + hero.name + "  Health: ", hero.health, "Luck: ", hero.luck)
	fmt.Println("Enemy: " + enemy.name + "  Health: ", enemy.health, "Luck: ", enemy.luck)

	attacker, defender := firstAttacker(&hero, &enemy)
	for i := 1;  i<=20; i++ {
		fmt.Println("Round " , i)
		attacker, defender = turn(attacker, defender)

		if attacker.health <= 0 {
			fmt.Println( attacker.name + " is dead.")
			os.Exit(0)
		} else if defender.health <= 0 {
			fmt.Println(defender.name + " is dead.")
			os.Exit(0)
		}
	}
}

func turn(warrior1 *Warrior, warrior2 *Warrior)  (*Warrior, *Warrior){
	fmt.Println("Turn for " + warrior1.name)
	for _, ability := range warrior1.abilities {
		abilityDamage := ability.(Ability).attack(warrior1, warrior2)
		defendLevel := float32(0.00)
		for _, defensive := range warrior2.abilities {
			defendLevel +=defensive.(Ability).defend(warrior1, warrior2)
		}
		damage := abilityDamage - defendLevel
		if damage > 0 {
			warrior2.health -= damage
		}
	}

	fmt.Println( warrior1.name + "    Health: ", warrior1.health)
	fmt.Println( warrior2.name + "    Health: ", warrior2.health)

	fmt.Println("Turn completed. Changing roles..")
	fmt.Println()
	return warrior2, warrior1
}

//Who is the first one. If both speed and luck are the same then we default to warrior1
func firstAttacker(warrior1 *Warrior, warrior2 *Warrior) (*Warrior, *Warrior) {
	if warrior1.speed < warrior2.speed {
		return warrior2, warrior1
	} else if warrior1.speed > warrior2.speed {
		return warrior1, warrior2
	} else {
		if warrior1.luck < warrior2.luck {
			return warrior2, warrior1
		}
		return warrior1, warrior2
	}
}
