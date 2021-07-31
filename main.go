package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

// buildBalloon takes a slice of strings of max width, prepends/appends margins
// on first and last line, and returns a string with the contents of the balloon
func buildBalloon(lines []string, maxwidth int) string {
	var borders []string
	count := len(lines)
	var ret []string

	borders = []string{"/", "\\", "\\", "/", "|", "<", ">"}

	top := " " + strings.Repeat("_", maxwidth + 2)
	bottom := " " + strings.Repeat("_", maxwidth + 2)

	ret = append(ret, top)
	if count == 1 {
		s := fmt.Sprintf("%s %s %s", borders[5], lines[0], borders[6])
		ret = append(ret, s)
	} else {
		s := fmt.Sprintf("%s %s %s", borders[0], lines[0], borders[1])
		ret = append(ret, s)
		i := 1
		for ; i < count - 1; i++ {
			s = fmt.Sprintf("%s %s %s", borders[4], lines[i], borders[4])
			ret = append(ret, s)
		}
		s = fmt.Sprintf("%s %s %s", borders[2], lines[i], borders[3])
		ret = append(ret, s)
	}

	ret = append(ret, bottom)
	return strings.Join(ret, "\n")
}

// tabsToSpaces converts all tabs found in the strings found in the `lines` slice
// to 4 spaces, to prevent misalignments in counting the runes
func tabsToSpaces(lines []string) []string {
	var ret []string
	for _, l := range lines {
		l = strings.Replace(l, "\t", "    ", -1)
		ret = append(ret, l)
	}
	return ret
}

// calculateMaxWidth given a slice of strings returns the length of the string
// with the max width
func calculateMaxWidth(lines []string) int {
	w := 0
	for _, l := range lines {
		inString := utf8.RuneCountInString(l)
		if inString > w {
			w = inString
		}
	}

	return w
}

// normalizeStringLength takes a slice of strings and appends to each one a
// number of spaces needed to have them all the same number of runes
func normalizeStringLength(lines []string, maxwidth int) []string {
	var ret []string
	for _, l := range lines {
		s := l + strings.Repeat(" ", maxwidth - utf8.RuneCountInString(l))
		ret = append(ret, s)
	}
	return ret
}

// printFigure given a figure name prints it.
func printFigure(name string) {
	var cow = `         \  ^__^
          \ (oo)\_______
	    (__)\       )\/\
	        ||----w |
	        ||     ||
		`

	var stegosaurus = `         \                      .       .
          \                    / ` + "`" + `.   .' "
           \           .---.  <    > <    >  .---.
            \          |    \  \ - ~ ~ - /  /    |
          _____           ..-~             ~-..-~
         |     |   \~~~\\.'                    ` + "`" + `./~~~/
        ---------   \__/                         \__/
       .'  O    \     /               /       \  "
      (_____,    ` + "`" + `._.'               |         }  \/~~~/
       ` + "`" + `----.          /       }     |        /    \__/
             ` + "`" + `-.      |       /      |       /      ` + "`" + `. ,~~|
                 ~-.__|      /_ - ~ ^|      /- _      ` + "`" + `..-'
                      |     /        |     /     ~-.     ` + "`" + `-. _  _  _
                      |_____|        |_____|         ~ - . _ _ _ _ _>

	`

	var alligator = `      \
       \
        \
           .-._   _ _ _ _ _ _ _ _
.-''-.__.-'00  '-' ' ' ' ' ' ' ' '-.
'.___ '    .   .--_'-' '-' '-' _'-' '._
 V: V 'vv-'   '_   '.       .'  _..' '.'.
   '=.____.=_.--'   :_.__.__:_   '.   : :
           (((____.-'        '-.  /   : :
                             (((-'\ .' /
                           _____..'  .'
                          '-._____.-'
`

	var whale = `              \
               \
     .-'        \     
'--./ /     _.---.
'-,  (__..-'       \\
   \\          .     |
    ',.__.   ,__.--/
     '._/_.'___.-'
`

	var cat = `     \   
      \
     .ﾊ,,ﾊ
     ( ﾟωﾟ)
     |つ  つ
     |    |
     |    |
     |    |
     |    |
     |    |
     |    |
     |    |
     |    |
     |    |
     |    |
     |    |
     |    |
     |    |
     |    |
     |    |
     |    |
     |    |
     |    |
     U "  U
`

	switch name {
	case "cow":
		fmt.Println(cow)
	case "stegosaurus":
		fmt.Println(stegosaurus)
	case "alligator":
		fmt.Println(alligator)
	case "whale":
		fmt.Println(whale)
	case "cat":
		fmt.Println(cat)
	default:
		fmt.Println("Unknown figure")
	}
}

func main() {
	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes")
		fmt.Println("Usage: fortune | cowsay")
		return
	}

	var figure string
	flag.StringVar(
		&figure,
		"f",
		"cow",
		"the figure name. Valid values are `cow`, `stegosaurus`, alligator, whale, cat",
		)
	flag.Parse()

	var lines []string

	reader := bufio.NewReader(os.Stdin)

	for {
		line, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		var res string
		for i, r := range string(line) {
			res = res + string(r)
			if (i + 1) % 150 == 0 || i == utf8.RuneCountInString(string(line)) - 1 {
				lines = append(lines, res)
				res = ""
			}
		}
	}

	lines = tabsToSpaces(lines)
	maxwidth := calculateMaxWidth(lines)
	messages := normalizeStringLength(lines, maxwidth)
	balloon := buildBalloon(messages, maxwidth)
	fmt.Println(balloon)
	printFigure(figure)
}
