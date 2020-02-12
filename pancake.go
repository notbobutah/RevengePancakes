package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

const blankside uint8 = '-'
const happyside uint8 = '+'

func main() {
	title()

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[1]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

// flip stack based on calculated best move
// 0th character in string array represents top of stack
// last character in string array represents bottom of stack
func flip(stack string) int {

    var flipstck int = cntFlips(stack);
	printStack(stack);
	return flipstck
}

func cntFlips(stack string) int  {
	if len(stack) == 1 && stack[0] == blankside {return 1}
	if len(stack) == 1 && stack[0] == happyside {return 0}

	// define regex to split string in splices
	m := regexp.MustCompile("-{1,}")
	p := regexp.MustCompile("\\+{1,}")

	// execute the regex and split the strings into splices
	var mmin = m.Split(stack, 100);
	var pmin = p.Split(stack, 100);

	// could not find an elegant way to return the splices as one
	// combining splices into single array
	var stackSlice []string

	// The for loop iterates over both of the arrays produced by the splice calls
	// Builds a new splce with each array entry in order, the regex.Splice command
	// leaves empty strings in first element of the array.
	// If the string starts with a - then the plus array will have
	// a blank string as the first in the array and should be processed second
		for i := 0; i < len(mmin); i++ {
			if len(mmin) > i && mmin[i] != "" {
				stackSlice = append(stackSlice, mmin[i])
			}
			if len(pmin) > i && pmin[i] != "" {
			stackSlice = append(stackSlice, pmin[i])
			}
		}
	// the total number of entries in the stackslice are the individual slices containing the target characters
	// number of flips required is the slice array lenght -1
	var flipcnt = len(stackSlice) - 1;

	// found that stacks ending in blank pancakes require an adjustment to their count.
	// most likely caused by the order of append in the for loop above.
	if stack[len(stack)-1] == '-' {
		flipcnt++;
	}
   return flipcnt
}

func printStack(stack string){
	m := regexp.MustCompile("-{1,}")
	p := regexp.MustCompile("\\+{1,}")

	var mmin = m.Split(stack, 100);
	var pmin = p.Split(stack, 100);

	var stackSlice []string

	for i := 0; i < len(mmin); i++ {
		if len(mmin) > i && mmin[i] != "" {
			stackSlice = append(stackSlice, mmin[i])
		}
		if len(pmin) > i && pmin[i] != "" {
			stackSlice = append(stackSlice, pmin[i])
		}
	}
    var flip int = 1;
    var idx int = 0;
	var matrix []string
    var row strings.Builder
	var workstr string
	var cellstr string
    var newflip = "+"
    var oldflip = "-"

    print("[")
	for i, cell := range stackSlice {
		row.WriteString(cell)
		print(cell)
		flip = i
	}
	println( "]", )

	for idx <= flip {
		workstr = "";
		if newflip == "+" {
			oldflip = "+"
			newflip = "-"
		} else {
			oldflip = "-"
			newflip = "+"
		}

		for i, cell := range stackSlice {
			if i < idx {
				cellstr = cell
				cellstr = strings.ReplaceAll(cellstr, oldflip, newflip);
				workstr = workstr + cellstr;
			} else {
				workstr = workstr + cell;
			}
		}
		matrix = append(matrix, workstr);
		println(" "+ matrix[idx]+ " : ", idx)
		idx++;
	}

}
