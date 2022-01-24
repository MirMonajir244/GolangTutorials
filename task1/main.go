package main

import "fmt"

var id int
var isKnightSleep bool   //decides fast attack
var isPrisonerSleep bool //decides signal prisoner
var isArcherSleep bool   //decides signal prisoner
var hasDog bool

//var isAllSleep bool
func CanFastAttack() bool {
	return isKnightSleep
}

//isAllSleep = isKnightSleep && isPrisonerSleep && isArcherSleep
func CanSpy() bool {
	return !(isKnightSleep && isPrisonerSleep && isArcherSleep)
}

func CanSignalPrisoner() bool {
	return (!isPrisonerSleep) && isArcherSleep
}

func CanFreePrisoner() bool {
	return ((hasDog && isArcherSleep) || (!isPrisonerSleep && isArcherSleep && isKnightSleep))

}

func main() {
	isArcherSleep, isPrisonerSleep, isKnightSleep, hasDog = true, true, true, true

	fmt.Println(" the result is: ", CanFastAttack(), CanFreePrisoner(), CanSignalPrisoner())
}
