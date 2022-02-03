package backtracking

var letters = map[string]string{
	"0": "0",
	"1": "1",
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

// Time: O(n*4^n); space: O(n*4^n); n = len(phoneNumber)
func PhoneNumberMnemonics(phoneNumber string) []string {
	var backtrack func(int, string) []string
	backtrack = func(i int, chosen string) []string {
		if i == len(phoneNumber) {
			return []string{chosen}
		}
		list := []string{}
		digit := string(phoneNumber[i])
		for _, c := range letters[digit] {
			list = append(list, backtrack(i+1, chosen+string(c))...)
		}

		return list
	}

	return backtrack(0, "")
}
