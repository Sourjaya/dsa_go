package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Sourjaya/dsa/dStructures"
	hashMapsAndSets "github.com/Sourjaya/dsa/hasMapsAndSets"
	linkedlist "github.com/Sourjaya/dsa/linkedList"
	"github.com/Sourjaya/dsa/twoPointers"
)

func main() {
	var choice, choice1 int
	for {
		fmt.Println("...........................................Main Menu...........................................")
		fmt.Println("1. Two Pointers")
		fmt.Println("2. Hash Maps and Hash Sets")
		fmt.Println("3. Linked List")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice : ")

		_, err := fmt.Scanf("%d", &choice)
		// Check for input errors
		if err != nil {
			fmt.Println("Invalid input! Please enter a valid integer.")
			clearBuffer()
			continue
		}
		switch choice {
		case 1:
			flag := false
			for {
				if flag {
					break
				}
				fmt.Println("...........................................Two Pointers...........................................")
				fmt.Println("1. Pair Sum - Sorted")
				fmt.Println("2. Triplet Sum")
				fmt.Println("3. Is Palindrome")
				fmt.Println("4. Largest Container")
				fmt.Println("5. Move Zeroes")
				fmt.Println("6. Find Next lexicographical permutation")
				fmt.Println("7. Go Back")
				_, err := fmt.Scanf("%d", &choice1)
				// Check for input errors
				if err != nil {
					fmt.Println("Invalid input! Please enter a valid integer.")
					clearBuffer()
					continue
				}
				switch choice1 {
				case 1:
					array, err1 := input("array")
					target, err2 := input("target")
					if err1 != nil || err2 != nil {
						log.Printf("Error in input : %v %v", err1, err2)
					}

					inputArray, err := toIntSlice(array.([]interface{}))
					if err != nil {
						log.Printf("Error in converting from interface slice to integer slice : %v", err)
					}

					fmt.Println("Output(PairSumSorted) : ", twoPointers.PairSumSorted(int(target.(float64)), inputArray))
				case 2:
					array, err := input("array")
					if err != nil {
						log.Printf("Error in input : %v", err)
					}

					inputArray, err := toIntSlice(array.([]interface{}))
					if err != nil {
						log.Printf("Error in converting from interface slice to integer slice : %v", err)
					}
					fmt.Println("Output(TripletSum) : ", twoPointers.TripletSum(inputArray))
				case 3:
					str, err := input("string")
					if err != nil {
						log.Printf("Error in input : %v", err)
					}
					fmt.Println("Output(IsPalindrome) : ", twoPointers.IsPalindrome(str.(string)))
				case 4:
					array, err := input("array")
					if err != nil {
						log.Printf("Error in input : %v", err)
					}

					inputArray, err := toIntSlice(array.([]interface{}))
					if err != nil {
						log.Printf("Error in converting from interface slice to integer slice : %v", err)
					}
					fmt.Println("Output(TripletSum) : ", twoPointers.LargestContainer(inputArray))
				case 5:
					array, err := input("array")
					if err != nil {
						log.Printf("Error in input : %v", err)
					}

					inputArray, err := toIntSlice(array.([]interface{}))
					if err != nil {
						log.Printf("Error in converting from interface slice to integer slice : %v", err)
					}
					fmt.Println("Output(Move Zeroes) : ", twoPointers.MoveZeroes(inputArray))
				case 6:
					str, err := input("string")
					if err != nil {
						log.Printf("Error in input : %v", err)
					}
					input := str.(string)

					fmt.Println("Output(NextPermutation) : ", twoPointers.NextPermutation([]rune(input)))
				case 7:
					flag = true
					continue
				default:
					fmt.Println("Wrong Choice!! Enter your choice again")
				}
			}
		case 2:
			flag := false
			for {
				if flag {
					break
				}
				fmt.Println("...........................................Two Pointers...........................................")
				fmt.Println("1. Pair Sum - Unsorted")
				fmt.Println("2. Verify Sudoku Board")
				fmt.Println("3. Zero Stripping")
				fmt.Println("4. Longest Chain of Consecutive numbers")
				fmt.Println("5. Geometric Sequence Triplets")
				fmt.Println("6. Go Back")
				_, err := fmt.Scanf("%d", &choice1)
				// Check for input errors
				if err != nil {
					fmt.Println("Invalid input! Please enter a valid integer.")
					clearBuffer()
					continue
				}
				switch choice1 {
				case 1:
					array, err1 := input("array")
					target, err2 := input("target")
					if err1 != nil || err2 != nil {
						log.Printf("Error in input : %v %v", err1, err2)
					}

					inputArray, err := toIntSlice(array.([]interface{}))
					if err != nil {
						log.Printf("Error in converting from interface slice to integer slice : %v", err)
					}

					fmt.Println("Output(PairSumSorted) : ", hashMapsAndSets.PairSumUnsorted(int(target.(float64)), inputArray))
				case 2:
					array, _ := input("array")
					inputArray, err := toInt2dSliceFromString(array.([]interface{}))
					if err != nil {
						log.Printf("Error in converting from interface slice to integer slice : %v", err)
					}
					if hashMapsAndSets.ValidSudoku(inputArray) {
						fmt.Println("Valid Board")
					} else {
						fmt.Println("Not a Valid Board")
					}
				case 3:
					array, _ := input("array")
					inputArray, err := toInt2dSlice(array.([]interface{}))
					if err != nil {
						log.Printf("Error in converting from interface slice to integer slice : %v", err)
					}
					hashMapsAndSets.ZeroStrippingInPlace(inputArray)
				case 4:
					array, _ := input("array")
					inputArray, err := toIntSlice(array.([]interface{}))
					if err != nil {
						log.Printf("Error in converting from interface slice to integer slice : %v", err)
					}
					fmt.Println("Longest Consecutive Sequence has length : ", hashMapsAndSets.LongestChain(inputArray))
				case 5:
					array, _ := input("array")
					r, _ := input("common ratio")
					inputArray, err := toIntSlice(array.([]interface{}))
					if err != nil {
						log.Printf("Error in converting from interface slice to integer slice : %v", err)
					}
					fmt.Println("No. of Geometric Triplets are : ", hashMapsAndSets.GeometricTriplets(int(r.(float64)), inputArray))
				case 6:
					flag = true
					continue
				default:
					fmt.Println("Wrong Choice!! Enter your choice again")
				}
			}
		case 3:
			flag := false
			for {
				if flag {
					break
				}
				fmt.Println("...........................................Two Pointers...........................................")
				fmt.Println("1. Linked List Reversal")
				fmt.Println("2. Remove Kth Last Node")
				fmt.Println("3. Zero Stripping")
				fmt.Println("4. Longest Chain of Consecutive numbers")
				fmt.Println("5. Geometric Sequence Triplets")
				fmt.Println("6. Go Back")
				_, err := fmt.Scanf("%d", &choice1)
				// Check for input errors
				if err != nil {
					fmt.Println("Invalid input! Please enter a valid integer.")
					clearBuffer()
					continue
				}
				switch choice1 {
				case 1:
					list := inputLinkedList()
					reversedList, _ := linkedlist.LinkedListReversalRecursive(list, list.Head)
					fmt.Print("Reversed Linked list is : ")
					reversedList.Print()
				case 2:
					list := inputLinkedList()
					k, _ := input("k")
					updatedList := linkedlist.RemoveKthLastNode(&list, int(k.(float64)))
					fmt.Print("Linked List after removing kth last node : ")
					updatedList.Print()
				case 3:
					array, _ := input("array")
					inputArray, err := toInt2dSlice(array.([]interface{}))
					if err != nil {
						log.Printf("Error in converting from interface slice to integer slice : %v", err)
					}
					hashMapsAndSets.ZeroStrippingInPlace(inputArray)
				case 4:
					array, _ := input("array")
					inputArray, err := toIntSlice(array.([]interface{}))
					if err != nil {
						log.Printf("Error in converting from interface slice to integer slice : %v", err)
					}
					fmt.Println("Longest Consecutive Sequence has length : ", hashMapsAndSets.LongestChain(inputArray))
				case 5:
					array, _ := input("array")
					r, _ := input("common ratio")
					inputArray, err := toIntSlice(array.([]interface{}))
					if err != nil {
						log.Printf("Error in converting from interface slice to integer slice : %v", err)
					}
					fmt.Println("No. of Geometric Triplets are : ", hashMapsAndSets.GeometricTriplets(int(r.(float64)), inputArray))
				case 6:
					flag = true
					continue
				default:
					fmt.Println("Wrong Choice!! Enter your choice again")
				}
			}
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Wrong Choice!! Enter your choice again")
		}
	}
}

func toIntSlice(input []interface{}) ([]int, error) {
	var result []int
	for _, v := range input {
		// Type assertion to ensure the value is float64
		num, ok := v.(float64)
		if !ok {
			return nil, fmt.Errorf("value %v is not an int", v)
		}
		result = append(result, int(num))
	}
	return result, nil
}

func toIntSliceFromString(input []interface{}) ([]int, error) {
	var result []int
	for _, v := range input {
		// Type assertion to ensure the value is float64
		num, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("value %v is not an int", v)
		}
		n, _ := strconv.Atoi(num)
		result = append(result, n)
	}
	return result, nil
}

func toInt2dSlice(input []interface{}) ([][]int, error) {
	i := 0
	output := make([][]int, len(input))
	for _, v := range input {
		data, ok := v.([]interface{})
		if !ok {
			fmt.Println("Not a 2d array : ")
		}
		result, _ := toIntSlice(data)
		output[i] = append(output[i], result...)
		i++
	}
	return output, nil
}

func toInt2dSliceFromString(input []interface{}) ([][]int, error) {
	i := 0
	output := make([][]int, len(input))
	for _, v := range input {
		data, ok := v.([]interface{})
		if !ok {
			fmt.Println("Not a 2d array : ")
		}
		result, _ := toIntSliceFromString(data)
		output[i] = append(output[i], result...)
		i++
	}
	return output, nil
}

func clearBuffer() {
	var garbage string
	// Read and discard the remaining input in the buffer
	fmt.Scanln(&garbage)
}

func inputLinkedList() dStructures.LinkedList[int] {
	list := dStructures.LinkedList[int]{}
	var num int
	var s string
	for {
		fmt.Print("Enter node : ")
		fmt.Scanf("%d", &num)
		list.Add(num)
		fmt.Print("Do you want to continue(Y/N)")
		fmt.Scanf("%s", &s)
		if s != "Y" && s != "y" {
			break
		}
	}
	fmt.Print("Original Linked List : ")
	list.Print()
	return list
}

func input(msg string) (interface{}, error) {
	var result interface{}

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter your input %s : ", msg)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	// Trim any surrounding whitespace or newline characters
	userInput = strings.TrimSpace(userInput)

	// Attempt to unmarshal the input string into a generic interface
	if msg == "string" {
		return userInput, nil
	}

	err = json.Unmarshal([]byte(userInput), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
