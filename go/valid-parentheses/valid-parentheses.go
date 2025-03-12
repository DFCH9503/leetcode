package main

import(
	"fmt"
)

func isValid(s string) bool {
    stack := []rune{} // Stack for opening brackets
    hash := map[rune]rune{')': '(', ']': '[', '}': '{'}
    
    for _, char := range s {
        if match, found := hash[char]; found {
            // Check if stack is non-empty and matches
            if len(stack) > 0 && stack[len(stack)-1] == match {
                stack = stack[:len(stack)-1] // Pop
            } else {
                return false // Invalid
            }
        } else {
            stack = append(stack, char) // Push opening bracket
        }
    }
    return len(stack) == 0 // Valid if stack is empty
}

func main(){
	input := "()"
	fmt.Println(isValid(input))
}