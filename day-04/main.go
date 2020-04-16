package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func processPasswordPotentials(min int, max int, isPartTwo bool) int {
	validPasswords := 0

	for i := min; i <= max; i++ {
		stringifiedPassword := strconv.Itoa(i)
		passwordSlice := strings.Split(stringifiedPassword, "")
		if checkProgressiveIncreases(passwordSlice) && checkMinimumDoubles(stringifiedPassword, passwordSlice, isPartTwo) {
			validPasswords++
		}
	}

	return validPasswords
}

func checkMinimumDoubles(password string, passwordSlice []string, isPartTwo bool) bool {
	matched, err := regexp.MatchString(`(00|11|22|33|44|55|66|77|88|99)`, password)
	check(err)

	if !matched {
		return false
	}

	if isPartTwo {
		isValid := false
		for n := 0; n < 5; n++ {
			if passwordSlice[n] == passwordSlice[n+1] {
				regNoOddPair := regexp.MustCompile("(([^" + passwordSlice[n] + "]|\\A)" + passwordSlice[n] + passwordSlice[n] + "([^" + passwordSlice[n] + "]|\\z))")
				count := len(regNoOddPair.FindAllString(password, -1))

				if count == 1 {
					isValid = true
				}
			}
		}

		return isValid
	}

	return true
}

func checkProgressiveIncreases(passwordSlice []string) bool {
	for i := 0; i < 5; i++ {
		current, currErr := strconv.Atoi(passwordSlice[i])
		check(currErr)

		next, nextErr := strconv.Atoi(passwordSlice[i+1])
		check(nextErr)

		if current > next {
			return false
		}
	}

	return true
}

func main() {
	min, max := 206938, 679128

	fmt.Println("Part 1", processPasswordPotentials(min, max, false))
	fmt.Println("Part 2", processPasswordPotentials(min, max, true))
}
