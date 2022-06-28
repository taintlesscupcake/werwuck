package extracthangul

var (
	hangulCHO = []string{"ㄱ", "ㄲ", "ㄴ", "ㄷ", "ㄸ", "ㄹ", "ㅁ", "ㅂ", "ㅃ", "ㅅ", "ㅆ", "ㅇ", "ㅈ", "ㅉ", "ㅊ", "ㅋ", "ㅌ", "ㅍ", "ㅎ"}
	hangulJUN = []string{"ㅏ", "ㅐ", "ㅑ", "ㅒ", "ㅓ", "ㅔ", "ㅕ", "ㅖ", "ㅗ", "ㅘ", "ㅙ", "ㅚ", "ㅛ", "ㅜ", "ㅝ", "ㅞ", "ㅟ", "ㅠ", "ㅡ", "ㅢ", "ㅣ"}
	hangulJON = []string{"", "ㄱ", "ㄲ", "ㄳ", "ㄴ", "ㄵ", "ㄶ", "ㄷ", "ㄹ", "ㄺ", "ㄻ", "ㄼ", "ㄽ", "ㄾ", "ㄿ", "ㅀ", "ㅁ", "ㅂ", "ㅄ", "ㅅ", "ㅆ", "ㅇ", "ㅈ", "ㅊ", "ㅋ", "ㅌ", "ㅍ", "ㅎ"}
)

const (
	// 한글의 시작은 '가' 끝은 '힣'
	hangulBASE = rune('가')
	hangulEND  = rune('힣')

	// 자음은 29개, 모음은 20개
	hangulJA = rune('ㄱ')
	hangulMO = rune('ㅏ')
)

func ExtractHangul(str string) string {
	var result string
	for _, r := range str {
		if hangulBASE <= r && r <= hangulEND {
			base := r - hangulBASE
			cho := base / 588
			jun := (base - cho*588) / 28
			jong := base % 28
			result += hangulCHO[cho] + hangulJUN[jun] + hangulJON[jong]
		} else {
			result += string(r)
		}
	}
	return result
}
