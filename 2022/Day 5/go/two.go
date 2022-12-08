package main

import "fmt"

func part2() {
	var (
		stack1 = []string{"H", "C", "R"}
		stack2 = []string{"B", "J", "H", "L", "S", "F"}
		stack3 = []string{"R", "M", "D", "H", "J", "T", "Q"}
		stack4 = []string{"S", "G", "R", "H", "Z", "B", "G"}
		stack5 = []string{"R", "P", "F", "Z", "T", "D", "C", "B"}
		stack6 = []string{"T", "H", "C", "G"}
		stack7 = []string{"S", "N", "V", "Z", "B", "P", "W", "L"}
		stack8 = []string{"R", "J", "Q", "G", "C"}
		stack9 = []string{"L", "D", "T", "R", "H", "P", "F", "S"}
	)

	scanner := readFile()

	for scanner.Scan() {
		line := scanner.Text()
		data := getData(line)

		buf := make([]string, data.Count)

		switch data.From {
		case 1:
			stack1, buf = liftCratesPart2(stack1, buf, data.Count)
			break
		case 2:
			stack2, buf = liftCratesPart2(stack2, buf, data.Count)
			break
		case 3:
			stack3, buf = liftCratesPart2(stack3, buf, data.Count)
			break
		case 4:
			stack4, buf = liftCratesPart2(stack4, buf, data.Count)
			break
		case 5:
			stack5, buf = liftCratesPart2(stack5, buf, data.Count)
			break
		case 6:
			stack6, buf = liftCratesPart2(stack6, buf, data.Count)
			break
		case 7:
			stack7, buf = liftCratesPart2(stack7, buf, data.Count)
			break
		case 8:
			stack8, buf = liftCratesPart2(stack8, buf, data.Count)
			break
		case 9:
			stack9, buf = liftCratesPart2(stack9, buf, data.Count)
			break
		default:
			panic("Invalid From")
		}

		switch data.To {
		case 1:
			stack1 = append(stack1, buf...)
			break
		case 2:
			stack2 = append(stack2, buf...)
			break
		case 3:
			stack3 = append(stack3, buf...)
			break
		case 4:
			stack4 = append(stack4, buf...)
			break
		case 5:
			stack5 = append(stack5, buf...)
			break
		case 6:
			stack6 = append(stack6, buf...)
			break
		case 7:
			stack7 = append(stack7, buf...)
			break
		case 8:
			stack8 = append(stack8, buf...)
			break
		case 9:
			stack9 = append(stack9, buf...)
			break
		default:
			panic("Invalid To")
		}
	}

	fmt.Println("Stack 1:", stack1)
	fmt.Println("Stack 2:", stack2)
	fmt.Println("Stack 3:", stack3)
	fmt.Println("Stack 4:", stack4)
	fmt.Println("Stack 5:", stack5)
	fmt.Println("Stack 6:", stack6)
	fmt.Println("Stack 7:", stack7)
	fmt.Println("Stack 8:", stack8)
	fmt.Println("Stack 9:", stack9)
	fmt.Println()

	top := ""
	top += string(stack1[len(stack1)-1])
	top += string(stack2[len(stack2)-1])
	top += string(stack3[len(stack3)-1])
	top += string(stack4[len(stack4)-1])
	top += string(stack5[len(stack5)-1])
	top += string(stack6[len(stack6)-1])
	top += string(stack7[len(stack7)-1])
	top += string(stack8[len(stack8)-1])
	top += string(stack9[len(stack9)-1])

	fmt.Println(top)
}
