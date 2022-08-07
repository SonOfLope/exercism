//Package bob provides conversational answers
package bob

import "strings"

//Hey will make conversation following normal rules regarding sentence punctuation in English.
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)
	var lastWord string
	if len(remark) == 0 {
		lastWord = ""
	} else {
		lastWord = remark[len(remark)-1:]
	}

	if lastWord == "?" {
		if strings.ToUpper(remark) == remark && isLetter(remark) {
			return "Calm down, I know what I'm doing!"
		} else {
			return "Sure."
		}
	} else if strings.ToUpper(remark) == remark && isLetter(remark) {
		return "Whoa, chill out!"
	} else if remark == "" {
		return "Fine. Be that way!"
	} else {
		return "Whatever."
	}
}

func isLetter(s string) bool {
	for _, r := range s {
		if ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') {
			return true
		}
	}
	return false
}
