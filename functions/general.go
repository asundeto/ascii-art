package ascart

import "fmt"

func ReadFile(content, style string) {
	if style == "" {
		style = "standard"
	}
	var words []string
	var word string
	var cntSlash int
	var arrSlashCnt []int
	var checkLast bool
	var checkFirstWord bool
	arr := ReadWholeFile("arts/" + style + ".txt")
	if arr == nil {
		return
	}
	for i := 0; i < len(content); i++ {
		if content[i] != '\\' {
			word += string(content[i])
			checkFirstWord = true
		} else {
			if i+1 < len(content) && content[i+1] == 'n' {
				if word == "" && !checkFirstWord {
					fmt.Println()
					i++
					continue
				}
				if word != "\n" && word != "" {
					words = append(words, word)
				}
				word = ""
				i++
				cntSlash++
				if i+1 == len(content) {
					checkLast = true
				}
				if i+1 < len(content) {
					if i+2 < len(content) {
						if content[i+1] == '\\' && content[i+2] == 'n' {
							continue
						} else {
							word += string(content[i+1])
							i++
							arrSlashCnt = append(arrSlashCnt, cntSlash)
							cntSlash = 0
						}
					}
				} else {
					arrSlashCnt = append(arrSlashCnt, cntSlash)
				}
			} else {
				word += string(content[i])
			}
		}
		if i+1 == len(content) {
			if word != "\n" && word != "" {
				words = append(words, word)
			}
		}
	}
	Call(checkLast, words, arr, arrSlashCnt)
}

func Call(checkLast bool, words []string, arr []rune, arrSlashCnt []int) {
	for i := 0; i < len(words); i++ {
		if checkLast {
			if i+1 == len(words) {
				Separate(arr, words[i], arrSlashCnt[i], true)
			} else {
				Separate(arr, words[i], arrSlashCnt[i], false)
			}
		} else {
			if i+1 == len(words) {
				Separate(arr, words[i], 0, true)
			} else {
				Separate(arr, words[i], arrSlashCnt[i], false)
			}
		}
	}
}

func Separate(arr []rune, content string, numNewLine int, boly bool) {
	var doubleArr [][]rune
	var res []rune
	if !boly {
		numNewLine -= 1
	}
	indexes := Index(content)
	for i := 0; i < len(indexes); i++ {
		num := indexes[i] * 9
		x := 0
		j := 0
		for j < len(arr) {
			if x < num {
				if arr[j] == 10 || arr[j] == 13 {
					x++
				}
			} else {
				break
			}
			j++
		}
		y := 0
		for k := j; k < len(arr); k++ {
			if arr[k] == 10 {
				y++
			}
			if y < 10 {
				res = append(res, arr[k])
			}
		}
		doubleArr = append(doubleArr, res)

		res = []rune{}
	}

	var resStr string
	if len(doubleArr) > 1 {
		resStr = RuneToString(Connect(doubleArr))
	} else {
		if len(doubleArr) != 0 {
			resStr = RuneToString(doubleArr[0])
		}
	}

	var newRes string
	cnt := 0
	for i := 0; i < len(resStr); i++ {
		if cnt > 0 && cnt < 9 {
			newRes += string(resStr[i])
		}
		if resStr[i] == '\n' {
			cnt++
		}
	}
	fmt.Print(newRes)
	for i := 0; i < numNewLine; i++ {
		fmt.Print("\n")
	}
}

func Connect(arr [][]rune) []rune {
	var res []rune
	var nums []int
	for i := 0; i < len(arr); i++ {
		nums = append(nums, 0)
	}
	i := 0
	j := 0
	cnt := 0
	for true {
		if j+1 < len(arr[i]) {
			j = nums[i]
			for true {
				if arr[i][j] == 10 {
					nums[i] = j + 1
					j = 0
					break
				}
				res = append(res, arr[i][j])
				j++
			}
		}
		myCheck := false
		if i+1 == len(arr) && cnt < 9 {
			cnt++
			res = append(res, 10)
		}
		if i+1 == len(arr) && cnt < 9 {
			i = 0
			myCheck = true
		}
		if i+1 == len(arr) {
			break
		}
		if !myCheck {
			i++
		}
	}
	return res
}

func RuneToString(r []rune) string {
	var res string
	for _, i := range r {
		res += string(i)
	}
	return res
}

func Index(s string) []int {
	var res []int
	for i := 0; i < len(s); i++ {
		res = append(res, int(s[i]-32))
	}
	return res
}
