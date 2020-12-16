package main

import (
	"fmt"
	"reflect"
)

/*
Problem 1

(*) Find the last element of a list.

Example in Haskell:

	λ> myLast [1,2,3,4]
	4
	λ> myLast ['x','y','z']
	'z'
*/
func myLast(arr interface{}) interface{} {
	arrValues := reflect.ValueOf(arr)
	length := arrValues.Len()

	return arrValues.Index(length - 1)
}

/*
Problem 2
(*) Find the last but one element of a list.

(Note that the Lisp transcription of this problem is incorrect.)

Example in Haskell:

	λ> myButLast [1,2,3,4]
	3
	λ> myButLast ['a'..'z']
	'y'
*/
func myButLast(arr interface{}) interface{} {
	arrValues := reflect.ValueOf(arr)
	length := arrValues.Len()

	return arrValues.Index(length - 2)
}

/*
Problem 3
(*) Find the K'th element of a list. The first element in the list is number 1.

Example in Lisp:

	* (element-at '(a b c d e) 3)
	c

Example in Haskell:

	λ> elementAt [1,2,3] 2
	2
	λ> elementAt "haskell" 5
	'e'
*/
func elementAt(arr interface{}, pos int) interface{} {
	arrValues := reflect.ValueOf(arr)

	return arrValues.Index(pos - 1).Interface()
}

/*
Problem 4
(*) Find the number of elements of a list.

Example in Haskell:

	λ> myLength [123, 456, 789]
	3
	λ> myLength "Hello, world!"
	13
*/
func myLength(arr interface{}) int {
	arrValues := reflect.ValueOf(arr)
	return arrValues.Len()
}

/*
Problem 5
(*) Reverse a list.

Example in Haskell:

	λ> myReverse "A man, a plan, a canal, panama!"
	"!amanap ,lanac a ,nalp a ,nam A"
	λ> myReverse [1,2,3,4]
	[4,3,2,1]
*/
func myReverse(arr interface{}) interface{} {
	arrValues := reflect.ValueOf(arr)
	length := arrValues.Len()

	switch arr.(type) {
	case string:
		r := []rune(arrValues.String())
		var res []rune
		for i := len(r) - 1; i >= 0; i-- {
			res = append(res, r[i])
		}
		return string(res)
	case []int:
		r := make([]int, length)
		var res []int
		for i := len(r) - 1; i >= 0; i-- {
			res = append(res, r[i])
		}
		return res
	default:
		res := make([]interface{}, length)
		for i := 0; i != length; i++ {
			res[i] = arrValues.Index(length - i - 1).Interface()
		}
		return res
	}
}

/*
Problem 6
(*) Find out whether a list is a palindrome. A palindrome can be read forward or backward; e.g. (x a m a x).

Example in Haskell:

	λ> isPalindrome [1,2,3]
	False
	λ> isPalindrome "madamimadam"
	True
	λ> isPalindrome [1,2,4,8,16,8,4,2,1]
	True
*/
func isPalindrome(arr interface{}) bool {
	var res bool
	arrValues := reflect.ValueOf(arr)
	length := arrValues.Len()

	switch arr.(type) {
	case []int:
		j := 0
		for i := length - 1; i > j; i-- {
			if arrValues.Index(i).Int() != arrValues.Index(j).Int() {
				return false
			}
			j++
		}
		res = true
	case string:
		res = arr.(string) == myReverse(arr)
	default:
		res = arr == myReverse(arr)
	}

	return res
}

/*
Problem 7
(**) Flatten a nested list structure.

Transform a list, possibly holding lists as elements into a `flat' list by replacing each list with its elements (recursively).

Example in Lisp:

	* (my-flatten '(a (b (c d) e)))
	(A B C D E)

Example in Haskell:

We have to define a new data type, because lists in Haskell are homogeneous.

 	data NestedList a = Elem a | List [NestedList a]
	λ> flatten (Elem 5)
	[5]
	λ> flatten (List [Elem 1, List [Elem 2, List [Elem 3, Elem 4], Elem 5]])
	[1,2,3,4,5]
	λ> flatten (List [])
	[]
*/

/*
Problem 8
(**) Eliminate consecutive duplicates of list elements.

If a list contains repeated elements they should be replaced with a single copy of the element. The order of the elements should not be changed.

Example in Lisp:

	* (compress '(a a a a b c c a a d e e e e))
	(A B C A D E)

Example in Haskell:

	λ> compress "aaaabccaadeeee"
	"abcade"
*/

/*
Problem 9
(**) Pack consecutive duplicates of list elements into sublists. If a list contains repeated elements they should be placed in separate sublists.

Example in Lisp:

	* (pack '(a a a a b c c a a d e e e e))
	((A A A A) (B) (C C) (A A) (D) (E E E E))

Example in Haskell:

	λ> pack ['a', 'a', 'a', 'a', 'b', 'c', 'c', 'a',
	             'a', 'd', 'e', 'e', 'e', 'e']
	["aaaa","b","cc","aa","d","eeee"]
*/

/*
Problem 10
(*) Run-length encoding of a list. Use the result of problem P09 to implement the so-called run-length encoding data compression method. Consecutive duplicates of elements are encoded as lists (N E) where N is the number of duplicates of the element E.

Example in Lisp:

	* (encode '(a a a a b c c a a d e e e e))
	((4 A) (1 B) (2 C) (2 A) (1 D)(4 E))

Example in Haskell:

	λ> encode "aaaabccaadeeee"
	[(4,'a'),(1,'b'),(2,'c'),(2,'a'),(1,'d'),(4,'e')]
*/

func main() {
	intList := []int{1, 2, 3, 4}
	stringList := []string{"x", "y", "z"}
	_string := "Hello, Golang World!"
	_jString := "こんにちは世界！"
	fmt.Println("Golang version for H-99: Lists")
	fmt.Println("\nQ1: myLast <list>")
	fmt.Printf("Input: %v\n", intList)
	fmt.Printf("Output: %v\n", myLast(intList))
	fmt.Printf("Input: %v\n", stringList)
	fmt.Printf("Output: %v\n", myLast(stringList))

	fmt.Println("\nQ2: myButLast <list>")
	fmt.Printf("Input: %v\n", intList)
	fmt.Printf("Output: %v\n", myButLast(intList))
	fmt.Printf("Input: %v\n", stringList)
	fmt.Printf("Output: %v\n", myButLast(stringList))

	fmt.Println("\nQ3: elementAt <list> <pos>")
	pos := 3
	fmt.Printf("Input: %v %d\n", intList, pos)
	fmt.Printf("Output: %v\n", elementAt(intList, pos))
	pos = 8
	fmt.Printf("Input: %q %d\n", _string, pos)
	fmt.Printf("Output: %q\n", elementAt(_string, pos))

	fmt.Println("\nQ4: myLength <list>")
	fmt.Printf("Input: %v\n", intList)
	fmt.Printf("Output: %v\n", myLength(intList))
	fmt.Printf("Input: %q\n", _string)
	fmt.Printf("Output: %d\n", myLength(_string))

	fmt.Println("\nQ5: myReverse <list>")
	fmt.Printf("Input: %v\n", intList)
	fmt.Printf("Output: %v\n", myReverse(intList))
	fmt.Printf("Input: %q\n", _string)
	fmt.Printf("Output: %q\n", myReverse(_string))
	fmt.Printf("Input: %q\n", _jString)
	fmt.Printf("Output: %q\n", myReverse(_jString))

	fmt.Println("\nQ6: isPalindrome <list>")
	paList := []int{1, 2, 4, 8, 16, 8, 4, 2, 1}
	paString := "madamimadam"
	fmt.Printf("Input: %v\n", intList)
	fmt.Printf("Output: %v\n", isPalindrome(intList))
	fmt.Printf("Input: %v\n", paList)
	fmt.Printf("Output: %v\n", isPalindrome(paList))
	fmt.Printf("Input: %q\n", _string)
	fmt.Printf("Output: %v\n", isPalindrome(_string))
	fmt.Printf("Input: %q\n", paString)
	fmt.Printf("Output: %v\n", isPalindrome(paString))

}

// TODO: Try to modify isPalindrome to reuse myReverse
