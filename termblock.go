// Turn Array to visiable terminal block map

package termblock

import (
	"fmt"
	"strconv"
	
)

var CharsetBlocks = `▉▊▋▌▍▎▏▎▍▌▋▊▉`
var CharsetDots    = `⠁⠁⠉⠙⠚⠒⠂⠂⠒⠲⠴⠤⠄⠄⠤⠠⠠⠤⠦⠖⠒⠐⠐⠒⠓⠋⠉⠈⠈`
var CharsetDirections = `←↑→↓`

var Colorset1 = []int{31,32,33,34,35,36,37}
var Colorset2 = []int{31,32,33,34,35,36}

type TermBlock struct {
	alphabet []rune
	color []int
}

func New(alphabet string) *TermBlock {
	tb := &TermBlock{}
	tb.SetAlphabet(alphabet)
	return tb
}

func ColorNew(alphabet string, color []int) *TermBlock {
	tb := &TermBlock{}
	tb.SetColorAlphabet(alphabet, color)
	return tb
}

func (tb *TermBlock) SetAlphabet(frame string) {
	tb.alphabet = []rune(frame)
}

func (tb *TermBlock) SetColorAlphabet(frame string, color []int) {
	tb.alphabet = []rune(frame)
	tb.color = color
}

func (tb *TermBlock) GetAlphabet() string {
	return string(tb.alphabet)
}

func (tb *TermBlock) GetColorAlphabet() (string, []int) {
	return string(tb.alphabet), tb.color
}

func (tb *TermBlock) Translate(myarray []int) string {
	s := ""
	for i := range myarray {
		alphabetPos := myarray[i] % len(tb.alphabet)
		color_start := ""
		if len(tb.color) > 0 {
			colorPos := myarray[i] % len(tb.color)
			color_start = "\033[" + strconv.Itoa(tb.color[colorPos])  + "m"
		}
		s += color_start + string(tb.alphabet[alphabetPos])
	}
	return s
}

func (tb *TermBlock) Print(myarray []int, label string) {
	s := tb.Translate(myarray)
	// fmt.Printf("\r  \033[36m%s\033[m %s ", label, s)
	fmt.Printf("\r  \033[36m%s\033[m %s  \033[m", label, s)
}
