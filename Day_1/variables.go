package main

import "fmt"

/**
global scope of variable
 */

var thisVariableIsAccessibleWithinFunction = "Yes it can!"

//outsideShortAssignWillFailed := "Will this failed?, uncomment this and try"

func main() {
	/**
	to declare a variable we use `var`
	followed by variable's name	and type
	if an initializer is present, the type can be omitted,
	the variable will take the type of initializer
	 */
	var variableWithInitializer = "123"
	fmt.Print(variableWithInitializer)

	/**
	variable without initializer MUST declare a type
	an initial value of all empty var is 0
	if integer, 0 literally means 0 (zero)
	if string, 0 means empty string ""
	if bool, 0 means false
	 */
	var variableWithoutInitializer string
	fmt.Print(variableWithoutInitializer)


	/**
	go allows to declare multiple variable in one line
	with the type at the end of statement
	 */

	var isRich, hasCar, ownedACompany, useIphoneX bool
	fmt.Print(isRich, hasCar, ownedACompany, useIphoneX)

	/**
	go can assign multiple value respectively into multiple variables
	 */

	var name, gender, isMarried, age = "eric", "gentlemen", true, 99
	fmt.Print(name, gender, isMarried, age)

	fmt.Print(thisVariableIsAccessibleWithinFunction)

	/**
	variable declaration can be done using := short assign
	NOTE: it works inside function only!
	 */

	iAmShort := "yes I am short!"

	fmt.Print(iAmShort)
	fmt.Print()
}
