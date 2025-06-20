# GoLang - An introduction by video tutorials

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

   - This means that variables cannot be assigned values that are not their type, i.e. an integer cannot be assigned a string. This also means that when attempting to perform operations on variables, certain variables cannot interact. This is in direct contrast to a language like `JavaScript` where you can add a string to a float and it will continue along as if nothing has happened. This allows the compiler to do checks and error detection before compiling your code, as well as allows your editor to have better autocomplete because it is clear what variables are.

3. `Go` is Compiled

   - This means that your program will be compiled down to a standalone binary. This allows a program to be run in multiple enviornments. This is in contrast to Python which interprets the code at run time. This also means that programs will run much faster as there is no interpreter interpretting the code as it runs, rather the code is already ready to be run by the machine.

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

When starting a `Go` project, you are just creating a module.

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

If the package name is not called main, this is not a requirement for your package to have.

To import a package you can use the `import` keyword followed by the package name in quotes.

```go
import "fmt"
```

This will now cause issue because `Go` _does not let you import a module without using it_. To complete our program we can now use the print function in `Go` to print out **Hello World** to the console.

```go
fmt.Println("Hello, World!")
```

### Building and running our file

To build and run this file you use the following:

```bash
go build {FILE PATH}
```

This produces a file titled `main` which is the binary for the program. You can run it in the console to see what it does. Alternatively you can use:

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

`int` will default to a specific value based on your system. There are ways to specify the the size and type of int which can be seen in this table:

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

It should be noted as well, `int` initialize to _0_ by default which in contrast to `JavaScript`'s `undefined` or `C++` having a garbled mess of whatever was there before you initialized the value.

The max and mins are also something to look out for. If you try to initialize a value over its max size, the compiller will throw an error. If you update a value however, the compiller is no longer involved, and there will be no error. This can lead to unexpected behavior.

```go
var numberOne int8 = 127 // No error

var numberTwo int8 = 128 // Overflow error

numberOne = numberOne + 1 // No error
```

Integer division results in an integer, which rounds any integer down. If you want the remainder you can get it using the modulo operator `%`.

### Floats

`Go` has two types of floats. You must specify the float type when declaring a float.

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

Now that we have two types of numbers, we can address type casting. You cannot directly add an `int` and a `float` together because they are different types. Because of this you need to type cast from an int to a float or vice versa to avoid and error.

```go
var integer int32 = 2
var float float32 = 2.0

var fail float32 = float + integer // Not quick maths

var success float32 = float + float32(integer) // 2.0 + 2.0 is 4.0
var result int32 = int32(success) - 1 // -1 that's 3, quick maths.
```

### String

The `string` keyword is used for strings. Depending on what you use to define the string, the code will allow different behaviors. For instance, the use of `" "` means that the string is single line terminating.

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

To get the length of the string you can try using the built in `len()` function. It should however be noted that this is not always the number of charcters in the string. This is because `len()` returns the number of bytes used. Becuase `Go` encodes in `UTF-8`, characters outside of this encoding will use more bytes, thereby shifting the result of the `len()` funciton.

```go
var stringOne string = "A"
var stringTwo string = "γ"

fmt.Println(stringOne) // 1
fmt.Println(stringTwo) // 2
```

To get around this if you expect strings with characters outside of `UTF-8`, import the `"unicode/utf8"` and use the following:

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

These are variables that don't vary. In other words, you can assign them, but once assigned, they don't change. You also have to assign a value to a const at declaration unlike variables where there is a default value.

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

For passing parameters, the name of the parameter comes first, followed by the type declaration. This also means that you cannot pass other data types into the function. For instance, if you wanted to print an integer and you have a function that takes in a string, you will first have to cast the value to a string before passing it in.

```go
func functionName(value string) {
   fmt.Println(value)
}
```

These function types have had void return type so far, but how about function that return things? You must also declare what value the function will be returning so that `Go`'s compiler can check for errors. You can also add multipler return values as shown below.

```go
func intDivision(a int, b int) (int, int) {
	var result int = a / b
	var remainder int = a % b
	return result, remainder
}
```

Now we can ask, what happens when we hit a division by zero error? `Go` throws a panic error and exits out of the program. This is not great as far as handling errors goes so what is standard practice is to add a return type `error` along with your function if your function can encounter errros. When an `error` is first defined it will default to a value of `nil`. This also introduces the syntax for if statements as a check will occur to see if the divisor is 0.

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

You can also immediately initialize an array using the following syntax:

```go
var intArray [3]int32 = [3]int32{1, 2, 3}

// And of course there is shorthand such as the following
var intArray = [3]int32{1, 2, 3}
intArray := [3]int32{1, 2, 3}
intArray := [...]int32{1, 2, 3} // The ... makes the compiler infer the size of the array.

```

The `...` is called the variadic parameter operator, and in `Go` is used to denote that a function or template can accept a varialbe number of arguments.

### Slices

Slices are a wrapper for arrays. In other words, a slice is a superset of the features of an array. The way that slice is defined is similar to an array, though it is not the same. Instead of adding a parameter for the size of the array, you leave the value empty. This creates a slice instead of an array. This opens up a variety of new methods, such as the append feature.

```go
var intSlice []int32 = []int32{4 5 6}
fmt.Printf("The length is %v with capacity of %v", len(intSlice), cap(intSlice))
// Cap is a built in function allowing you to see what the capacity of a slice is.

intSlice = append(intSlice, 7)
fmt.Printf("The length is %v with capacity of %v", len(intSlice), cap(intSlice))
```

The underlying workings of a slice are as follows:

1. A new array with the exact length is allocated in memory.
2. When running the append, the method checks if there is enough room for a new value.
3. If there is not enough room, a new array with increased capacity will be allocated (this could be larger than needed, such as doulbing capacity to add a single value).
4. The appended value is copied in, followed by copying in the values from the previous array.

In other words, a new array in a new location in memory is returned after these operations. You _should not assume that the location remains the same while using a slice._

Interestingly you cannot index into the unassigned slots in the slice. So if there are six values in a slice with a capacity of ten, you will get an `index out of range` error.

You can also add multiple values to the slice at once depicted below.

```go
throwAway := []int32{8,9}
intSlice = append(intSlice, throwAway...) // Used after a value, this will spread out the values contained in the array/slice.
```

You can also use the `make` syntax to create a slice. This can be good if you know that there is a maximum size that your slice will take up and you want to avoid reallocating for resizing.

```go
var intSlices3 []int32 = make(int32[], 3, 8)
```

### Maps

As with other languages, this is a table of `key:value` pairs. Here is the syntax for defining one. You can initialize with or without values. Referencing values is fairly standard.

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

You can iterate over maps, arrays, and slices using a `range`. This will return an iterator allowing you to iterate through each value (similar to python).

```go
for name:= range myMap {
	fmt.Printf("Name: %s, Age: %d\n", name, myMap[name])
}
```

For an array you can do the following.

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

## Sources

[Development Docs](https://go.dev/doc/)
[Learn GO Fast: Full Tutorial](https://www.youtube.com/watch?v=8uiZC0l4Ajw&t=490s)
[Go (Golang) programming language in 3 minutes](https://www.youtube.com/watch?v=8U8erH5qOZ8)

```

```
