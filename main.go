package main

import (
	"fmt"
	"github.com/Poportss/annalyn-s-infiltration-leet-code/annalyn"
)

func main() {
	knightIsAwake := false
	archerIsAwake := true
	prisonerIsAwake := false
	petDogIsPresent := true

	fmt.Println("Can Fast Attack:", annalyn.CanFastAttack(knightIsAwake))
	fmt.Println("Can Spy:", annalyn.CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake))
	fmt.Println("Can Signal Prisoner:", annalyn.CanSignalPrisoner(archerIsAwake, prisonerIsAwake))
	fmt.Println("Can Free Prisoner:", annalyn.CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent))
}
