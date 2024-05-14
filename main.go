package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println("___main")
	// lets see some pointers
	myPointer()

	// launch gundams
	myStructs()

	//lets see how maps work
	myMaps()

	// lets see how slices work
	mySlices()

	// lets see some arrays at work
	myArrays()

	var divideResult, remainder, err = intDivide(100, 50)
	// switches are a good alternative to highly nested if/esleif/else
	switch {
	case err != nil:
		// pasing 0 would give compile error
		fmt.Printf("err.Error()\n")
	case remainder == 0:
		fmt.Printf("Divide Result: %v\n", divideResult)
	default:
		fmt.Printf("Divide Result: %v, remainder %v\n", divideResult, remainder)
	}
	// you can even get extra specific your switch statements
	switch remainder {
	case 0:
		fmt.Printf("The division was exact\n")
	case 1, 2:
		fmt.Printf("The division was close\n")
	default:
		fmt.Printf("The division was not close\n")
	}

	const name = "Cole"
	printMe(name)
	basicIfLogic()
	goConstVarAndBasicDataTypes()
	fmt.Println("============================")
}

// map is a set of kv pairs where you can do look up by key
func myMaps() {
	fmt.Println("___myMaps")
	// can make a new map with the make function again
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	// another way to make a map
	var myMap2 map[string]uint8 = map[string]uint8{"Adam": 23, "Sarah": 45, "Cole": 31}
	fmt.Println(myMap2)
	fmt.Println(myMap2["Adam"])
	// if you try to use an invalid key you get the default value of that type
	// so you have to be careful bc maps will always return something even if maps don't exsist
	fmt.Println(myMap2["Jason"])
	// but maps also return an optional second value so you can handle when when something
	// doesn't exsist or better stated, when nothing is returned by the map
	var age, ok = myMap2["Adam"]
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}
	// to delete a value of a map go has the built in delete function
	delete(myMap, "Adam")

	// thing of range just basically as length of the array in this case, don't over think
	// what you have seen people complain about on the internet
	// IMPORTANT: when itteration over a map no order is preseverd, so different order everytime
	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}

	fmt.Println("============================")
}

// slices are just wrappers around arrays
// general, powerful, and convenient interface to the sequences of data
func mySlices() {
	fmt.Println("___mySlices")
	// by omiting the length value we now have a slice
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capcity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capcity %v\n", len(intSlice), cap(intSlice))

	// can also append multiple values by using the spread opperator
	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	// us the make function to to make a new slice
	// if you do it this way (type,length,capacity)
	// if you leave capacity blank then capcity defaults to slice length
	// using make and setting both length and capcitity can aid in better
	// memory mangment/performance, this keeps you from having to do memory reallocation
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	fmt.Println("============================")
}

func myArrays() {
	fmt.Println("___myArrays")
	// initialize an array of size 3 type int32
	// arrays are 0 indexed
	// arrays in go = contiguous in memory
	// so in this case since an int32=4bytes of memory
	// this array is 4bytes x 3array items
	// thats 12 bytes total
	var intArr [3]int32
	// below is some other wasy to init an array
	var intArr2 [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(intArr2[0])
	intArr3 := [3]int32{1, 2, 3}
	fmt.Println(intArr3[0])
	// still fixed sized 3 bc of items in array
	intArr4 := [...]int32{1, 2, 3}
	fmt.Println(intArr4[0])

	// access index of array
	fmt.Println(intArr[0])
	// access slice of array
	fmt.Println(intArr[1:3])
	// lets print out the bytes of memory with &
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// basic for loop in go
	for i, v := range intArr2 {
		fmt.Printf("Index:, %v, Value%v \n", i, v)
	}

	//NOTE: go doesn't have a built in while loop, so you build it
	//basic while loop here
	var i int = 0
	for {
		if i > 10 {
			break
		}
		fmt.Println(i)
		i = i + 1
	}
	// same while loop here but condensed
	// initalization(i:=0),condition(i<10),post(i++)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("============================")
}

func intDivide(i int, y int) (int, int, error) {
	fmt.Println("___intDivde")
	var err error

	if y == 0 {
		err = errors.New("can't divide by zero")
		return 0, 0, err
	}

	fmt.Println("============================")
	return (i / y), (i % y), err
}

func printMe(name string) {
	fmt.Println("___printMe")
	fmt.Printf("Hey %v, I'm printing from the printMe func\n", name)
	fmt.Println("============================")
}

func goConstVarAndBasicDataTypes() {
	fmt.Println("___goConstVarAndBasicDataTypes")
	fmt.Println("Hello World!")

	// all default unassigned values for number types is 0
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum32 float32 = 12345678.9
	var floatNum64 float32 = 12345678.9
	fmt.Println(floatNum32)
	fmt.Println(floatNum64)

	// you cain't do math with mixed types you have to cast like below
	var result float32 = floatNum32 + float32(intNum)
	fmt.Println(result)

	// in division rounds down
	var intNum1 = 3
	var intNum2 = 2
	fmt.Println(intNum1 / intNum2)
	// have to use modulos sign to get non rounded result
	fmt.Println(intNum1 % intNum2)

	// strings and runes
	var myString string = "Hey Mom!"
	fmt.Println(myString)
	// get the len or the number of bytes in a string its not the number of charaters
	fmt.Println(len(myString))
	// if you use fancy charactes it won't be a 1-1 comparison
	fmt.Println(len("γ"))
	// can use utf8 package and util below to get the count of rune
	fmt.Println(utf8.RuneCountInString("γ"))
	// if you use single quotes like this you can get a rune
	var myRune rune = 'a'
	fmt.Println(myRune)

	// the default for all unassinged bool's is fals
	var myBoolean bool = false
	fmt.Println(myBoolean)

	// can omit types if we want to let the compiler infer the type like below
	var myInferredType = "myInferredType"
	fmt.Println(myInferredType)
	// you can also use shorthand and inferred types
	myInferredShortHandVar := "myInferredShortHandVar"
	fmt.Println(myInferredShortHandVar)
	// you can also do multi var initialization
	var1, var2 := 11, 22
	// and print multiple vars like so
	fmt.Println(var1, var2)
	const myConst string = "constvalue"
	fmt.Println(myConst)

	stringsRunesBytes()

	fmt.Println("============================")
}

func basicIfLogic() {
	fmt.Println("___basicIfLogic")
	if 1 != 2 && 2 != 3 {
		fmt.Println("check passed")
	} else {
		fmt.Println("check did not pass")
	}

	if 1 != 2 || 2 != 3 {
		fmt.Println("other check passed")
	} else {
		fmt.Println("check did not pass")
	}

	fmt.Println("============================")
}

func stringsRunesBytes() {
	fmt.Println("___stringsRunesBytes")
	// strings are reperesented by bytes, ascii
	// so you gotta remember that when dealing with strings/itterating over them
	// an easier way to deal with this to cast to a rune([]rune)
	// runes are just an alias for int32
	// you can also declare a rune type like this with single quotes
	// var myRune = 'a'
	var myRunes = []rune("r𝔼sum𝔼") //var myRunes = "r𝔼sum𝔼"
	var indexed = myRunes[0]
	fmt.Printf("%v,%T\n", indexed, myRunes)
	for i, v := range myRunes {
		fmt.Println(i, v)
	}

	//strings are immutable in go so to avoid memoryk issues and create a new string
	// each time you access things you can use string builder
	var stringSlice = []string{"C", "o", "l", "e"}
	var stringBuilder strings.Builder
	for i := range stringSlice {
		stringBuilder.WriteString(stringSlice[i])
	}
	var catStr = stringBuilder.String()
	fmt.Printf("\n%v\n", catStr)

	fmt.Println("Slices contain pointers to the same data, so if you copy a slice and update it you will update the other slice as well")

	fmt.Println("============================")
}

type gundam struct {
	pilot        string
	unit         string
	canTransform bool
}

func (g gundam) transform() (string, error) {
	if !g.canTransform {
		return "", errors.New("TMS support only")
	}
	fmt.Println("")
	return g.unit, nil
}

type mobileArmor struct {
	pilot        string
	unit         string
	canTransform bool
}

func (ma mobileArmor) transform() (string, error) {
	if !ma.canTransform {
		return "", errors.New("TMS support only")
	}
	fmt.Println("Transforming!")
	return ma.unit, nil
}

type mobileSuit interface {
	transform() (string, error)
}

func myStructs() {
	fmt.Println("___myStructs")
	// how to init structs
	var myGundam gundam = gundam{pilot: "Cole", unit: "MSZ006Zeta"}
	var amurosGundam gundam = gundam{pilot: "Amuro", unit: "HivNu"}
	// how to init annoymous structs
	var theEnemyUnit1 = struct {
		pilot string
		unit  string
	}{"Zeon Ace", "Zaku II"}
	var theEnemyUnit2 = struct {
		pilot string
		unit  string
	}{"Char", "Zaku II"}
	var theEnemyUnitMobileArmor = mobileArmor{"Zeon Ace", "Mobile Armor", true}

	fmt.Printf("Gundam Unit 1, %v: %v, launching!\n", myGundam.pilot, myGundam.unit)
	fmt.Printf("Gundam Unit 2, %v: %v, launching!\n", amurosGundam.pilot, amurosGundam.unit)
	fmt.Printf("Enemy Unit Spotted! %v: %v\n", theEnemyUnit1.pilot, theEnemyUnit1.unit)
	fmt.Printf("Enemy Unit Spotted! %v: %v\n", theEnemyUnit2.pilot, theEnemyUnit2.unit)
	fmt.Printf("Enemy Unit Spotted! %v: %v\n", theEnemyUnitMobileArmor.pilot, theEnemyUnitMobileArmor.unit)
	fmt.Printf("This is %v. Enemy spotted, moving to intercept!\n", myGundam.pilot)

	myUnit, err := myGundam.transform()
	if err != nil {
		fmt.Println("Error:", err)
		// Handle the error, perhaps return or exit the function
		return
	}
	fmt.Println("Transforming! ", myUnit)

	enemyMA, err := theEnemyUnitMobileArmor.transform()
	if err != nil {
		fmt.Println("Error:", err)
		// Handle the error, perhaps return or exit the function
		return
	}
	fmt.Println("Transforming! ", enemyMA)

	fmt.Println("============================")
}

func myPointer
	// pointers store memory locations to variables/values/etc..

	// nil pointer(unless initalized)that will hold an int32
	var p *int32 = new(int32)
	var i int32 = 21

	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value i points to is: %v\n", i)

	// ok lets set the value of what p i point to to 10
	*p = 10 // see the * is used for initalization and referencing
	fmt.Printf("The value p points to is: %v\n", *p)
	// now lets make p hold the same value of i with the & syntax
	p = &i
	fmt.Printf("The value p points to is: %v\n", *p)
	// now lets change the value of the pointer p and see that i will change
	*p = 1
	fmt.Printf("The value i points to is: %v\n", i)

	// if you use pointers with functions you can save a ton on memory
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe mem location of thing1 array is: %p", &thing1)
	var result [5]float64 = myPointerFuncSquareHeavyMemory(thing1)
	fmt.Printf("\nThe result is: %v", result)
	fmt.Printf("\nThe value of thing1 is: %v", thing1)
	fmt.Printf("\nMaking a compy for our function so potentially more memory")

	var thing3 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe mem location of thing3 array is: %p", &thing3)
	var result [5]float64 = myPointerFuncSquareHeavyMemory(thing3)
	fmt.Printf("\nThe result is: %v", result)
	fmt.Printf("\nThe value of thing3 is: %v", thing3)
	fmt.Printf("\nPointers are great when you have to pass in a lot bc you don't have to create copys so you save on memory")

}

func myPointerFuncSquareHeavyMemory(thing2 [5]float64) [5]float64 {
	// square each item in the array
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
}


func myPointerFuncSquareLessMemory(thing2 *[5]float64) [5]float64 {
	// square each item in the array
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
}

