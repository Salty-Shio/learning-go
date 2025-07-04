# GoLang - An introduction by video tutorials

These notes were taken primarily taken while watching **Learn Go Fast: Full Tutorial** by *Alex Mux* with supplmentation from various sources.

The code generated for each sement of the video except for the last can be found in `./tutorials` and the project at the end of the video can be found in `./project`.

## Introduction

`Go` is an open source language created by Google. The language itself was created to improve the efficiency of programmers. It was designed with multicore, networked machines, and large codebases in mind. While this is true, it is also a general use lanaguage, which can be used for a variety of use cases. Because it was built with multicore machines in mind, it is especially useful for concurrency. `Go` is a compiled language, making it faster in comparisson to interpreted languages such as Python, and compiles directly to `machine code`, making it comparable to `C` or `C++`. It also has memory management and garbage collection built in to avoid memory leaks.

Getting more specific:

1. `Go` is a Statically Typed Language
   - This means that variables, when defined, will have its type declared explicitly, or will infer the type based on the initial value passed in

```Go
// Defined variable without initializing it to a value. The type is explicilty defined.
var numberOne int

// Defined variable by initializing it to a value. The type is inferred.
var numberTwo 0
```

2. `Go` is a strongly typed language

   - This means that variables cannot be assigned values that are not their type, i.e. an integer cannot be assigned a string. This also means that when attempting to perform operations on variables, certain variables cannot interact. This is in direct contrast to a language like `JavaScript` where we can add a string to a float and it will continue along as if nothing has happened. This allows the compiler to do checks and error detection before compiling our code, as well as allows our editor to have better autocomplete because it is clear what variables are.

3. `Go` is Compiled

   - This means that our program will be compiled down to a standalone binary. This allows a program to be run in multiple enviornments. This is in contrast to Python which interprets the code at run time. This also means that programs will run much faster as there is no interpreter interpretting the code as it runs, rather the code is already ready to be run by the machine.

4. `Go` has built in Concurrency

   - There are built in packages that support concurrency. This is because `Go` was built with this features specifically in mind.

5. `Go` was built with simplicity in mind
   - The syntax of `Go` is consistant and simple, making writing code faster, and debugging easier.

### Installation

Go to the website and install `Go` by following the prompts: [go.dev](https://go.dev/doc/install)

To test if the install was successfull go to a terminal and run `go version`.

## Modules and Packages

`Go` works within a paradigm of packages and modules.

A _Package_ is a folder containing various `.go` files.

A _Module_ is a collection of packages.

When starting a `Go` project, we are just creating a module.

To get started run: `go mod init {MODULE NAME}`. This initializes the module by creating a `go.mod` file which will contain the version of `Go` and the module name.

## Tutorial 1 - Writing and Running a program

Every go file is part of a package. This means to start we start every file with the package name.

```go
package main
```

`main` is a special package name telling Go that the compiller to look for the entrypoint function within the main package. The syntax for the entrypoint function is as follows:

```go
func main() {

}
```

If the package name is not called main, this is not a requirement for the package to have.

To import a package we can use the `import` keyword followed by the package name in quotes.

```go
import "fmt"
```

This will now cause issue because `Go` _does not let us import a module without using it_. To complete our program we can now use the print function in `Go` to print out **Hello World** to the console.

```go
fmt.Println("Hello, World!")
```

### Building and running our file

To build and run this file we use the following:

```bash
go build {FILE PATH}
```

This produces a file titled `main` which is the binary for the program. We can run it in the console to see what it does. Alternatively we can use:

```bash
go run {FILE PATH}
```

This will build and run the code in one command.

## Tutorial 2 - Data Types, Variables, and Consts

To declare a variable use the `var` keyword, followed by the name, and then the type, such as `int`;

```go
var integer int
```

### Integers

`int` will default to a specific value based on the system. There are ways to specify the the size and type of int which can be seen in this table:

| Keyword  | Bits     | Min Value                  | Max Value                  |
| -------- | -------- | -------------------------- | -------------------------- |
| `int`    | 32 or 64 | -2^31 or -2^63             | 2^31-1 or 2^63-1           |
| `int8`   | 8        | -128                       | 127                        |
| `int16`  | 16       | -32,768                    | 32,767                     |
| `int32`  | 32       | -2,147,483,648             | 2,147,483,647              |
| `int64`  | 64       | -9,223,372,036,854,775,808 | 9,223,372,036,854,775,807  |
| `uint`   | 32 or64  | 0                          | 2^32-1 or 2^64-1           |
| `uint8`  | 8        | 0                          | 255                        |
| `uint16` | 16       | 0                          | 65,535                     |
| `uint32` | 32       | 0                          | 4,294,967,295              |
| `uint64` | 64       | 0                          | 18,446,744,073,709,551,615 |

It should be noted as well, `int` initialize to _0_ by default which in contrast to `JavaScript`'s `undefined` or `C++` having a garbled mess of whatever was there before we initialized the value.

The max and mins are also something to look out for. If we try to initialize a value over its max size, the compiller will throw an error. If we update a value however, the compiller is no longer involved, and there will be no error. This can lead to unexpected behavior.

```go
var numberOne int8 = 127 // No error

var numberTwo int8 = 128 // Overflow error

numberOne = numberOne + 1 // No error
```

Integer division results in an integer, which rounds any integer down. If we want the remainder, we can get it using the modulo operator `%`.

### Floats

`Go` has two types of floats. We *must* specify the float type when declaring a float.

| Keyword   | Bits | Min Value             | Max Value            |
| --------- | ---- | --------------------- | -------------------- |
| `float32` | 32   | approx -3.402823e+38  | approx 3.402823e+38  |
| `float64` | 64   | approx -1.797693e+308 | approx 1.797693e+308 |

```go
var decimalOne float // Throws error because the type of float is undefined

var decimalTwo float32 = 3.14159
```

Once again, be on the lookout for the same issues with floats as integers, and remember that the default intialization is _0_.

### Type Casting

Now that we have two types of numbers, we can address type casting. We cannot directly add an `int` and a `float` together because they are different types. Because of this we need to type cast from an int to a float or vice versa to avoid and error.

```go
var integer int32 = 2
var float float32 = 2.0

var fail float32 = float + integer // Not quick maths

var success float32 = float + float32(integer) // 2.0 + 2.0 is 4.0
var result int32 = int32(success) - 1 // -1 that's 3, quick maths.
```

### String

The `string` keyword is used for strings. Depending on what we use to define the string, the code will allow different behaviors. For instance, the use of `" "` means that the string is single line terminating.

```go
var stringOne string = "Hello World!" // Valid

var stringTwo string = "Hello
World" // Not valid

var stringThree string = "Hello\nWorld" // Valid
```

Strings can be added together for concatenation.

```go
var stringOne string = "Hello"
var stringTwo string = "World"

var stringThree string = stringOne + " " + stringTwo
```

To get the length of the string we can try using the built in `len()` function. It should however be noted that this is not always the number of charcters in the string. This is because `len()` returns the number of bytes used. Becuase `Go` encodes in `UTF-8`, characters outside of this encoding will use more bytes, thereby shifting the result of the `len()` funciton.

```go
var stringOne string = "A"
var stringTwo string = "γ"

fmt.Println(stringOne) // 1
fmt.Println(stringTwo) // 2
```

To get around this, if we expect strings with characters outside of `UTF-8`, import the `"unicode/utf8"` and use the following:

```go
import "unicode/utf8"

fmt.Println(utf8.RuneCountInString(stringTwo)) // 1
```

This works becasue a `rune` is a datatype in go representing characters. A rune is actually defined as follows using the `' '`:

```go
var character rune = 'a'
```

The default value for a rune is `'0'`.
The default value for a string is `" "`.

### Booleans

Booleans are booleans are booleans are booleans.

```go
var george bool = true
```

The default is `false`

### Variables

There are a number of ways to define variables

```go
var firstName = "Louis" // Type inference
lastName := "Armstrong" // Short variable declaration operator

var num1, num2 int32 = 1, 2 // Multiple variables at once
var num3, num4 = 1, 2 // Hey look, type inference still works
num5, num6 := 1, 2 // this even still works
```

It should be noted that it best practice to use the `:=` when it is obvious what is being assigned, and the `var {NAME} type` syntax when it is not apparent what is being assinged.

```go
x := foo() // What is x?

var y string = bar() // This is clear
```

### Constants

These are variables that don't vary. In other words, we can assign them, but once assigned, they don't change. We also have to assign a value to a const at declaration unlike variables where there is a default value.

```go
const name string // Error

const name string = "Dick"

name = "Richard" // Cannot assign to a const
```

## Functions and Control Structures

### Functions

As shown earlier, the function sytnax in `Go` is:

```go
func functionName() {
   fmt.Println("Hello, World!")
}
```

It should be noted that `Go` enforces the K&R brace style for functions. This means that Allman style braces are invalid sytax.

```go
func functionName()
{  // This throws an error

}
```

For passing parameters, the name of the parameter comes first, followed by the type declaration. This also means that we cannot pass other data types into the function. For instance, if we wanted to print an integer and we have a function that takes in a string, we will first have to cast the value to a string before passing it in.

```go
func functionName(value string) {
   fmt.Println(value)
}
```

These function types have had void return type so far, but how about function that return things? We must also declare what value the function will be returning so that `Go`'s compiler can check for errors. We can also add multipler return values as shown below.

```go
func intDivision(a int, b int) (int, int) {
	var result int = a / b
	var remainder int = a % b
	return result, remainder
}
```

Now we can ask, what happens when we hit a division by zero error? `Go` throws a panic error and exits out of the program. This is not great as far as handling errors goes so what is standard practice is to add a return type `error` along with our function if our function can encounter errros. When an `error` is first defined it will default to a value of `nil`. This also introduces the syntax for if statements as a check will occur to see if the divisor is 0.

If statements are similar to most other languages. The one thing that is noteworthy with `Go` is that parentheses are not necessary for the conditional. This contrasts `JavaScript` where the conditional must be wrapped in parentheses.

```go
func intDivision(a int, b int) (int, int, error) {
   var err error
   if b == 0 {
      err = error.New("Cannot divide by zero.")
      return 0, 0, err
   }
	var result int = a / b
	var remainder int = a % b
	return result, remainder
}
```

#### Function not defined?

Because `Go` is compiled, no matter where it is placed, so long as the function is declared in the same package, `Go` will know about the function. In other words, so long as at some point the function is delcared in the same scope, the function will be defined and we will not encounter an error where the function has not yet been defined. This is because when `Go` compiles the language it has already processed the entire file.

### If statements

Finally if, else if, and else statements all follow the standard syntax accordingly. The else and else if statements **must** be on the smae line as the final bracket of the if statement precessing them.

```go
if condition1 {

} else if condition2 {

} else {

}
```

#### Operators in Go

With conditionals and talk of arithmetic, I feel that it is pertinent to list the standard operators in `Go`. The set of operators is pretty standard as far as programming languages _go_.

| Operator | Description                                      |
| :------- | :----------------------------------------------- |
| `+`      | Addition                                         |
| `-`      | Subtraction                                      |
| `*`      | Multiplication                                   |
| `/`      | Division                                         |
| `%`      | Modulus (remainder)                              |
| `&`      | Bitwise AND                                      |
| `\|`     | Bitwise OR                                       |
| `^`      | Bitwise XOR                                      |
| `&^`     | Bitwise clear (AND NOT)                          |
| `<<`     | Left shift                                       |
| `>>`     | Right shift                                      |
| `==`     | Equal to                                         |
| `!=`     | Not equal to                                     |
| `<`      | Less than                                        |
| `<=`     | Less than or equal to                            |
| `>`      | Greater than                                     |
| `>=`     | Greater than or equal to                         |
| `&&`     | Logical AND                                      |
| `\|\|`   | Logical OR                                       |
| `!`      | Logical NOT                                      |
| `+=`     | Add and assign                                   |
| `-=`     | Subtract and assign                              |
| `*=`     | Multiply and assign                              |
| `/=`     | Divide and assign                                |
| `%=`     | Modulo and assign                                |
| `&=`     | Bitwise AND and assign                           |
| `\|=`    | Bitwise OR and assign                            |
| `^=`     | Bitwise XOR and assign                           |
| `&^=`    | Bitwise clear and assign                         |
| `<<=`    | Left shift and assign                            |
| `>>=`    | Right shift and assign                           |
| `++`     | Increment (postfix only)                         |
| `--`     | Decrement (postfix only)                         |
| `=`      | Assignment                                       |
| `<-`     | Channel send/receive                             |
| `.`      | Selector (access struct fields, package members) |
| `&`      | Address of                                       |
| `*`      | Dereference                                      |

#### Shortcircuiting

`Go` also has shortcircuiting behavior when it comes to evaluating boolean expressions. This means that for the following `x != 0 || y > 10` if `x = 1` it will not evalute the second condition in the conditional.

### Switch statements

Switch statements are another tool we can use to avoid evaulating a larger number of conditionals. In `Go` breaks are implied, so breaking out of each case with a `break` keyword is not necessary.

```go
func main() {
	var numerator int = 10
	var denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator)

	switch {
		case err != nil:
			fmt.Println("Error occurred:", err)
		case remainder == 0:
			fmt.Printf("The result is %v with no remainder\n", result)
		default:
			fmt.Printf("The result is %v with a remainder of %v\n", result, remainder)
	}
}
```

Note, that this is an expressionless switch case statment, and thus it effectivly is a simpler definiton of an `if-else if-else` ladder, where if statements are placed in series. In other words, each case is evaluated sequentially. In some cases, `Go` may use a binary search to optimize switch cases if the cases are constant and there are enough cases to warrant a search. Aside from this however, I find that this is a more readable approach to switch cases statements, though it is not always practical to write a switch case this way.

```go
switch condition {
   case 1:
      foo()
   case 2:
      bar()
   case 3:
      isThatATypeOfChoclate()
   default:
      itShouldBe()
}
```

## Arrays, Slices, and Maps

### Arrays

Arrays in Go are similar to that of an array in `C` or `C++` in that they have a set size and a set data type associated with them. Arrays are zero indexed.

```go
var intArr [3]int32
// Index
fmt.Println(intArr[0])

// Slice
fmt.Println(intArr[1:3])
```

Arrays store their contents in continguous memory segments. This can be shown by printing the pointers to each item in the array. This can be done using the address operator `&` before the call to the array.

```go
var intArr [3]int32

// Show addressess
fmt.Println(&intArr[0])
fmt.Println(&intArr[1])
fmt.Println(&intArr[2])
```

It should be noted that this is similar behavior to a language such as `C` which stores an array as a pointer to a location in memory, and indexes through the array by incrememting the memory location by the index.

We can also immediately initialize an array using the following syntax:

```go
var intArray [3]int32 = [3]int32{1, 2, 3}

// And of course there is shorthand such as the following
var intArray = [3]int32{1, 2, 3}
intArray := [3]int32{1, 2, 3}
intArray := [...]int32{1, 2, 3} // The ... makes the compiler infer the size of the array.

```

The `...` is called the variadic parameter operator, and in `Go` is used to denote that a function or template can accept a varialbe number of arguments.

### Slices

Slices are a wrapper for arrays. In other words, a slice is a superset of the features of an array. The way that slice is defined is similar to an array, though it is not the same. Instead of adding a parameter for the size of the array, we leave the value empty. This creates a slice instead of an array. This opens up a variety of new methods, such as the append feature.

```go
var intSlice []int32 = []int32{4 5 6}
fmt.Printf("The length is %v with capacity of %v", len(intSlice), cap(intSlice))
// Cap is a built in function allowing us to see what the capacity of a slice is.

intSlice = append(intSlice, 7)
fmt.Printf("The length is %v with capacity of %v", len(intSlice), cap(intSlice))
```

The underlying workings of a slice are as follows:

1. A new array with the exact length is allocated in memory.
2. When running the append, the method checks if there is enough room for a new value.
3. If there is not enough room, a new array with increased capacity will be allocated (this could be larger than needed, such as doulbing capacity to add a single value).
4. The appended value is copied in, followed by copying in the values from the previous array.

In other words, a new array in a new location in memory is returned after these operations. We _should not assume that the location remains the same while using a slice._

Interestingly we cannot index into the unassigned slots in the slice. So if there are six values in a slice with a capacity of ten, we will get an `index out of range` error.

We can also add multiple values to the slice at once depicted below.

```go
throwAway := []int32{8,9}
intSlice = append(intSlice, throwAway...) // Used after a value, this will spread out the values contained in the array/slice.
```

We can also use the `make` syntax to create a slice. This can be good if we know that there is a maximum size that our slice will take up and we want to avoid reallocating for resizing.

```go
var intSlices3 []int32 = make(int32[], 3, 8)
```

### Maps

As with other languages, this is a table of `key:value` pairs. Here is the syntax for defining one. We can initialize with or without values. Referencing values is fairly standard.

```go
var myMap map[string]uint8 = make(map[string]uint8)
// Or
var myMap = map[string]unint8 = {"Adam": 23, "Sarah":45}

fmt.Println(myMap["Adam"])
```

When trying to access a value that does not exist (i.e. the key leads to nothing) the map will return the default value of the datatype for the value. For `unint8` that would be 0. Maps also return a second, _ok_ value, which is a boolean. This is true if the key was found, and false otherwise.

```go
var myMap = map[string]unint8 = {"Adam": 23, "Sarah":45}
age, ok = myMap["Adam"] // returns 23, true
age, ok = myMap["James"] // returns 0, false
```

Deleting a value is as follows:

```go
delete(map, key);
```

Maps do not preserve order in `Go` so iterating over them will result in different orderings each time.

#### Iterating

We can iterate over maps, arrays, and slices using a `range`. This will return an iterator allowing us to iterate through each value (similar to python).

```go
for name:= range myMap {
	fmt.Printf("Name: %s, Age: %d\n", name, myMap[name])
}
```

For an array we can do the following.

```go
for index, value := range intSlice {
	fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

Interestingly, there is no dedicated `while` loop in `Go`, rather, the effect of a while loop can be achieved with the for syntax as well.

```go
// Stnadard for loop
for i:=0; i<10; i++ {
   fmt.Println(i)
}

// Standard while loop
i := 0
for i < 10 {
   i++
   fmt.Println(i)
}

// No condition needed
i = 0
for {
   if i >= 10 break
   fmt.Println(i)
}
```

## Strings, Runes, and Bytes

`Strings` can be indexed just like arrays. This is done using the `string[index]` syntax. It should be noted, that this will return numerical values. This is because the characters are encoded in `UTF-8`. The computer will read out the values as the numeric representation of the binary values rather than the encoded value.

```go
var myString = "résumé"
var indexed = mystring[0]
fmt.Println(indexed)
```

`Strings` allocate a byte array for each character. This means that it allocates the necessary memory to hold each character in the `string`. This is useful to know for a number of reasons. For instance, in the word _résumé_, there are two characters that are represented using 2 bytes, rather than one. _KEEP THIS IN MIND._ Because certain characters have different sizes we can get a number of differing results from our `string` without understanding why.

```go
var myString = "résumé"
var indexedFirst = myString[1]
fmt.Println(indexed) // Outputs 195, becasue the first 8 bits are stored at index 1 of the string. In ASCII this is Ã which is not in the string.

var indexedSecond = myString[2]
fmt.Println(indexed) // Outputs 169, the second 8 bits of the character at index 1. This correlates to the ASCII © symbol which is also not in the string.

for i, v := range myString {
   fmt.Println(i, v)
}
// Returns the following list
// 0 114
// 1 223
// 3 115
// 4 117
// 5 109
// 6 233
```

It should now be apparent that there are some oddities working with `strings`. For instance, we can see that indexing straight to a position does not always give us the correct value. We also can see that using the range on myString does some work in decoding the proper positions of the values. This can be see by noticing that there is never an index of 2 printed out. That is because it is part of the character for position 1. We also know that iterating through the `string` we are not seeing the full size of the array. For instance, we know that there are 8 slots in the array, but only 7 are shown because the final slot, once again, is part of the character in slot 6. This is also why taking the `len(string)` returns the number of bytes and not the number of charcters, because a `string` is actually an array of bytes representing the `string`. It should also be noted that the type of each value at each point on the array, because it is a byte, is an `int8`.

To get around this odd behavior it is possible to instead cast a `string` to an array of `runes` (`runes` are the `chars` of `Go`). This changes the underlying representation from an array of bytes representing the string to an array of unicode characters representing the string. In other words, what was output from the iteration in the previous code block is how the string would be represesnted

```go
var mystring = []rune("résumé")
var indexed  = myString[2]
fmt.Println(indexed) // Output 115, the Unicode representation of 's'
```

It should be noted that `rune` is an alias for `int32`.

Earlier it was noted that we can concatenate `strings` in `Go` using the `+` operator. This does work but there is a caviate to this. `Strings` are immutable and as such, each string that is concatenated together is actually a new string altogether, which is not efficient.

There is also a strings library in `Go` to help us create and work with strings more efficiently. For instance:

```go
// Lame normal concatenation
var bString = ""
var stringSlice = []string{"H", "e", "l", "l", "o" }
for i := range strSlice {
   bString += strSlice[i] // Creates a new string each iteration.
}

import "strings"

// Chad efficient string building
var strBuilder strings.Builder
for i := range strSlice {
   strBuilder.WriteString(strSlice[i])
}
var aString = strBuilder.String()
```

The string library allocates an array internally and appends values until the string method is called.

## Structs and Interfaces

### Structs

#### Defining a Struct

Like in many other languages, `Go` allows the programmer to define their own types. These types can then be used, referenced and manipulated, like in many other languages.

```go
type gasEngine struct {
   mpg uint8
   gallons uint8
}

var engine gasEngine
fmt.Println(engine.mpg, engine.gallons) // Prints 0 0
```

There are also multiple ways to assign values to the attributes of a given struct.

```go
var engine1 gasEngine{mpg: 100, galons: 1} // Struct Literal Syntax
fmt.Println(engine1.mpg, engine1.gallons) // Prints 100 1

var engine2 gasEngine{20, 20} // Assigns attributes values in order
fmt.Println(engine2.mpg, engine2.gallons) // Prints 20 20

var engine3 gasEngine
engine3.mpg = 1
engine3.gallons = 100
fmt.Println(engine3.mpg, engine3.gallons) // Prints 1 100
```

Because a struct is just another type, structs can have other structs as attribtues.

```go
type make struct {
   name string
}

type gasEngine struct {
   mpg uint8
   gallons uint8
   manufacturer make
}

var chevyCruze gasEngine = gasEngine{30, 16, make{"Chevrolet"}}
fmt.Println(chevyCruze.mpg, chevyCruze.gallons, chevyCruze.manufacturer.name) // Prints 30 16 Chevrolet

// This creates the following object
// {
//    mpg: 30
//    gallons: 15
//    manufacturer.name: "Chevrolet"
// }
```

#### Embedding: A GoLang Composition Alternative

On top of traditional composition, there is also *embedding*. This allows for easier composition of structs. In the tradional composition form as shown above, one would get the name of the manufacturer by accessing `gasEngine.manufacturer.name`. Embedding instead allows the attributes of a lower order struct to be promoted to part of the higher order struct. This allows us to access the attributes as if they are declared directly in the higher order struct.


```go
type make struct {
   name string
}

type gasEngine struct {
   mpg uint8
   gallons uint8
   make
}

var chevyCruze gasEngine = gasEngine{30, 16, make{"Chevrolet"}}
fmt.Println(chevyCruze.mpg, chevyCruze.gallons, chevyCruze.name) // Name is now directly accessed through the name attribute.


// This creates the following object
// {
//    mpg: 30
//    gallons: 15
//    name: "Chevrolet"
// }
```

#### Anonymouse Structs

We can also create anonymouse structs. This requires us to define, declare, and initialize the struct all at once. This is similar to anonymous functions.

```go
var anonEngine = struct {
   mpg uint8
   gallons uint8
}{25, 15}
```

#### Methods

Structs can also have methods. These methods have access to the values of the instance of the struct.

```go
// Define the struct
type gasEngine struct {
   mpg uint8
   gallons uint8
}

// Define a function => tie the method to the struct and give the instance a name => method name => return type
func (engine gasEngine) milesLeft() uint8 {
   return engine.gallons * engine.mpg
}

var engine = gasEngine{10, 20}
fmt.Println(engine.milesLeft()) // Returns 200
```

### Interfaces

```go
// Gas Engine Definition
type gasEngine struct {
   mpg uint8
   gallons uint8
}

func (engine gasEngine) milesLeft() uint8 {
   return engine.gallons * engine.mpg
}

// Electric "Engine" definition
type electricEngine struct {
   mpkwh uint8
   kwh uint8
}

func (motor electricEngine) milesLeft() uint8 {
   return motor.mpkwh * motor.kwh
}

func canMakeIt(e gasEngine, miles uint8) {
   if miles <= e.milesLeft() {
      fmt.Println("You can make it there!")
   } else {
      fmt.Println("Need to fuel up first")
   }
}
```

In the above example, we have two types of 'engines' that power cars. We have a traditional gas powered vehicle and an EV. There is also a funciton called `canMakeIt()` which takes in the engine and the distance and returnes wether or not the car can make it to it's destionation. The issue is that it only takes in `gasEngines`. This is a problem as now we have no way of telling if an EV will make it. We could define a second function `canElectricVehicleMakeIt()`, but aside from being a mouthful, there is a way to make our funciton more flexible. This is where interfaces come into play.

In this very simple case we can define an interface that allows any type with a `milesLeft()` method into the funciton. This allows for the function to take in an EV or gas vehicle, so long as both types have the proper method defined.

```go
type engine interface {
   milesLeft() uint8 // Method name and return type = method signature
}

func canMakeIt(e engine, miles uint8) { // Go now checks to see if the value passed in for e has the method.
   if miles <= e.milesLeft() {
      fmt.Println("You can make it there!")
   } else {
      fmt.Println("Need to fuel up first")
   }
}
```

## Pointers

Welcome to the endless suffering of pointers. In `Go`, pointers are a type. They are values that store memory addresses rather than values. The syntax of pointers in `Go` is very similar to C languages.

A pointer initializes to `nil` by default in `Go`, and will either be 32 or 64 bytes long depending on the operating system. A key point to remember about a pointer is that it is stored the same way that any variable is stored in memory. This means that if at location `0x1b00` there is an `int32` with a value of `42`, then in address `0x1b01` there can be a `pointer` with a values of `0x1b00`, pointing to the int32. In other words, a pointer is a variable with a memory address as the value.

```go
var p *int32 // Creates a pointer to a 32 bit integer, but initializes to nil, lets pretend it is at 0x1b00. It currently points to nowhere. It will take up the locations from 0x1b00 to 0x1b07 because we are on a 64 bit OS.
```

We can also initialize the pointer to some value such as in the following example

```go
var p *int32 = new(int32) // Same pointer as before, but it initializes with value of 0x1b0c, pointing to a new 32 bit integer. The 32 bit integer will take up addresses 0x1b0c to 0x1b0f.

var i int32 // Creates a 32 bit integer with value 0, the address will be 0x1b08 and will span to 0x1b0b because it will take up 4 bytes.
```

To get the value of the pointer, we can dereference the pointer using the dereference operator `*` again.

```go
fmt.Printf("The value p points to is : %v", *p)
```

To change the value of the int at the pointers locaiton, we dereference the pointer, and then assing the value.

```go
*p = 42 // Set the value at 0x1b0c to 42
```

A common bug is having a nil pointer error. This means that we did not set the value of the pointer before attempting to set the value at the address.

```go
var p *int32
*p = 42 // Nil pointer error because p is pointing to nowhere
```

We can assign the value of a pointer using the reference operator `&`. As with earlier where we declared our pointer, and a value `i` at location `0x1b08`. We can now reference the address of `i` using this operator.

```go
var p *int32 = new(int32) // Points to a new integer

var i int32 // Points to a different new integer

p = &i // Now points to i instead of the new integer.

*p = 42 // Changes the value of i
```

This is where pointers diverge from regular variables. For instance, if there are two varialbes of `x` and `y`. If we write `x = y` the value from `y` is copied from `y` and then placed in the value at the location of `x`. This is how *most* things work in `Go`. For instance, functions in go are pass by value. This means that if we define a function and pass in a variable, the variable is not altered because the value is passed in, not the reference. If we instead wanted to alter the original variable we could pass in a reference to the variable so that the function would edit the original variable (though there are debates to be had about whether or not anyone should be doing that). Some other types in `Go` also funciton by reference rather than value. For instance, if we have a two slices `slice1` and `slice2`, and we set `slice1 = slice2`. Then when we edit `slice1` we are modifying `slice2` and vice versa. This is because a slice works by using pointers to an array under the hood, and so instead of creating a copy of the slice, we are actually just setting both `slice1` and `slice2` to point at the same array.

To summarize, the point of that paragraph, if there are unexpected errors occuring where data is changing that shouldn't be, it would be pertinent to consult either the documenation, the internet, or the sweaty guy in the neighbor's basement about whether or not the data we are using is being passed around by reference or by value.

### Pointers in Functions

So with all that said, why would we want to pass by reference instead of by value? There are a number of very valid reasons to do this, but first a discussion of reasons why it would be frowned upon. Passing many values around by reference can cause unpredictable behavior. A good way to think about it is that if everyone had a remote to one TV, and they all had different preferences, how long would the channel remain the same before changing. What are the odds that we could watch one episode of our favorite show all the way through without interuptions? Its the same for passing by reference. Allowing multiple things to access the same address can cause the value at that location to become difficult to predict, and if that value is imperative to the operation of multiple parts of our program it can cause cascading issues all the way down. In other words, explicitly defining, setting, and changing a variable generally produces cleaner, easier to maintain code. In regards to maintainability, pointers also can be a difficult concept to grasp, understand, and visualize, even with lots of experience. Pointers add a layer of abstraction to a value, where instead of `variable = value` there is now `variable = value that points to another place and that other place contains a value and the place that I am pointing to could change and then it would be both a different value and locaiton`. It can be a pain to debug code that points to other places.

So with that said, why would we want to use pointers? They are more efficient in some cases. Take for instance the following example.

```go
var arr = [5]int64{2, 4, 6, 8, 10}
var result [5]int64 = square(arr)

func square(arr [5]int64) [5]int64 {
   for i := range arr {
      arr[i] = arr[i] * arr[i]
   }
   return arr
}

fmt.Println(arr) // [2, 4, 6, 8, 10]
fmt.Println(result) // [4, 16, 36, 64, 100]
```

Because arrays are passed in by value, the function is recieving a copy of the array. This means that to run this function, first the array must be copied over to a new location, then that new array is operated on, and then the values are passed back and copied into result. This means that there is time spent copying the array as well as we have doubled our memory usage. What if we just need the squared values and nothing more?

```go
var arr = [5]int64{2, 4, 6, 8, 10}
var result [5]int64 = square(&arr)

func square(arr *[5]int64) [5]int64 {
   for i := range arr {
      arr[i] = arr[i] * arr[i]
   }
   return *arr
}

fmt.Println(arr) // [4, 16, 36, 64, 100]
fmt.Println(result) // [4, 16, 36, 64, 100]
```

In the modified version of the code we have passed in the address of the array, and we dereference the pointer in the function. This modifies the orignial array. When we pass it back, it is passed back, it is passing back the original array. `Result` and `arr` end up pointing to the same location in memory. This helps to avoid creating copies of large pieces of data.

### Garbage Collection and Go

This is a side tangent that feels most fitting to place here. In the following example (an example from earlier), it should be clear that we can create a pointer pointing to one value, and then point away from that value later on with no way of returning to the first address.

```go
var p *int32 = new(int32) // Points to a new integer

var i int32 // Points to a different new integer

p = &i // Now points to i 
```

What happens to that first integer? Due to modern technical advancements in programming languages, we have something called *Garbage Collection*. The program will notice that the first int no longer has any way to be reached and will clear up the space for us, but that is probably one of the most ambiguos descriptions of that process we could get. 

`Go` uses a *Tri Color Mark and Sweep* garbage collector. This is a process that runs in the `Go` runtime during execution concurrently with our program. This is important because it means that `Go` will not stop to complete the garbage collection process, rather it will take small slices of time during our program to free up memory. The *Tri Color Mark and Sweep* garbage collector (which is now just going to be called the GC)effectivly uses a tree algorithm to determine if objects are reachable or not. First, the GC looks for *root* objects. These are objects that definitely have a reference to them, and include things like globabl variables, variables on the stack, and registers. Next, the GC will mark everything as trash by setting its color to `white`. Each root object is set to `gray`, meaning that it is possible to reach that item but not all attached nodes have been checked yet. Each connected node is set to gray and added to the list of nodes to check. When a node has had all of it's attached node, the GC will set it's color to black and will move on. This is done until no more nodes are left in the list to check. When this is all said and done, items that have no way to reach them will be removed to clear up heap space.

When does this happen? GC is handled by the `Go` runtime. Generally though, this will happen when the heap size grows or shrinks, during large amounts of memory allocation, and whenever the GC feels like running based on internal metrics and methods.

## Goroutines

A Goroutine is a way to implement *Concurrency* in `Go`. This is *NOT* necessarily parallelism. A recap of the difference is that concurrency is desining our program to handle multiple tasks at once. This may mean flipping back and forth between the tasks, or working on another task while one tasks is waiting on something like an API call. This is in contrast to parallelism which uses multiple CPU cores to to multiple tasks at the same time. In other words, say there are two tasks each taking 10 cycles of a CPU core. With just concurrency and no parallelism, it will still take 20 cycles of the core to complete both tasks, they will just be done simultaneously. In parallelism, it will take 10 cycles because two cores will cycles through the tasks at the same time. *Goroutines can be parallel though*, so long as the system allows it. Go has an underlying schedule that will map `m` goroutines to `n` OS threads. As a general rule, concurrency is used to solve *I/O bottlenecked problems* such as waiting on a server to respond, asking users to input data. Parallelism is used for *CPU bound* or *Computationally Bottlenecked* problems such as processing visual data for a game or running calculations for neural nets (both problems which are generally offloaded to the GPU, which is really just a specialized miniature computer for handling embarassingly parallel computations).

In this example, there is a mock database. We will attempt to call the database once for each ID. This will take a substantial amount of time to consecutively call this database.

```go
import (
	"fmt"
	"math/rand"
	"time"
)

var dbData = []string{
	"id1", "id2", "id3", "id4", "id5",
	"id6", "id7", "id8", "id9", "id10",
	"id11", "id12", "id13", "id14", "id15",
	"id16", "id17", "id18", "id19", "id20",
}

func main() {
	t0 := time.Now()
	for i:=0; i<len(dbData); i++ {
		dbCall(i)
	}
	fmt.Printf("\nTotal Execution Time: %v\n", time.Since(t0))
}

func dbCall(i int) {
	var delay float32 = rand.Float32()*2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("\nThe result is: %v\n", dbData[i])
}
```

This will take approximately 20 seconds to run as the calls are being made consecutively. We can instead use the built in Goroutines to call these concurrently. This is done with the `go` keyword. This on its own is generally not enough though. This is because adding it in front of the call to the `dbCall()` function will spawn 20 goroutines, and then the for block will exit and the main will exit before the calls can complete. To fix this we can use wait groups from the `sync` package. A wait group is effectively just a counter that we can increment with `Add()` method, and it will decrement when we call the `Done` method. After the loop we call the `Wait` method which waits for the counter to count back down to zero before continuing. This is an example of how `Go` has very good tools for concurrency built in.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

var wg sync.WaitGroup
var dbData = []string{
	"id1", "id2", "id3", "id4", "id5",
	"id6", "id7", "id8", "id9", "id10",
	"id11", "id12", "id13", "id14", "id15",
	"id16", "id17", "id18", "id19", "id20",
}

func main() {
	t0 := time.Now()
	for i:=0; i<len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal Execution Time: %v\n", time.Since(t0))
}

func dbCall(i int) {
	var delay float32 = rand.Float32()*2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("\nThe result is: %v\n", dbData[i])
	wg.Done()
}
```

Now, what about storing data and avoiding race conditions? The sync package comes with the mutual exclusion, or the `Mutex`. This has the `Lock` and `Unlock` methods

```go
var m = sync.Mutex{}

func dbCall(i int) {
	var delay float32 = 2000 
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("\nThe result is: %v\n", dbData[i])
	m.Lock()
	results = append(results, dbData[i])
	m.Unlock()
	wg.Done()
}
```

There are even more options within the sync package to really specify how our program allows reading and writing to data. For instance there is the `RWMutex` which also provides the ability to lock and unlock reading separately from writing.

```go
var m = sync.RWMutex{}

func dbCall(i int)

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("\nThe result is: %v\n", dbData[i])
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}
 
 func log() {
	m.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock()
}
```

`RLock` will check if there is any full lock in place. If there isn't, it will aquire a read lock. Many goroutines can have read locks at once. A full lock cannot be put in place until all read locks are clears. Once all read locks are cleared, a full lock can be put in place and no other locks can be aquired until that lock is cleared. In other words, only one goroutine can write to the slice at a time, and nothing can read from the slice until the write is done, however many routines can read from the slice all at once.

## Channels

Channels are the next step up from locks. They are a built in, thread safe way to synchronize and share data, as well as coordinate and manage goroutines. Thye are a typed data pipline that does not require explicitly defining locks. It should, however, be noted that this does not mean that channels don't rely on these things themselves. The underlying implementation of a channel is an underlying *circular queue* with a *mutex* to avoid race conditions. When go routines interact with channels, `Go` automatically handles things like defining, aquiring, and releasing locks for us. In other words, channels are a way to abstract away a lot of the complications of concurrent programming.

```go
func main() {

	// Create a channel containing and integer with a size of 1. There is currently no value in c
	var c = make(chan int) 

	// Insert the value of 1 into the channel
	c <- 1  

	// Retrieves the value and sets i equal to the value in c
	var i = <- c
}
```

Here we can see how this is implemented with a queue. the `<-` operator is being used to insert a value and retrieve it later. This is in contrast to say setting a variable `y = 1` and then saying `x = y`. In that case, y still maintains its value. The channel, however, does not maintain value 1, as it has popped the value out of the circular queue.

The code as is will cause a deadlock error though. This is because when writing to an unbuffered channel, the code will block until something reads from the channel. Because of this, the code stops when inserting the value of 1 and never move forward to retrieve and print the value. Rather than hanging, the creators of `Go` were smart enough to see this possiblity, and a *deadlock* error will be thrown instead.

This example lacks the meat and bones of how channels are actually supposed to be used.

```go
func main() {
	var c = make(chan int)
	go process(c)
	fmt.Println(<-c)
}

func process(c chan int) {
	c <- 42
}
```

This example will actually work. Earlier it was stated that when writing to an unbuffered channel in `Go`, the code will block. This is because just as two people must be present for a hand shake to occur, channels must have a corresponding `send` and `recieve` flag to complete the exchange of data.

We can also use channels with loops as shown below:

```go
func main() {
	var c = make(chan int)
	go process(c)
	// For loop will continually return to c until notified that no more transfers will occur
	for v := range c {
		// Print the value from the channel
		fmt.Println(v)
	}
}

func process(c chan int) {
	// Add the values between 0 to 99 to the channel
	for i:=0; i < 100; i++{
		// Wait to pass off the data
		c <- i
	}
	// Notify anywhere referencing this channel that the transfers are over
	close(c)
}
```

The `close(c)` method is imperative to this function as without it the program will throw another deadlock error. This is because unless we notify the for loop in `main()`, we never tell the program that there are no more transfers to happen. This in turn results in sitting at the entrance to the loop waiting to hear back from someone who is never going to call.

We can also use the `defer` keyword in the function instead of putting the `close(c)` at the end of the funciton. The `defer` keyword deferrs a given action until after the function returns.

```go
func main() {
	var c = make(chan int)
	go process(c)
	for v := range c {
		fmt.Println(v)
	}
}

func process(c chan int) {
	defer close(c)
	for i:=0; i < 100; i++{
		c <- i
	}
}
```

Earlier there was mention of an unbuffered channel, meaning a channel with a single slot. How does a channel behave if it has multiple slots contained within it? It functions just as we might expect. It allows for multiple values to be written to the channel before blocking the code. This means that a process can work ahead while the recieving side churns away at it's process.

```go
func main() {
	var c = make(chan int, 10)
	go process(c)
	for v := range c {
		fmt.Println(i)
		time.Sleep(time.Second*1)
	}
}

func process(c chan int) {
	defer close(c)
	for i:=0; i < 100; i++{
		c <- i
	}
	fmt.Println("Exiting Process")
}
```

### Select Statements

Lets also look at a select statement. What is a select statment? It's a special switch case statement that only works with channel operations, designed for concurrent communiation. Each case in a select statement must contain a channel operation. It's behavior allows for a goroutine to handle multiple channels of communication. When multiple cases are ready at the same time it randomly picks one case to process. It also has a default statement that will run immediately if no channels are ready.

```go
func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
		case tofuWebsite := <-tofuChannel:
			fmt.Printf("\nEMAIL SENT: Found a deal on tofu at %s\n", tofuWebsite)
		case chickenWebsite := <-chickenChannel:
			fmt.Printf("\nTEXT SENT: Found a deal on chicken at %s\n", chickenWebsite)
	}
}
```

## Generics

Generics are a way to avoid duplicating logic for functions that do similar things. An example of this is a sum function which takes in a slice of floats and returns a float. That's great, but what if we also want to add integers? Now we need to add a function that takes in integers. Howabout unsigned integers? That's another function. This can lead to a mass amount of code duplication. Instead we can use generics. Generics allow for the definition of a function that takes in multiple types and uses that type in the function regardless of what it is. This can be done by the following:

```go
var intSlice = []int{1, 2, 3}
var intSum = sumSlice(intSlice)

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}
```

In this example, before defining the parameters in the `()` for the function, we declare our generic type `T` to be an `int`, `float32` or a `float64`. From this point onward, `T` can be mentally replaced with the type that is inserted into the function. In this case, `int` was inserted, so the slice was of type `int`, and the sum was declared as type `int`. Furthermore, we can get even more generic with the `any` type. This type allows for any type to be passed into the funciton. In our example though, this would throw an error. This is because not all types allow for the `+` operator and `Go` is smart enough to catch this. A place where an any type would make sense is a function that checks if a slice contains a certain number of elements. Regardless of what is inside the slice, any slice can be checked for the length. In other words, the function declaration would look like

```go
func checkLength[T any](slice []T) int {}
```

While there have been a number of example where generics are being used for functions, they can also be used for structs as well.

## Building an API

### Project structure

There isn't much to say about this other than the fact that `Go` has general guidlines for structuring projects. To get more information on the general layout of this, go to [this repo](https://github.com/golang-standards/project-layout) to find more information on the subject.

### Installing Packages

One way to install packages is to import them. After declaring imports you can run `go mod tidy` which will search through your imports and install the packages for you.

### Public or Private

Functions in `Go`, like many other languages, can be public or private. What deliniates the two from each other? Whether or not the function name is capitalized. In other words `sum()` is a private method that cannot be imported but `Sum()` is a public method that can be. It's not like that makes it difficult to find errors or anything though.

## Sidenotes

### Make VS New

What is the difference between the `make` keyword and the `new` keyword? They are used for different data types and return different values. Generally speaking, `new` is used to allocate and initialize memory for specified types. After new has allocated memory and zeroed out the value it returns a pointer to the value. For instance, `var x = new(int)` will go and zero out a space in memory and then return the address to x, where x now points to where new created the value. `Make` on the other hand is only used for specific types. These types are slices, maps, and channels. It creates the values and returns an initialized, non-zeroed, value. Another way to think about this is that `new` initializes data and `make` initializes a data structure.

## Sources

[Development Docs](https://go.dev/doc/)
[Learn GO Fast: Full Tutorial](https://www.wetube.com/watch?v=8uiZC0l4Ajw&t=490s)
[Go (Golang) programming language in 3 minutes](https://www.youtube.com/watch?v=8U8erH5qOZ8)

```

```
