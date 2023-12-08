package pm

import (
	secrand "crypto/rand"
	"log"
	"math/big"
	mrand "math/rand"
	"strings"
	"time"
)

func NewPass(letterCharset, numberCharset, specialCharset string,
	minUppercase, minLowercase, minNumber, minSpecials, size int) Pass {

	return Pass{
		Config: Config{
			LetterCharset:  letterCharset,
			NumberCharset:  numberCharset,
			SpecialCharset: specialCharset,
		},
		Upper:   minUppercase,
		Lower:   minLowercase,
		Special: minSpecials,
		Number:  minNumber,
		Length:  size,
	}
}

type Config struct {
	LetterCharset  string
	NumberCharset  string
	SpecialCharset string
}

type Pass struct {
	Config  Config
	Lower   int
	Upper   int
	Number  int
	Special int
	Length  int
}

func (p Pass) Generate() string {

	lowerCharset := strings.ToLower(p.Config.LetterCharset)
	upperCharset := strings.ToUpper(p.Config.LetterCharset)
	numberCharset := p.Config.NumberCharset
	specialCharset := p.Config.SpecialCharset

	min := p.Lower + p.Upper + p.Number + p.Special

	if min > p.Length {
		log.Fatalf("count (%d) < length (%d)", min, p.Length)
	}

	sb := strings.Builder{}

	pickAndWrite(&sb, lowerCharset, p.Lower)
	pickAndWrite(&sb, upperCharset, p.Upper)
	pickAndWrite(&sb, numberCharset, p.Number)
	pickAndWrite(&sb, specialCharset, p.Special)

	allChars := strings.Join([]string{
		lowerCharset,
		upperCharset,
		numberCharset,
		specialCharset,
	}, "")

	rest := p.Length - len(sb.String())
	pickAndWrite(&sb, allChars, rest)

	passrunes := []rune(sb.String())

	mrand.Seed(time.Now().Unix())

	for i := 0; i < 50; i++ {
		shuffle(passrunes)
	}

	return string(passrunes)
}

func shuffle(passrunes []rune) {
	mrand.Shuffle(len(passrunes), func(i, j int) {
		passrunes[i], passrunes[j] = passrunes[j], passrunes[i]
	})
}

func pickAndWrite(sb *strings.Builder, charset string, count int) {
	if len(charset) == 0 {
		return
	}
	for i := 0; i < count; i++ {
		n, err := secrand.Int(secrand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			log.Fatal("Erro ao gerar número aleatório")
		}

		sb.WriteByte(charset[n.Int64()])
	}

}
