package main

import "fmt"

func add(a string, b string) string {

	//retString := "0"
	aLen := len(a)
	bLen := len(b)
	newB := b
	newA := a

	if aLen > bLen {
		// pad b
		for i:=0 ; i < (aLen - bLen); i++ {
			newB = "0" + newB
		}
	}

	if aLen < bLen {
		// pad a
		for i:=0 ; i < ( bLen - aLen); i++ {
			newA = "0" + newA
		}
	}

	// continue since they are equal

	// perform addition one by one
	var carry int = 0
	var retval string = ""
	for i:= len(newA) - 1  ; i >= 0  ; i--{

		// 1 + 0
		if newA[i] == '0' {

			if newB[i] == '0' {
				// a 0 b 0 c 1
				if carry == 1 {
					retval = "1" + retval
					carry = 0
				}else {
					// a 0 b 0 c 0
					retval = "0" + retval
				}
			}else{ // b 1
				// a 0 b 1 c 1
				if carry == 1 {
					retval = "0" + retval
					carry = 1
				}else { // a 0 b 1 c 0
					retval = "1" + retval
				}

			}
		}else{ // a == 1
			if newB[i] == '0' {
				// a 1 b 0 c 1
				if carry == 1 {
					retval = "1" + retval
					carry = 0
				}else { // a 1 b 0 c 0
					retval = "0" + retval
				}
			}else{ // b == 1
				if carry == 1 { // a 1 b 1 c 1
					retval = "1" + retval
					carry = 1
				}else { // a 1 b 1 c 0
					retval = "0" + retval
					carry = 1
				}
			}
		}
	}

	// prepend the final carry
	if carry == 1 {
		retval = "1" + retval
	}
	return retval
}



func main() {
	fmt.Println(add("111", "1111"))
}
