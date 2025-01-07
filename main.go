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
	"github.com/Sourjaya/dsa/fastAndSlowPointers"
	hashMapsAndSets "github.com/Sourjaya/dsa/hasMapsAndSets"
	linkedlist "github.com/Sourjaya/dsa/linkedList"
	"github.com/Sourjaya/dsa/twoPointers"
)

func main() {
	var choice, choice1, ch int
	for {
		fmt.Println("...........................................Main Menu...........................................")
		fmt.Println("1. Two Pointers")
		fmt.Println("2. Hash Maps and Hash Sets")
		fmt.Println("3. Linked List")
		fmt.Println("4. Fast and Slow Pointers")
		fmt.Println("5. Exit")
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
						continue
					}

					inputArray, err := toIntSlice(array.([]interface{}))
					if err != nil {
						log.Printf("Error in converting from interface slice to integer slice : %v", err)
						continue
					}

					fmt.Println("Output(PairSumSorted) : ", twoPointers.PairSumSorted(int(target.(float64)), inputArray))
				case 2:
					array, err := input("array")
					if err != nil {
						log.Printf("Error in input : %v", err)
						continue
					}

					inputArray, err := toIntSlice(array.([]interface{}))
					if err != nil {
						log.Printf("Error in converting from interface slice to integer slice : %v", err)
						continue
					}
					fmt.Println("Output(TripletSum) : ", twoPointers.TripletSum(inputArray))
				case 3:
					str, err := input("string")
					if err != nil {
						log.Printf("Error in input : %v", err)
						continue
					}
					fmt.Println("Output(IsPalindrome) : ", twoPointers.IsPalindrome(str.(string)))
				case 4:
					array, err := input("array")
					if err != nil {
						log.Printf("Error in input : %v", err)
						continue
					}

					inputArray, err := toIntSlice(array.([]interface{}))
					if err != nil {
						log.Printf("Error in converting from interface slice to integer slice : %v", err)
						continue
					}
					fmt.Println("Output(TripletSum) : ", twoPointers.LargestContainer(inputArray))
				case 5:
					array, err := input("array")
					if err != nil {
						log.Printf("Error in input : %v", err)
						continue
					}

					inputArray, err := toIntSlice(array.([]interface{}))
					if err != nil {
						log.Printf("Error in converting from interface slice to integer slice : %v", err)
						continue
					}
					fmt.Println("Output(Move Zeroes) : ", twoPointers.MoveZeroes(inputArray))
				case 6:
					str, err := input("string")
					if err != nil {
						log.Printf("Error in input : %v", err)
						continue
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
						continue
					}

					inputArray, err := toIntSlice(array.([]interface{}))
					if err != nil {
						log.Printf("Error in converting from interface slice to integer slice : %v", err)
						continue
					}

					fmt.Println("Output(PairSumSorted) : ", hashMapsAndSets.PairSumUnsorted(int(target.(float64)), inputArray))
				case 2:
					array, _ := input("array")
					inputArray, err := toInt2dSliceFromString(array.([]interface{}))
					if err != nil {
						log.Printf("Error in converting from interface slice to integer slice : %v", err)
						continue
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
						continue
					}
					hashMapsAndSets.ZeroStrippingInPlace(inputArray)
				case 4:
					array, _ := input("array")
					inputArray, err := toIntSlice(array.([]interface{}))
					if err != nil {
						log.Printf("Error in converting from interface slice to integer slice : %v", err)
						continue
					}
					fmt.Println("Longest Consecutive Sequence has length : ", hashMapsAndSets.LongestChain(inputArray))
				case 5:
					array, _ := input("array")
					r, _ := input("common ratio")
					inputArray, err := toIntSlice(array.([]interface{}))
					if err != nil {
						log.Printf("Error in converting from interface slice to integer slice : %v", err)
						continue
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
				fmt.Println("3. Linked List Intersection")
				fmt.Println("4. LRU Cache")
				fmt.Println("5. Palindrome List")
				fmt.Println("6. Flatten Multi-Level Linked List")
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
					listA := inputLinkedList()
					listB := inputLinkedList()

					if ptr := linkedlist.LinkedListIntersection(listA, listB); ptr != nil {
						fmt.Println("The node where the two lists intersect is : ", ptr.Value)
					} else {
						fmt.Println("No intersection")
					}
				case 4:
					var cache *linkedlist.LRUCache
					flag := false
					for {
						if flag {
							break
						}
						fmt.Println("...........................................LRU Cache...........................................")
						fmt.Println("1. Initialize")
						fmt.Println("2. Put")
						fmt.Println("3. Get")
						fmt.Println("4. Exit")
						_, err := fmt.Scanf("%d", &ch)
						// Check for input errors
						if err != nil {
							fmt.Println("Invalid input! Please enter a valid integer.")
							clearBuffer()
							continue
						}

						switch ch {
						case 1:
							capacity, _ := input("capacity")
							cache = linkedlist.Constructor(int(capacity.(float64)))

						case 2:
							if cache == nil {
								fmt.Println("Cache not initialized! Please initialize first.")
								continue
							}
							key, _ := input("key")
							value, _ := input("value")
							cache.Put(int(key.(float64)), int(value.(float64)))
						case 3:
							if cache == nil {
								fmt.Println("Cache not initialized! Please initialize first.")
								continue
							}
							key, _ := input("key")
							if value := cache.Get(int(key.(float64))); value != -1 {
								fmt.Println("Value is : ", value)
							} else {
								fmt.Println("Key not in Cache")
							}
						case 4:
							flag = true
							continue
						default:
							fmt.Println("Wrong Choice!! Enter your choice again")
						}
					}
				case 5:
					list := inputLinkedList()
					if linkedlist.IsPalindrome(list) {
						fmt.Println("Linked List is Palindromic")
					} else {
						fmt.Println("Linked List is not a Palindrome List")
					}
				case 6:
					head := CreateMultiLevelList[int]()
					fmt.Println("Original List : ")
					dStructures.Print(head, 1)
					head = linkedlist.FlattenList(head)
					fmt.Println("Flattened List : ")
					dStructures.Print(head, 1)
				case 7:
					flag = true
					continue
				default:
					fmt.Println("Wrong Choice!! Enter your choice again")
				}
			}
		case 4:
			flag := false
			for {
				if flag {
					break
				}
				fmt.Println("...........................................Fast and Slow Pointers...........................................")
				fmt.Println("1. Flyod's Cycle Detection")
				fmt.Println("2. Finding Midpoint")
				fmt.Println("3. Happy Number")
				fmt.Println("4. Go Back")
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
					if fastAndSlowPointers.CycleDetection(list) {
						fmt.Println("Cycle detected")
					} else {
						fmt.Println("No cycle detected")
					}
				case 2:
					list := inputLinkedList()
					fmt.Println("Midpoint of the linked list is : ", fastAndSlowPointers.MidpointList(list))
				case 3:
					num, err := input("number")
					if err != nil {
						log.Printf("Error in input : %v", err)
						continue
					}
					if fastAndSlowPointers.HappyNumber(int(num.(float64))) {
						fmt.Println("The number is Happy Number")
					} else {
						fmt.Println("The number is not a Happy Number")
					}
				case 4:
					flag = true
					continue
				default:
					fmt.Println("Wrong Choice!! Enter your choice again")
				}
			}
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Wrong Choice!! Enter your choice again")
		}
	}
}

// CreateList interactively builds a multilevel linked list using user input.
func CreateMultiLevelList[T comparable]() *dStructures.MultiLevelListNode[T] {
	reader := bufio.NewReader(os.Stdin)

	var head *dStructures.MultiLevelListNode[T]
	var tail *dStructures.MultiLevelListNode[T]

	for {
		// Get value for the new node
		var val T
		fmt.Print("Enter value for the node: ")
		_, err := fmt.Scanf("%v", &val)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid value.")
			clearBuffer()
			continue
		}

		newNode := dStructures.NewMultiLevelListNode(val)
		if head == nil {
			head = newNode
			tail = newNode
		} else {
			tail.Next = newNode
			tail = newNode
		}

		// Ask if the user wants to add a child node
		fmt.Print("Does this node have a child? (Yes/No): ")
		childResponse, _ := reader.ReadString('\n')
		childResponse = strings.TrimSpace(strings.ToLower(childResponse))

		if childResponse == "y" || childResponse == "Y" {
			fmt.Printf("Creating child list for node with value %v:\n", val)
			newNode.Child = CreateMultiLevelList[T]()
		}

		// Ask if the user wants to add another node at this level
		fmt.Print("Do you want to add another node at this level? (Yes/No): ")
		nextResponse, _ := reader.ReadString('\n')
		nextResponse = strings.TrimSpace(strings.ToLower(nextResponse))

		if nextResponse != "y" || childResponse == "Y" {
			break
		}
	}

	return head
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

func inputDoublyLinkedList() dStructures.DoublyLinkedList[int] {
	list := dStructures.DoublyLinkedList[int]{}
	var num, key int
	var s string
	for {
		fmt.Print("Enter node value : ")
		fmt.Scanf("%d", &num)
		fmt.Print("Enter node key : ")
		fmt.Scanf("%d", &key)
		list.Add(key, num)
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
