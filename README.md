# exercism-exercises

# Exercism Exercises in Go

# Exercise 1: Basics -- Gopher's Gorgeous Lasagna

## Instructions and Info

## Introduction

[Go](https://golang.org/) is a statically typed, compiled programming language. This exercise introduces three major language features: Packages, Functions, and Variables.

### Packages

Go applications are organized in packages. A package is a collection of source files located in the same directory. All source files in a directory must share the same package name. When a package is imported, only entities (functions, types, variables, constants) whose names start with a capital letter can be used / accessed. The recommended style of naming in Go is that identifiers will be named using `camelCase`, except for those meant to be accessible across packages which should be `PascalCase`.

```Go
package lasagna
```

### Variables

Go is statically-typed, which means all variables [must have a defined type](https://en.wikipedia.org/wiki/Type_system) at compile-time.

Variables can be defined by explicitly specifying a type:

```Go
var explicit int // Explicitly typed
```

You can also use an initializer, and the compiler will assign the variable type to match the type of the initializer.

```Go
implicit := 10   // Implicitly typed as an int
```

Once declared, variables can be assigned values using the `=` operator. Once declared, a variable's type can never change.

```Go
count := 1 // Assign initial value
count = 2  // Update to new value

count = false // This throws a compiler error due to assigning a non `int` type
```

### Constants

Constants hold a piece of data just like variables, but their value cannot change during the execution of the program.

Constants are defined using the `const` keyword and can be numbers, characters, strings or booleans:

```Go
const Age = 21 // Defines a numeric constant 'Age' with the value of 21
```

### Functions

Go functions accept zero or more parameters. Parameters must be explicitly typed, there is no type inference.

Values are returned from functions using the `return` keyword.

A function is invoked by specifying the function name and passing arguments for each of the function's parameters.

Note that Go supports two types of comments. Single line comments are preceded by `//` and multiline comments are inserted between `/*` and `*/`.

```Go
package greeting

// Hello is a public function.
func Hello (name string) string {
    return hi(name)
}

// hi is a private function.
func hi (name string) string {
    return "hi " + name
}
```

## Instructions

In this exercise you're going to write some code to help you cook a brilliant lasagna from your favorite cooking book.

You have four tasks, all related to the time spent cooking the lasagna.

Define the `OvenTime` constant with how many minutes the lasagna should be in the oven. According to the cooking book, the expected oven time in minutes is 40:

```Go
OvenTime
// => 40
```

## How to debug

When a test fails, a message is displayed describing what went wrong and for which input. You can also use the fact that [console output](https://pkg.go.dev/fmt#Println) will be shown too. You can write to the console using:

```Go
import "fmt"
fmt.Println("Debug message")
```

Note: When debugging with the in-browser editor, using e.g. `fmt.Print` will not add a newline after the output which can provoke a bug in `go test --json` and result in messed up test output.

## My Solution
```Go
package lasagna
import "fmt"

// TODO: define the 'OvenTime' constant
const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
    fmt.Println(OvenTime)
	fmt.Println(OvenTime - actualMinutesInOven)
    return (OvenTime - actualMinutesInOven)
}
// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	fmt.Println(2*numberOfLayers)
    return (2*numberOfLayers)
}
// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	var x int
    x = PreparationTime(numberOfLayers)
    return (x + actualMinutesInOven)
}

```

## Learnings and Points to Note
1. const variable declarations use ```=``` and not ```:=```
2. fmt.Println and other executable statements at the package level only work inside the function and not outside

# Exercise 2: Comments -- Weather Forecast 

## Instructions and Info

## Introduction

In the previous exercise, we saw that there are two ways to write comments in Go: single-line comments that are preceded by `//`, and multiline comment blocks that are wrapped with `/*` and `*/`.

### Documentation comments

In Go, comments play an important role in documenting code. They are used by the `godoc` command, which extracts these comments to create documentation about Go packages. A documentation comment should be a complete sentence that starts with the name of the thing being described and ends with a period.

Comments should precede packages as well as exported identifiers, for example exported functions, methods, package variables, constants, and structs, which you will learn more about in the next exercises.

A package-level variable can look like this:

```go
// TemperatureCelsius represents a certain temperature in degrees Celsius.
var TemperatureCelsius float64
```

### Package comments

Package comments should be written directly before a package clause (`package x`) and begin with `Package x ...` like this:

```go
// Package kelvin provides tools to convert
// temperatures to and from Kelvin.
package kelvin
```

### Function comments

A function comment should be written directly before the function declaration. It should be a full sentence that starts with the function name. For example, an exported comment for the function `Calculate` should take the form `Calculate ...`. It should also explain what arguments the function takes, what it does with them, and what its return values mean, ending in a period):

```go
// CelsiusFreezingTemp returns an integer value equal to the temperature at which water freezes in degrees Celsius.
func CelsiusFreezingTemp() int {
	return 0
}
```

## Instructions

Goblinocus is a country that takes its weather forecast very seriously. Since you are a renowned, responsible and proficient developer, they asked you to write a program that can forecast the current weather condition of various cities in Goblinocus. You were busy at the time and asked one of your friends to do the job instead. After a while, the president of Goblinocus contacted you and said they do not understand your friend's code. When you check the code, you discover that your friend did not act as a responsible programmer and there are no comments in the code. You feel obligated to clarify the program so goblins can understand them as well.

Since goblins are not as smart as you are, they forgot what the package should do for them. Please write a comment for `package weather` that describes its contents. The package comment should introduce the package and provide information relevant to the package as a whole.

The president of Goblinocus is a bit paranoid and fears uncommented variables are used to destroy their country. Please clarify the usage of the package variables `CurrentCondition` and `CurrentLocation` and put the president's mind at ease. This should tell any user of the package what information the variables store, and what they can do with it.

Goblinocus forecast operators want to know what the `Forecast()` function does (but do not tell them how it works, since unfortunately, they will get more confused). Please write a comment for this function that describes what the function does, but not how it does it.
## How to debug

When a test fails, a message is displayed describing what went wrong and for which input. You can also use the fact that [console output](https://pkg.go.dev/fmt#Println) will be shown too. You can write to the console using:

```Go
import "fmt"
fmt.Println("Debug message")
```

Note: When debugging with the in-browser editor, using e.g. `fmt.Print` will not add a newline after the output which can provoke a bug in `go test --json` and result in messed up test output.

## My Solution

```Go
// Package weather provides the tools to forecast current weather conditons in various cites in Goblinocus.
package weather
// CurrentCondition represents the current weather condition in a city as a string.
var CurrentCondition string
// CurrentLocation represent the Current Location where the weather condition is being calculated.
var CurrentLocation string
// Forecast calculates and returns a string value that represents the current location along with the current weather conditon which is also a string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
```

## Learnings and Points to note
Comments should begin with the identifier name (e.g. function, variable etc.. name)
Comments should end in a period ```.```
```godoc``` command can be used to generate documentation by extracting comments from the code. This requires that the comments precede packages as well as exported identifiers, for example exported functions, methods, package variables, constants, and structs.


# Exercise 3: Numbers and Arithmetic Operators -- Cars Assemble

## Instructions and Info
## Introduction

### Numbers

Go contains basic numeric types that can represent sets of either integer or floating-point values. There are different types depending on the size of value you require and the architecture of the computer where the application is running (e.g. 32-bit and 64-bit).

For the sake of this exercise you will only be dealing with:

- `int`: e.g. `0`, `255`, `2147483647`. A signed integer that is at least 32 bits in size (value range of: -2147483648 through 2147483647). But this will depend on the systems architecture. Most modern computers are 64 bit, therefore `int` will be 64 bits in size (value rate of: -9223372036854775808 through 9223372036854775807).
    
- `float64`: e.g. `0.0`, `3.14`. Contains the set of all 64-bit floating-point numbers.
    
- `uint`: e.g. `0`, `255`. An unsigned integer that is the same size as `int` (value range of: 0 through 4294967295 for 32 bits and 0 through 18446744073709551615 for 64 bits)
    

Numbers can be converted to other numeric types through Type Conversion.

### Arithmetic Operators

Go supports many standard arithmetic operators:

| Operator | Example        |
| -------- | -------------- |
| `+`      | `4 + 6 == 10`  |
| `-`      | `15 - 10 == 5` |
| `*`      | `2 * 3 == 6`   |
| `/`      | `13 / 3 == 4`  |
| `%`      | `13 % 3 == 1`  |

For integer division, the remainder is dropped (e.g. `5 / 2 == 2`).

Go has shorthand assignment for the operators above (e.g. `a += 5` is short for `a = a + 5`). Go also supports the increment and decrement statements `++` and `--` (e.g. `a++`).

### Converting between types

Converting between types is done via a function with the name of the type to convert to. For example:

```Go
var x int = 42 // x has type int
f := float64(x) // f has type float64 (ie. 42.0)
var y float64 = 11.9 // y has type float64
i := int(y) // i has type int (ie. 11)
```

### Arithmetic operations on different types

In many languages you can perform arithmetic operations on different types of variables, but in Go this gives an error. For example:

```Go
var x int = 42

// this line produces an error
value := float32(2.0) * x // invalid operation: mismatched types float32 and int

// you must convert int type to float32 before performing arithmetic operation
value := float32(2.0) * float32(x)
```

## Instructions

In this exercise you'll be writing code to analyze the production in a car factory.

### 1. Calculate the number of working cars produced per hour

The cars are produced on an assembly line. The assembly line has a certain speed, that can be changed. The faster the assembly line speed is, the more cars are produced. However, changing the speed of the assembly line also changes the number of cars that are produced successfully, that is cars without any errors in their production.

Implement a function that takes in the number of cars produced per hour and the success rate and calculates the number of successful cars made per hour. The success rate is given as a percentage, from `0` to `100`:

```Go
CalculateWorkingCarsPerHour(1547, 90)
// => 1392.3
```

**Note:** the return value should be a `float64`.

### 2. Calculate the number of working cars produced per minute

Implement a function that takes in the number of cars produced per hour and the success rate and calculates how many cars are successfully produced each minute:

```Go
CalculateWorkingCarsPerMinute(1105, 90)
// => 16
```

**Note:** the return value should be an `int`.

### 3. Calculate the cost of production

Each car normally costs $10,000 to produce individually, regardless of whether it is successful or not. But with a bit of planning, 10 cars can be produced together for $95,000.

For example, 37 cars can be produced in the following way: 37 = 3 x groups of ten + 7 individual cars

So the cost for 37 cars is: 3*95,000+7*10,000=355,000

Implement the function `CalculateCost` that calculates the cost of producing a number of cars, regardless of whether they are successful:

```Go
CalculateCost(37)
// => 355000

CalculateCost(21)
// => 200000
```

**Note:** the return value should be an `uint`.

## My Solution

```Go
package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	pr := float64(productionRate) / 100 // I got an error here because I was doing float64 type 
    // conversion for the whole operation i.e. float64(productionRate / 100) that led to an error
    // for multiple test cases. This is because integer division messed things up.
    // float64(productionRate / 100) has both operands being int so the Go compiler does integer 
    // division first.
    // So if productionRate was 221 then 221 by 100 would be 2, which would then be converted to 
    // float64 type becoming 2.00 . However, if I did float64(productionRate) / 100, I would get 2.21
    // which would be correct. Because untyped constants adopt types based on the context. So 
    // multiplying by 0.01 or dividing by 100 works because the operand they are interacting, in 
    // this case, productionRate is of type floating point.
    sr := float64(successRate) * 0.01
    wr := pr * sr
    return wr * 100 // Apparently you can multiply a float variable like wr with 100 but if I were to 
    // declare a variable called percent of type int with a value 100 and and multiply it I get an
    // error. This is due to the same reason as above wherein untyped constants get typed based on c
    // context.
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
/*
Mistakes I made earlier and why they were wrong:

1) Integer division at the wrong time
   - I had: carsPerMin := productionRate / 60
   - Both operands were int, so Go did integer division first (truncation). Example: 221/60 == 3.
     That lost the fractional cars/minute and skewed every subsequent calculation.

2) Treating a percentage as a whole number instead of a fraction
   - I multiplied by successRate directly (e.g., 90) instead of using 0.90.
   - Converting with int(successRate) or float64(successRate) alone is not enough; I must divide by 100
     (or multiply by 0.01) to convert a percent to its fractional form.

3) Type-mixing confusion (constants vs. typed values)
   - In Go, untyped numeric constants (like 60 or 100) adopt the type of the other operand.
     If both sides are int, integer math happens; if one side is float64, the operation is floating-point.
     This is why float64(productionRate)/60 works (60 becomes 60.0), but productionRate/60 converted
     afterwards would already be truncated.

Why this version works:

1) float64 before dividing by 60.0
   - carsPerMin := float64(productionRate) / 60.0
   - Forcing float math preserves the fractional part (no truncation).

2) Convert percent to fraction
   - successRate is given in [0,100]; multiplying by 0.01 converts it to [0,1].
   - minSucc := carsPerMin * (successRate * 0.01)

3) Truncate only once at the end
   - The spec/tests want an int count per minute. Returning int(minSucc) truncates toward zero at the final step.

Sanity checks:
- (221, 100%)  → (221/60) ≈ 3.683 → 3
- (426, 80%)   → (426*0.8)/60 = 5.68 → 5
- (6824, 20.5%)→ (6824*0.205)/60 ≈ 23.31 → 23
*/
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	
    carsPerMin := float64(productionRate)/60.0
    minSucc := (float64(carsPerMin) * (float64(successRate) * 0.01))
    return int(minSucc)
    
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    carsGroup := carsCount/10
    carsSingle := carsCount%10
    return uint(95000*carsGroup) + uint(10000*carsSingle)
}
```

## Learnings and Points to Note

```Go
1. Convert to float **before** dividing to avoid integer truncation: use float64(productionRate) / 60.0, not productionRate / 60.
    
2. Percent ≠ fraction: multiply by 0.01 (or divide by 100.0) so 90 becomes 0.90; never use int(successRate) for this.
    
3. Untyped constants adopt context: in float64(productionRate) / 60, the 60 becomes 60.0; if you store sixty := 60 (typed int), you must convert: float64(productionRate) / float64(sixty).
    
4. Do all math in float64, then convert **once at the end** with int(...) to match tests’ truncation behavior.
    
5. If rounding is ever required, use math.Round explicitly; don’t rely on casts for rounding.
    
6. Prefer clear literals like 60.0 and 100.0 (or * 0.01) to make floating-point intent obvious.
    
7. Reuse existing logic to stay DRY: int(CalculateWorkingCarsPerHour(rate, success) / 60.0) keeps both functions consistent.
   
8. Keep GoDoc comments above the function, start with the identifier name, and end with a period; this helps godoc and future you.
```


# Exercise 4: Booleans -- Annalyn's Infiltration

## Instructions and Info

## Introduction

Booleans in Go are represented by the predeclared boolean type `bool`, which values can be either `true` or `false`. It's a defined type.

```go
var closed bool    // boolean variable 'closed' implicitly initialized with 'false'
speeding := true   // boolean variable 'speeding' initialized with 'true'
hasError := false  // boolean variable 'hasError' initialized with 'false' 
```

Go supports three logical operators that can evaluate expressions down to Boolean values, returning either `true` or `false`.

|Operator|What it means|
|---|---|
|`&&` (and)|It is true if both statements are true.|
|`\|` (or)|It is true if at least one statement is true.|
|`!` (not)|It is true only if the statement is false.

## Instructions

In this exercise, you'll be implementing the quest logic for a new RPG game a friend is developing. The game's main character is Annalyn, a brave girl with a fierce and loyal pet dog. Unfortunately, disaster strikes, as her best friend was kidnapped while searching for berries in the forest. Annalyn will try to find and free her best friend, optionally taking her dog with her on this quest.

After some time spent following her best friend's trail, she finds the camp in which her best friend is imprisoned. It turns out there are two kidnappers: a mighty knight and a cunning archer.

Having found the kidnappers, Annalyn considers which of the following actions she can engage in:

- _Fast attack_: a fast attack can be made if the knight is sleeping, as it takes time for him to get his armor on, so he will be vulnerable.
- _Spy_: the group can be spied upon if at least one of them is awake. Otherwise, spying is a waste of time.
- _Signal prisoner_: the prisoner can be signaled using bird sounds if the prisoner is awake and the archer is sleeping, as archers are trained in bird signaling so they could intercept the message.
- _Free prisoner_: Annalyn can try sneaking into the camp to free the prisoner. This is a risky thing to do and can only succeed in one of two ways:
    - If Annalyn has her pet dog with her she can rescue the prisoner if the archer is asleep. The knight is scared of the dog and the archer will not have time to get ready before Annalyn and the prisoner can escape.
    - If Annalyn does not have her dog then she and the prisoner must be very sneaky! Annalyn can free the prisoner if the prisoner is awake and the knight and archer are both sleeping, but if the prisoner is sleeping they can't be rescued: the prisoner would be startled by Annalyn's sudden appearance and wake up the knight and archer.

You have four tasks: to implement the logic for determining if the above actions are available based on the state of the three characters found in the forest and whether Annalyn's pet dog is present or not.

### 1. Check if a fast attack can be made

Define the `CanFastAttack()` function that takes a boolean value that indicates if the knight is awake. This function returns `true` if a fast attack can be made based on the state of the knight. Otherwise, returns `false`:

```go
var knightIsAwake = true
fmt.Println(CanFastAttack(knightIsAwake))
// Output: false
```

### 2. Check if the group can be spied upon

Define the `CanSpy()` function that takes three boolean values, indicating if the knight, archer and the prisoner, respectively, are awake. The function returns `true` if the group can be spied upon, based on the state of the three characters. Otherwise, returns `false`:

```go
var knightIsAwake = false
var archerIsAwake = true
var prisonerIsAwake = false
fmt.Println(CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake))
// Output: true
```

### 3. Check if the prisoner can be signaled

Define the `CanSignalPrisoner()` function that takes two boolean values, indicating if the archer and the prisoner, respectively, are awake. The function returns `true` if the prisoner can be signaled, based on the state of the two characters. Otherwise, returns `false`:

```go
var archerIsAwake = false
var prisonerIsAwake = true
fmt.Println(CanSignalPrisoner(archerIsAwake, prisonerIsAwake))
// Output: true
```

### 4. Check if the prisoner can be freed

Define the `CanFreePrisoner()` function that takes four boolean values. The first three parameters indicate if the knight, archer and the prisoner, respectively, are awake. The last parameter indicates if Annalyn's pet dog is present. The function returns `true` if the prisoner can be freed based on the state of the three characters and Annalyn's pet dog presence. Otherwise, it returns `false`:

```go
var knightIsAwake = false
var archerIsAwake = true
var prisonerIsAwake = false
var petDogIsPresent = false
fmt.Println(CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent))
// Output: false
```

## My Solution

```Go
package annalyn

// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
	if knightIsAwake == true{
        return false
    } else {
        return true
    }
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	if knightIsAwake || archerIsAwake || prisonerIsAwake {
        return true
    } else {
        return false
    }
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
    if prisonerIsAwake && !archerIsAwake {
        return true
    } else {
        return false
    }
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	if petDogIsPresent && !archerIsAwake {
            return true
        } else if !petDogIsPresent && !archerIsAwake && prisonerIsAwake && !knightIsAwake {
            return true
        } else {
            return false
        } 
}

```

## Learnings and Points to Note

1. I initially had defined the last function in the following manner:
```Go
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	if petDogIsPresent && !archerIsAwake {
            return true
        } else {
            return false
        }
    if !petDogIsPresent && !archerIsAwake  && prisonerIsAwake && !knightIsAwake {
            return true
        } else {
            return false
        } 
}
```

However, one test case failed: Knight and archer are sleeping. Prisoner is awake. Dog is not present.
```go

func TestCanFreePrisoner(t *testing.T) {
	tt := charactersState{
		desc:            "Knight and archer are sleeping. Prisoner is awake. Dog is not present.",
		knightIsAwake:   false,
		archerIsAwake:   false,
		prisonerIsAwake: true,
		dogIsPresent:    false,
		expected:        true,
	}

	if got := CanFreePrisoner(tt.knightIsAwake, tt.archerIsAwake, tt.prisonerIsAwake, tt.dogIsPresent); got != tt.expected {
		t.Errorf("CanFreePrisoner(%v,%v,%v,%v) = %v; want %v", tt.knightIsAwake, tt.archerIsAwake, tt.prisonerIsAwake, tt.dogIsPresent, got, tt.expected)
	}

}
```

```
#### Test Failure

=== RUN   TestCanFreePrisoner/Knight_and_archer_are_sleeping._Prisoner_is_awake._Dog_is_not_present.

    annalyns_infiltration_test.go:290: CanFreePrisoner(false,false,true,false) = false; want true

--- FAIL: TestCanFreePrisoner/Knight_and_archer_are_sleeping._Prisoner_is_awake._Dog_is_not_present. (0.00s)

```

This was because the first else condition would execute and return a false and the second if statement would never even have a chance to execute. Changing the code to an if else-if else structure led to the else-if getting executed and therefore the test case passing.

2. In the function where the prisoner is awake and the archer is sleeping and signalling can happen, I did this:
```Go
	// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
    if prisonerIsAwake && !archerIsAwake {
        return true
    } else {
        return false
    }
}
```

This was surprising because my initial instinct was to have a multiple if statement structure or make the condition more complex but I tried this because it intuitively made sense. And it worked!! This passes all permutations and combinations.

# Exercise 5: Strings and Strings Package -- Welcome to Tech Palace

## Instructions and Info

## Introduction

A `string` in Go is an immutable sequence of bytes, which don't necessarily have to represent characters.

A string literal is defined between double quotes:

```go
const name = "Jane"
```

Strings can be concatenated via the `+` operator:

```go
"Jane" + " " + "Austen"
// => "Jane Austen"
```

Some special characters need to be escaped with a leading backslash, such as `\t` for a tab and `\n` for a new line in strings.

```go
"How is the weather today?\nIt's sunny"  
// =>
// How is the weather today?
// It's sunny
```

The `strings` package contains many useful functions to work on strings. For more information about string functions, check out the [strings package documentation](https://pkg.go.dev/strings). Here are some examples:

```go
import "strings"

// strings.ToLower returns the string given as argument with all its characters lowercased
strings.ToLower("MaKEmeLoweRCase")
// => "makemelowercase"

// strings.Repeat returns a string with a substring given as argument repeated many times
strings.Repeat("Go", 3)
// => "GoGoGo"
```

## Instructions

There is an appliance store called "Tech Palace" nearby. The owner of the store recently installed a big display to use for marketing messages and to show a special greeting when customers scan their loyalty cards at the entrance. The display consists of lots of small LED lights and can show multiple lines of text.

The store owner needs your help with the code that is used to generate the text for the new display.

For most customers who scan their loyalty cards, the store owner wants to see `Welcome to the Tech Palace,`  followed by the name of the customer in capital letters on the display.

Implement the function `WelcomeMessage` that accepts the name of the customer as a `string` argument and returns the desired message as a `string`.

```go
WelcomeMessage("Judy")
// => Welcome to the Tech Palace, JUDY
```

For loyal customers that buy a lot at the store, the owner wants the welcome display to be more fancy by adding a line of stars before and after the welcome message. They are not sure yet how many stars should be in the lines so they want that to be configurable.

Write a function `AddBorder` that accepts a welcome message (a `string`) and the number of stars per line (type `int`) as arguments. It should return a `string` that consists of 3 lines, a line with the desired number of stars, then the welcome message as it was passed in, then another line of stars.

```go
AddBorder("Welcome!", 10)
```

Should return the following:

```plain
**********
Welcome!
**********
```

Before installing this new display, the store had a similar display that could only show non-configurable, static lines. The owner would like to reuse some of the old marketing messages on the new display. However, the data already includes a star border and some unfortunate whitespaces. Your task is to clean up the messages so they can be re-used.

Implement a function `CleanUpMessage` that accepts the old marketing message as a string. The function should first remove all stars from the text and afterwards remove the leading and trailing whitespaces from the remaining text. The function should then return the cleaned up message.

```go
message := `
**************************
*    BUY NOW, SAVE 10%   *
**************************
`

CleanUpMessage(message)
// => BUY NOW, SAVE 10%
```

## My Solution

```Go
package techpalace
import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    s := strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
    return s
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	s := strings.ReplaceAll(oldMsg, "*", " ")
    t := strings.TrimSpace(s)
    return t
}

```

## Learnings and Points to Note

1. string package is super useful for string manipulation: https://pkg.go.dev/strings
2. For the third function, the way to clean it up was something I had to think about. Since we don't know the input string, the solution was to replace all ```*``` with a whitespace using ```strings.ReplaceAll``` and then trim all the whitespaces at the front and back using ```strings.TrimSpace```

# Exercise 6: Conditionals If and Comparison -- Vehicle Purchase

## Instructions and Info

## Introduction

### Comparison

In Go numbers can be compared using the following relational and equality operators.

|Comparison|Operator|
|---|---|
|equal|`==`|
|not equal|`!=`|
|less|`<`|
|less or equal|`<=`|
|greater|`>`|
|greater or equal|`>=`|

The result of the comparison is always a boolean value, so either `true` or `false`.

```go
a := 3

a != 4 // true
a > 5  // false
```

The comparison operators above can also be used to compare strings. In that case a lexicographical (dictionary) order is applied. For example:

```Go
	"apple" < "banana"  // true
	"apple" > "banana"  // false
```

### If Statements

Conditionals in Go are similar to conditionals in other languages. The underlying type of any conditional operation is the `bool` type, which can have the value of `true` or `false`. Conditionals are often used as flow control mechanisms to check for various conditions.

For checking a particular case an `if` statement can be used, which executes its code if the underlying condition is `true` like this:

```go
var value string

if value == "val" {
    return "was val"
}
```

In scenarios involving more than one case many `if` statements can be chained together using the `else if` and `else` statements.

```go
var number int
result := "This number is "

if number > 0 {
    result += "positive"
} else if number < 0 {
    result += "negative"
} else {
    result += "zero"
}
```

If statements can also include a short initialization statement that can be used to initialize one or more variables for the if statement. For example:

```go
num := 7
if v := 2 * num; v > 10 {
    fmt.Println(v)
} else {
    fmt.Println(num)
}
// Output: 14
```

> Note: any variables created in the initialization statement go out of scope after the end of the if statement.

## Instructions

In this exercise, you are going to write some code to help you prepare to buy a vehicle.

You have three tasks, one to determine if you need a license, one to help you choose between two vehicles and one to estimate the acceptable price for a used vehicle.

Some vehicle kinds require a driver's license to operate them. Assume only the kinds `"car"` and `"truck"` require a license, everything else can be operated without a license.

Implement the `NeedsLicense(kind)` function that takes the kind of vehicle and returns a boolean indicating whether you need a license for that kind of vehicle.

```go
needLicense := NeedsLicense("car")
// => true

needLicense = NeedsLicense("bike")
// => false

needLicense = NeedsLicense("truck")
// => true
```

You evaluate your options of available vehicles. You manage to narrow it down to two options but you need help making the final decision. For that, implement the function `ChooseVehicle(option1, option2)` that takes two vehicles as arguments and returns a decision that includes the option that comes first in lexicographical order.

```go
vehicle := ChooseVehicle("Wuling Hongguang", "Toyota Corolla")
// => "Toyota Corolla is clearly the better choice."

ChooseVehicle("Volkswagen Beetle", "Volkswagen Golf")
// => "Volkswagen Beetle is clearly the better choice."
```

Now that you made a decision, you want to make sure you get a fair price at the dealership. Since you are interested in buying a used vehicle, the price depends on how old the vehicle is. For a rough estimate, assume if the vehicle is less than 3 years old, it costs 80% of the original price it had when it was brand new. If it is at least 10 years old, it costs 50%. If the vehicle is at least 3 years old but less than 10 years, it costs 70% of the original price.

Implement the `CalculateResellPrice(originalPrice, age)` function that applies this logic using `if`, `else if` and `else` (there are other ways if you want to practice). It takes the original price and the age of the vehicle as arguments and returns the estimated price in the dealership.

```go
CalculateResellPrice(1000, 1)
// => 800

CalculateResellPrice(1000, 5)
// => 700

CalculateResellPrice(1000, 15)
// => 500
```

**Note** the return value is a `float64`.

## My Solution

```Go
package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
    if kind == "car" || kind == "truck"{
        return true
    } else {
        return false
    }
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 < option2{
        return option1 + " is clearly the better choice."
    } else if option2 < option1 {
        return option2 + " is clearly the better choice."
    } else {
        return "They are both great options."
    }
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age >= 3 && age < 10{
        return 0.7 * originalPrice
    } else if age >= 10 {
        return 0.5 * originalPrice
    } else {
        return 0.8 * originalPrice
    }
}
```

## Learnings and Points to Note

In task 2, when I did this:

```Go
// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order. 
func ChooseVehicle(option1, option2 string) string { 
	if option1 < option2 { 
		return option1 + " is clearly the better choice." 
	} else if option2 < option1 { 
		return option2 + " is clearly the better choice." 
	} 
}
```

I get the following error:
```Error
We received the following error when we ran your code: FAIL purchase [build failed] # purchase [purchase.test] ./vehicle_purchase.go:19:1: missing return '/usr/local/go/bin/go test --short --json .' returned exit code 1: exit status 1
```
This is because:

The reason the code block is failing with a "missing return" error in Go is that **not all possible execution paths of the function return a value**, while the function signature specifies a return type (`string`). 

Here's why:

- **The Problem:** The `if/else if` structure handles two scenarios: `option1 < option2` and `option2 < option1`. However, it doesn't explicitly handle the case where `option1 == option2` (meaning the strings are identical). In this scenario, neither the `if` nor the `else if` conditions would be met, and the function would reach its end without a return statement, resulting in the "missing return" error.
- **The Fix (second code block):** The working code adds an `else` block:
    
    ```Go
    else {
        return "They are both great options."
    }
    ```

- This `else` block acts as a fallback, guaranteeing a return value even if the `if` and `else if` conditions aren't met (i.e., when `option1 == option2`). This makes sure all possible paths in the function explicitly return a string, satisfying the function's return type declaration. 

In Go, if a function is declared to return a value, the compiler enforces that a return statement is present for every possible execution path through the function. This helps prevent unexpected behavior and ensures that the function always provides a value of the specified type.


# Exercise 7: String Formatting and Packages -- Party Robot

## Instructions and Info

## Introduction

### Packages

In Go an application is organized in packages. A package is a collection of source files located in the same folder. All source files in a folder must have the same package name at the top of the file. By convention packages are named to be the same as the folder they are located in.

```go
package greeting
```

Go provides an extensive standard library of packages which you can use in your program using the `import` keyword. Standard library packages are imported using their name.

```go
package greeting

import "fmt"
```

An imported package is then addressed with the package name:

```go
world := "World"
fmt.Sprintf("Hello %s", world)
```

Go determines if an item can be called by code in other packages through how it is declared. To make a function, type, variable, constant or struct field externally visible (known as _exported_) the name must start with a capital letter.

```go
package greeting

// Hello is a public function (callable from other packages).
func Hello(name string) string {
    return "Hello " + name
}

// hello is a private function (not callable from other packages).
func hello(name string) string {
    return "Hello " + name
}
```

### String Formatting

Go provides an in-built package called `fmt` (format package) which offers a variety of functions to manipulate the format of input and output. The most commonly used function is `Sprintf`, which uses verbs like `%s` to interpolate values into a string and returns that string.

```go
import "fmt"

food := "taco"
fmt.Sprintf("Bring me a %s", food)
// Returns: Bring me a taco
```

In Go floating point values are conveniently formatted with Sprintf's verbs: `%g` (compact representation), `%e` (exponent) or `%f` (non exponent). All three verbs allow the field's width and numeric position to be controlled.

```go
import "fmt"

number := 4.3242
fmt.Sprintf("%.2f", number)
// Returns: 4.32
```

You can find a full list of available verbs in the [format package documentation](https://pkg.go.dev/fmt).

`fmt` contains other functions for working with strings, such as `Println` which simply prints the arguments it receives to the console and `Printf` which formats the input in the same way as `Sprintf` before printing it.


## Instructions

Once there was an eccentric programmer living in a strange house with barred windows. One day he accepted a job from an online job board to build a party robot. The robot is supposed to greet people and help them to their seats. The first edition was very technical and showed the programmer's lack of human interaction. Some of which also made it into the next edition.

Implement the `Welcome` function to return a welcome message using the given name:

```go
Welcome("Christiane")
// => Welcome to my party, Christiane!
```


Implement the `HappyBirthday` function to return a birthday message using the given name and age of the person. Unfortunately the programmer is a bit of a show-off, so the robot has to demonstrate its knowledge of every guest's birthday.

```go
HappyBirthday("Frank", 58)
// => Happy birthday Frank! You are now 58 years old!
```


Implement the `AssignTable` function to give directions. It should accept 5 parameters.

- The name of the new guest to greet (`string`)
- The table number (`int`)
- The name of the seatmate (`string`)
- The direction where to find the table (`string`)
- The distance to the table (`float64`)

The exact result format can be seen in the example below.

The robot provides the table number in a 3 digits format. If the number is less than 3 digits it gets extra leading zeroes to become 3 digits (e.g. 3 becomes 003). The robot also mentions the distance of the table. Fortunately only with a precision that's limited to 1 digit.

```go
AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298)
// =>
// Welcome to my party, Christiane!
// You have been assigned to table 027. Your table is on the left, exactly 23.8 meters from here.
// You will be sitting next to Frank.
```


## My Solution

```Go 
package partyrobot
import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
    panic("Please implement the Welcome function")
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	return fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", name, table, direction, distance, neighbor)
}

```

## Learnings and Points to Note

1. In the first task there were three key learnings:
	1. Don't forget to import ```"fmt"```
	2. Don't forget to do ```fmt.Sprintf``` and not just  ```Sprintf```
	3. Sprintf does not print the function but rather returns it. In Go, `fmt.Sprintf` **returns the formatted string as a string value**.  Unlike `fmt.Printf` which directly prints to the standard output, `fmt.Sprintf` generates the formatted string and makes it available to be assigned to a variable or used in subsequent operations within the code. tl;dr: It provides a formatted string _result_ rather than printing it immediately.
		1. When I didn't type return before Sprintf the test failed
	4. If I use Printf instead of Sprintf, I get this error:
		1. ``` We received the following error when we ran your code: FAIL partyrobot [build failed] # partyrobot [partyrobot.test] ./party_robot.go:6:9: too many return values have (int, error) want (string) '/usr/local/go/bin/go test --short --json .' returned exit code 1: exit status 1```
		2. ```The error message "too many return values" in the code, `return fmt.Printf("Welcome to my party, %s!", name)`, means that the `fmt.Printf` function returns more values than the `Welcome` function is designed to handle. Since `fmt.Printf` returns two values (an `int` and an `error`), and the `Welcome` function expects only one (`string`), the Go compiler flags this as a "too many return values" error. To fix this, we use `fmt.Sprintf` instead, which returns only the formatted string, without printing it to the console.
	5.  **`fmt.Printf` returns two values:** The `fmt.Printf` function in Go is used to print formatted output to the console. It returns two values:
	    1. The number of bytes written to the standard output.
	    2. An error value (if any) encountered during the print operation.
		    1. **The `Welcome` function expects one return value:** The signature of the`Welcome` function is `func Welcome(name string) string`, which indicates that it's designed to return only a single value, and that value should be a string.```
2.  In the third function there are a few things to note:
	1. You cannot do multiple return functions. Initially I had three ```Sprintf``` functions to print the three sentences and returned all of them but that does not work.
	2. To print the integer to exactly three numbers with leading zeroes if necessary, we use:
		1. The %0Nd format specifier that allows us to print an integer to a specific number of digits, along with padding with leading zeros if necessary.
		2. This is used in conjunction with `fmt.Sprintf`.
		3. The key is to use the `%0Nd` verb, where: 
			1. `0` specifies that padding should be done with zeros instead of spaces.
			2. `N` is an integer representing the total desired width (number of digits) of the output string.
			3. `d` indicates that the argument is a decimal integer.
	3. Finally, it is critical to ensure that formatting is exactly what the test case needs. We needed to use the `\n` escape character to push to a new line for example.


# Exercise 8: Slices -- Card Tricks

## Instructions and Info

## Introduction
### Slices

Slices in Go are similar to lists or arrays in other languages. They hold several elements of a specific type (or interface).

Slices in Go are based on arrays. Arrays have a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view of the elements of an array.

A slice is written like `[]T` with `T` being the type of the elements in the slice:

```go
var empty []int                 // an empty slice
withData := []int{0,1,2,3,4,5}  // a slice pre-filled with some data
```

You can get or set an element at a given zero-based index using the square-bracket notation:

```go
withData[1] = 5
x := withData[1] // x is now 5
```

You can create a new slice from an existing slice by getting a range of elements. Once again using square-bracket notation, but specifying both a starting (inclusive) and ending (exclusive) index. If you don't specify a starting index, it defaults to 0. If you don't specify an ending index, it defaults to the length of the slice.

```go
newSlice := withData[2:4]
// => []int{2,3}
newSlice := withData[:2]
// => []int{0,1}
newSlice := withData[2:]
// => []int{2,3,4,5}
newSlice := withData[:]
// => []int{0,1,2,3,4,5}
```

You can add elements to a slice using the `append` function. Below we append `4` and `2` to the `a` slice.

```go
a := []int{1, 3}
a = append(a, 4, 2)
// => []int{1,3,4,2}
```

`append` always returns a new slice, and when we just want to append elements to an existing slice, it's common to reassign it back to the slice variable we pass as the first argument as we did above.

`append` can also be used to merge two slices:

```go
nextSlice := []int{100,101,102}
newSlice  := append(withData, nextSlice...)
// => []int{0,1,2,3,4,5,100,101,102}
```

### Variadic Functions

Usually, functions in Go accept only a fixed number of arguments. However, it is also possible to write variadic functions in Go.

A variadic function is a function that accepts a variable number of arguments.

If the type of the last parameter in a function definition is prefixed by ellipsis `...`, then the function can accept any number of arguments for that parameter.

```go
func find(a int, b ...int) {
    // ...
}
```

In the above function, parameter `b` is variadic and we can pass 0 or more arguments to `b`.

```go
find(5, 6)
find(5, 6, 7)
find(5)
```

Caution

The variadic parameter must be the last parameter of the function.

The way variadic functions work is by converting the variable number of arguments to a slice of the type of the variadic parameter.

Here is an example of an implementation of a variadic function.

```go
func find(num int, nums ...int) {
    fmt.Printf("type of nums is %T\n", nums)

    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            return
        }
    }

    fmt.Println(num, "not found in ", nums)
}

func main() {
    find(89, 90, 91, 95)
    // =>
    // type of nums is []int
    // 89 not found in  [90 91 95]

    find(45, 56, 67, 45, 90, 109)
    // =>
    // type of nums is []int
    // 45 found at index 2 in [56 67 45 90 109]

    find(87)
    // =>
    // type of nums is []int
    // 87 not found in  []
}
```

In line `find(89, 90, 91, 95)` of the program above, the variable number of arguments to the find function are `90`, `91` and `95`. The `find` function expects a variadic int parameter after `num`. Hence these three arguments will be converted by the compiler to a slice of type `int` `[]int{90, 91, 95}` and then it will be passed to the find function as `nums`.

Sometimes you already have a slice and want to pass that to a variadic function. This can be achieved by passing the slice followed by `...`. That will tell the compiler to use the slice as is inside the variadic function. The step described above where a slice is created will simply be omitted in this case.

```go
list := []int{1, 2, 3}
find(1, list...) // "find" defined as shown above
```
## Instructions

As a magician-to-be, Elyse needs to practice some basics. She has a stack of cards that she wants to manipulate.

To make things a bit easier she only uses the cards 1 to 10.

When practicing with her cards, Elyse likes to start with her favorite three cards of the deck: 2, 6 and 9. Write a function `FavoriteCards` that returns a slice with those cards in that order.

```go
cards := FavoriteCards()
fmt.Println(cards)
// Output: [2 6 9]
```

Return the card at position `index` from the given stack.

```go
card := GetItem([]int{1, 2, 4, 1}, 2) // card == 4
```

If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to return `-1`:

```go
card := GetItem([]int{1, 2, 4, 1}, 10) // card == -1
```

Note

By convention in Go, an error is returned instead of returning an "out-of-band" value. Here the "out-of-band" value is `-1` when a positive integer is expected. When returning an error, it's considered idiomatic to return the [`zero value`](https://www.geeksforgeeks.org/zero-value-in-golang/) with the error. Returning an error with the proper return value will be covered in a future exercise.

Exchange the card at position `index` with the new card provided and return the adjusted stack. Note that this will modify the input slice which is the expected behavior.

```go
index := 2
newCard := 6
cards := SetItem([]int{1, 2, 4, 1}, index, newCard)
fmt.Println(cards)
// Output: [1 2 6 1]
```

If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to append the new card to the end of the stack:

```go
index := -1
newCard := 6
cards := SetItem([]int{1, 2, 4, 1}, index, newCard)
fmt.Println(cards)
// Output: [1 2 4 1 6]
```

Add the card(s) specified in the `value` parameter at the top of the stack.

```go
slice := []int{3, 2, 6, 4, 8}
cards := PrependItems(slice, 5, 1)
fmt.Println(cards)
// Output: [5 1 3 2 6 4 8]
```

If no argument is given for the `value` parameter, then the result equals the original slice.

```go
slice := []int{3, 2, 6, 4, 8}
cards := PrependItems(slice)
fmt.Println(cards)
// Output: [3 2 6 4 8]
```

Remove the card at position `index` from the stack and return the stack. Note that this may modify the input slice which is ok.

```go
cards := RemoveItem([]int{3, 2, 6, 4, 8}, 2)
fmt.Println(cards)
// Output: [3 2 4 8]
```

If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to leave the stack unchanged:

```go
cards := RemoveItem([]int{3, 2, 6, 4, 8}, 11)
fmt.Println(cards)
// Output: [3 2 6 4 8]
```

## My Solution

```Go
package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
    // Solution 1
    var fav []int
    fav = []int{2,6,9}
    return fav
    /* Solution 2
    fav := []int{2,6,9}
    return fav
    */
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
    var x int
    y := len(slice)
    if index > y || index < 0{
    	return -1
    } else{
        x = slice[index]
        return x
    }
	panic("Please implement the GetItem function")
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
    var a []int
    y := len(slice)
    if index >= y || index < 0{
        a = append(slice, value)
        return a
    } else {
        slice[index] = value
        return slice
    }
    /*
     Alternate solution:
    y := len(slice)
    if index >= y || index < 0{
        slice = append(slice, value)
        return slice
    } else {
        slice[index] = value
        return slice
    }
    */
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
    prepend := append(values, slice...) //Need to use the unpack operator "..." after slice for the append command to work correctly.
    return prepend
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
    var y int
    y = len(slice)
    if index > y || index < 0{
        return slice
    } else {
        preSlice := slice[:index]
        postSlice := slice[index+1:]
        newSlice:= append(preSlice, postSlice...)
        return newSlice
    }
}

```

## Learnings and Points to Note
1. In the first function, there are two ways to do the task. One is to define the slice and then save it using the `slice name = []type{x,y,z}` method. The second way is to directly do `slice name := []type{x,y,z}`. And then return it. Both ways work perfectly well.
2. In the second function, there are two edge cases that need to be accounted for:
	1. If the index passed is greater than the length of the slice
	2. If the index passed is negative i.e. <0
	3. To deal with this we need to check the length of the slice
		1. I did not know how to do this and initially was going to use a for loop or something to that effect to check but Go offers a pre-built len() function of the form `func (v Type) int` that returns the length of the slice as an integer
		2. We use an if conditional to check if the index is greater than the length of the slice or lless than 0. If it is, a value of -1 is returned as an out-of-bounds error value 
		3. If not, we return the value within the slice at the passed index value
3. In the third function when I did this:
```Go
// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended. 
func SetItem(slice []int, index, value int) []int { 
	var a int 
	y := len(slice) 
	if index > y || index < 0{ 
	a = append(slice, value) 
	return a 
	} else { 
	slice[index] = value 
	return slice 
	} 
}
```

I got the following error:
```
FAIL cards [build failed] # cards [cards.test] ./card_tricks.go:34:13: cannot use append(slice, value) (value of type []int) as int value in assignment ./card_tricks.go:35:16: cannot use a (variable of type int) as []int value in return statement '/usr/local/go/bin/go test --short --json .' returned exit code 1: exit status 1
```

The reason is because I've declared a as an int and the append function returns a slice. Trying to assign the return value of the append function which is a slice to an integer leads to the error, as this a type mismatch and Go's type system does not allow it.
Furthermore, I then try to return a, which is an integer but my function signature expects me to return an `[]int` i.e. a slice of type integer. This again causes a type mismatch because of Go's type system that leads to an error.

When doing the if conditionals on this exercise, I initially did if `index > len(slice)`
However, if the index was equal to the length of the slice, then appending it would be necessary. In the second function, a read was happening, so index = len(slice) would work and the read would be successful. But in the third function, it is a write operation. In a slice of one element, len(slice) would be 1 but index would be 0. So if I passed an index of 1, which is equal to the length of the slice, it would need to be appended as opposed to replacing a value.

4. In the fourth function, we are given to slices and we need to prepend the second one onto the first. Initially, I did `prepend :=append(values, slice)` but that resulted in the following error: `FAIL cards [build failed] # cards [cards.test] ./card_tricks.go:55:31: cannot use slice (variable of type []int) as int value in argument to append '/usr/local/go/bin/go test --short --json .' returned exit code 1: exit status 1`
	1. The reason for this because the append function expects the first argument to be the slice on which the append operation occurs and the subsequent arguments to be individual elements that will be appended or elements from another slice that are unpacked.
	2. When I use `append(values, slice)`, Go interprets `slice` (which is of type `[]int`, a slice of integers) as a single element that we are trying to append to `values`. However, `values` is also a slice of type  (`[]int`), so it expects `int` values as additional arguments, not another `[]int`. This is why we get the error: "cannot use slice (variable of type []int) as int value in argument to append". We are trying to pass a whole slice where an individual integer element is expected.
	3. The solution is to then use `prepend := append(values, slice...)` 
		1. The `...` operator is called the unpack/spread operator. It unpacks the elements of the slice into individual arguments. So if slice was `[]int{1,2,3}` then passing `slice...` would be equivalent to passing `1, 2, 3` as separate arguments to `append` .
		   When we write `append(values, slice...)`, we are effectively saying: "Append all the individual elements from `slice` to the `values` slice." This is the correct way to combine two slices where the elements of the second slice should be added individually to the first.
5. For the fifth and final function, we have to delete an element from the slice and then return the new slice. My approach was to create a new slice from the existing slice by getting a range of elements. I used the square bracket notation and specified the first slice as `[:index]`. The implication was that the 1st element with index 0 would be included and all subsequent elements till the index element would be included but the index element itself would not be included. I then created another slice with of the form `[index+1:]`. This means that the element after the index element until the last element would be included. Appending the two lists (while ensuring that the second list is followed by the unpacking/spread operator `...` since the append function expects it) would lead to the element whose index was passed being deleted.


# Exercise 9: Conditionals Switch -- Blackjack

## Instructions and Info

## Introduction

Like other languages, Go also provides a `switch` statement. Switch statements are a shorter way to write long `if ... else if` statements. To make a switch, we start by using the keyword `switch` followed by a value or expression. We then declare each one of the conditions with the `case` keyword. We can also declare a `default` case, that will run when none of the previous `case` conditions matched:

```go
operatingSystem := "windows"

switch operatingSystem {
case "windows":
    // do something if the operating system is windows
case "linux":
    // do something if the operating system is linux
case "macos":
    // do something if the operating system is macos
default:
    // do something if the operating system is none of the above
} 
```

One interesting thing about switch statements, is that the value after the `switch` keyword can be omitted, and we can have boolean conditions for each `case`:

```go
age := 21

switch {
case age > 20 && age < 30:
    // do something if age is between 20 and 30
case age == 10:
    // do something if age is equal to 10
default:
    // do something else for every other case
}
```

## Instructions

In this exercise we will simulate the first turn of a [Blackjack](https://en.wikipedia.org/wiki/Blackjack) game.

You will receive two cards and will be able to see the face up card of the dealer. All cards are represented using a string such as "ace", "king", "three", "two", etc. The values of each card are:

|card|value|card|value|
|---|---|---|---|
|ace|11|eight|8|
|two|2|nine|9|
|three|3|ten|10|
|four|4|jack|10|
|five|5|queen|10|
|six|6|king|10|
|seven|7|_other_|0|

**Note**: Commonly, aces can take the value of 1 or 11 but for simplicity we will assume that they can only take the value of 11.

Depending on your two cards and the card of the dealer, there is a strategy for the first turn of the game, in which you have the following options:

- Stand (S)
- Hit (H)
- Split (P)
- Automatically win (W)

Although not optimal yet, you will follow the strategy your friend Alex has been developing, which is as follows:

- If you have a pair of aces you must always split them.
- If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, a face card (Jack/Queen/King) or a ten then you automatically win. If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
- If your cards sum up to a value within the range [17, 20] you should always stand.
- If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.
- If your cards sum up to 11 or lower you should always hit.

Implement a function to calculate the numerical value of a card:

```go
value := ParseCard("ace")
fmt.Println(value)
// Output: 11
```

Write a function that implements the decision logic as described above:

```go
func FirstTurn(card1, card2, dealerCard string) string
```

Here are some examples for the expected outcomes:

```go
FirstTurn("ace", "ace", "jack") == "P"
FirstTurn("ace", "king", "ace") == "S"
FirstTurn("five", "queen", "ace") == "H"
```

## My Solution

```Go
package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
    case "ace":
        return 11
    case "two":
        return 2
    case "three":
        return 3
    case "four":
        return 4
	case "five":
        return 5
	case "six":
        return 6
    case "seven":
        return 7
    case "eight":
        return 8
    case "nine":
        return 9
    case "ten":
        return 10
    case "jack":
        return 10
    case "queen":
        return 10
    case "king":
        return 10
    default:
    	return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	Card1 := ParseCard(card1)
    Card2 := ParseCard(card2)
    Card3 := ParseCard(dealerCard)
    if Card1 == 11 && Card2 == 11{
        return "P"
    } else if (Card1 + Card2) == 21 && (Card3 != 11 && Card3 != 10){
        return "W"
    } else if (Card1 + Card2) == 21 && (Card3 == 11 || Card3 == 10){
        return "S"
    } else if (Card1 + Card2) >= 17 && (Card1 + Card2) <= 21{
        return "S"
    } else if (Card1 + Card2) >= 12 && (Card1 + Card2) <= 16 && Card3 < 7{
        return "S"
    } else if (Card1 + Card2) >= 12 && (Card1 + Card2) <= 16 && Card3 >= 7{
        return "H"
    } else if (Card1 + Card2) <= 11 {
        return "H"
    }
	return "S"
}

```

## Learnings and Points to Note

1. Suspect OR error: 

```markdown
### We received the following error when we ran your code:
FAIL	blackjack [build failed]
# blackjack
# [blackjack]
./blackjack.go:45:41: suspect or: Card3 != 11 || Card3 != 10
'/usr/local/go/bin/go test --short --json .' returned exit code 1: exit status 1
```

This was because of this line in the code: `else if (Card1 + Card2) == 21 && (Card3 != 11 || Card3 != 10){`

As there are two not statements that comprise the OR condition, it will always evaluate to true. This is because OR just needs one of its operands to be true for it to be true. For e.g. 
- Consider `Card3` is 5: `5 != 11` is `true`. `5 != 10` is `true`. So, `true || true` is `true`
- Consider `Card3` is 11: `11 != 11` is `false`. `11 != 10` is `true`. So, `false || true` is `true`
- Consider `Card3` is 10: `10 != 11` is `true`. `10 != 10` is `false`. So, `true || false` is `true`
No matter what value `Card3` has, one of the conditions `Card3 != 11` or `Card3 != 10` will _always_ be true, causing the entire `||` expression to always evaluate to true. This is why the Go compiler flags it as "suspect or", because it likely indicates a logical flaw where the condition will always be met.

The fact that the compiler can detect a logical flaw like this is mind-blowing to me. Beautiful stuff man.

Also, important to have a default return at the end so the compiler does not throw an `expected return error.`

The hint for task 2 suggested a bunch of stuff like using switch statements with nested ifs or using functions to handle tasks. Eg.:

1. "Use a big switch statement on the player score (maybe with nested if-statements on the dealer-score in some cases)"
	1. This approach suggests that you could primarily use a `switch` statement based on the player's total score (the sum of `Card1` and `Card2`). Within each `case` of the switch (e.g., for a player score of 21, for scores between 17 and 20, etc.), you might need to use `if` or `else if` statements to further check the `dealerCard` to determine the final decision.
```Go
func FirstTurn(card1, card2, dealerCard string) string {
    playerScore := ParseCard(card1) + ParseCard(card2)
    dealerScore := ParseCard(dealerCard)

    switch playerScore {
    case 21:
        if dealerScore != 11 && dealerScore != 10 {
            return "W"
        } else {
            return "S"
        }
    case 17, 18, 19, 20: // Using a range of values for a case
        return "S"
    case 12, 13, 14, 15, 16:
        if dealerScore < 7 {
            return "S"
        } else {
            return "H"
        }
    default: // Covers scores <= 11 and also handles the pair of aces implicitly if you place it last.
        return "H"
    }
}
```

2. "Distinguish separate player score categories... and write separate functions for all (or some) of these categories."
	1. This approach encourages you to break down the decision logic into smaller, more focused functions. You would first determine which category the player's score falls into (small, medium, or large hands). Then, based on that category, you would call a dedicated function to make the final decision, possibly passing in the dealer's card.

```Go
func handleSmallHand(dealerScore int) string {
    return "H" // For scores <= 11
}

func handleMediumHand(playerScore, dealerScore int) string {
    if playerScore >= 12 && playerScore <= 16 && dealerScore < 7 {
        return "S"
    } else {
        return "H"
    }
}

func handleLargeHand(playerScore, dealerScore int) string {
    if playerScore >= 17 && playerScore <= 20 { // Stand for 17-20
        return "S"
    }
    // Handle special cases within large hand if needed
    // e.g., for 21 with dealer having a 10 or Ace
    if playerScore == 21 && (dealerScore == 11 || dealerScore == 10){
        return "S"
    } else if playerScore == 21 && (dealerScore != 11 && dealerScore != 10){
        return "W"
    }
    return "S" // Default to stand for other large hands, might need refinement based on exact rules
}

func FirstTurn(card1, card2, dealerCard string) string {
    playerScore := ParseCard(card1) + ParseCard(card2)
    dealerScore := ParseCard(dealerCard)

    if ParseCard(card1) == 11 && ParseCard(card2) == 11 { // Check for a pair of aces first
        return "P"
    } else if playerScore <= 11 {
        return handleSmallHand(dealerScore)
    } else if playerScore >= 12 && playerScore <= 16 {
        return handleMediumHand(playerScore, dealerScore)
    } else { // playerScore >= 17
        return handleLargeHand(playerScore, dealerScore)
    }
}

```

Supposedly, the hint is pushing me towards:

- **Switch Statements:** To potentially make the code more readable when we have multiple, distinct cases based on a single value (like `playerScore`).
- **Function Decomposition:** To break down the logic into smaller, more manageable units. This can be particularly helpful if the decision logic becomes more complex in the future, as it makes the code easier to test, understand, and modify.
- **Readability:** A `switch` statement on the `playerScore` can make it easier to see at a glance how different player scores lead to different actions. Separating into functions for score categories also improves readability.
- **Maintainability:** If the Blackjack rules change (e.g., new rules for a specific score range), it's easier to modify a smaller, focused function or a specific `case` within a `switch` statement rather than navigating through a long chain of `if...else if` statements.
- **Testability:** Smaller functions are easier to write unit tests for. You can test `handleSmallHand`, `handleMediumHand`, and `handleLargeHand` independently, reducing the complexity of testing the entire `FirstTurn` function.
