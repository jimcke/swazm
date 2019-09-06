package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

// Spawn limit values for heroes and enemies
const (
	heroHealthMin   = 70.00
	heroHealthMax   = 100.00
	heroStrengthMin = 70.00
	heroStrengthMax = 80.00
	heroDefenceMin  = 45.00
	heroDefenceMax  = 55.00
	heroSpeedMin    = 40.00
	heroSpeedMax    = 50.00
	heroLuckMin     = 10.00
	heroLuckMax     = 20.00

	enemyHealthMin   = 60.00
	enemyHealthMax   = 90.00
	enemyStrengthMin = 60.00
	enemyStrengthMax = 90.00
	enemyDefenceMin  = 40.00
	enemyDefenceMax  = 60.00
	enemySpeedMin    = 40.00
	enemySpeedMax    = 60.00
	enemyLuckMin     = 25.00
	enemyLuckMax     = 40.00
)

type Character interface {
	strike(warrior Warrior)
}

type Warrior struct {
	Human
	health    float32
	strength  float32
	defence   float32
	speed     float32
	luck      float32
	abilities []interface{}
}

type Human struct {
	name   string
	house  string
	isHero bool
}

func NewCharacter(name string, house string, isHero bool) Warrior {
	return Warrior{
		Human:     Human{name, house, isHero},
		health:    getDefaultHealth(isHero),
		strength:  getDefaultStrength(isHero),
		defence:   getDefaultDefence(isHero),
		speed:     getDefaultSpeed(isHero),
		luck:      getDefaultLuck(isHero),
		abilities: getDefaultAbilities(isHero),
	}
}

func (attacker *Warrior) strike(victim Warrior) {
	damage := attacker.strength - victim.defence
	if !attacker.isHero {
		victim.health -= attacker.strength - victim.defence
		fmt.Println("Strike from enemy. Damage done: ", damage)
	} else {
		victim.health -= damage
		fmt.Println("Strike from hero. Damage done: ", damage)
		chance := rand.Float64() * 100
		fmt.Println("Critical strike chance hit: " + strconv.FormatFloat(chance, 'f', -1, 64))
		if chance < 10 {
			fmt.Println("Hero used critical strike ...1 more strike")
			victim.health -= damage
			if chance < 1 {
				fmt.Println("..and another one")
				victim.health -= damage
			}
		}
	}
}

func getDefaultHealth(isHero bool) float32 {
	if isHero {
		return heroHealthMin + rand.Float32()*(heroHealthMax-heroHealthMin)
	}
	return enemyHealthMin + rand.Float32()*(enemyHealthMax-enemyHealthMin)
}

func getDefaultStrength(isHero bool) float32 {
	if isHero {
		return heroStrengthMin + rand.Float32()*(heroStrengthMax-heroStrengthMin)
	}
	return enemyStrengthMin + rand.Float32()*(enemyStrengthMax-enemyStrengthMin)
}

func getDefaultDefence(isHero bool) float32 {
	if isHero {
		return heroDefenceMin + rand.Float32()*(heroDefenceMax-heroDefenceMin)
	}
	return enemyDefenceMin + rand.Float32()*(enemyDefenceMax-enemyDefenceMin)
}

func getDefaultSpeed(isHero bool) float32 {
	if isHero {
		return heroSpeedMin + rand.Float32()*(heroSpeedMax-heroSpeedMin)
	}
	return enemySpeedMin + rand.Float32()*(enemySpeedMax-enemySpeedMin)
}

func getDefaultLuck(isHero bool) float32 {
	if isHero {
		return heroLuckMin + rand.Float32()*(heroLuckMax-heroLuckMin)
	}
	return enemyLuckMin + rand.Float32()*(enemyLuckMax-enemyLuckMin)
}

func getDefaultAbilities(isHero bool) []interface{} {
	if isHero {
		return []interface{}{CriticalStrike{}, Resilience{true}}
	}
	return []interface{}{NormalStrike{}}
}
