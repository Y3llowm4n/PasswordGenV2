package passwordgen

import	(
	"flag"
	"bytes"
	"time"
	"math/rand"
)

// global variable which includes all the letters, numbers and tokens.
var (
	LowerCount     int
	UpperCount     int
	SpecialCount   int
	LengthPassword int
)

// here's all the input that the user can specify.
func init() {
	flag.IntVar(&LowerCount, "l", 0, "Number of lowercase letters in password")
	flag.IntVar(&UpperCount, "u", 0, "Number of uppercase letters in password")
	flag.IntVar(&SpecialCount, "s", 0, "Number of special characters in password")
	flag.IntVar(&LengthPassword, "p", 0, "Give a value for the length off your password")
	flag.Parse()
}
// Generate customized password with given inputs by user.
func GenPassword(lowerCount, upperCount, specialCount int) string {

	lowercase := []rune("abcdefghijklmnopqrstuvwxyz")
	uppercase := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	special := []rune("1234567890~!@#$%^&*()")

	rand.Seed(time.Now().UnixNano())

	//save the password inside the buffer
	var buffer bytes.Buffer

	//add the given amount of lowercase characters.
	for i := 0; i < lowerCount; i++ {
		buffer.WriteRune(lowercase[rand.Intn(len(lowercase))])
	}

	//add the given amount of uppercase characters.
	for i := 0; i < upperCount; i++ {
		buffer.WriteRune(uppercase[rand.Intn(len(uppercase))])
	}

	//add the given amount of special characters.
	for i := 0; i < specialCount; i++ {
		buffer.WriteRune(special[rand.Intn(len(special))])
	}

	// adds random character from the three options to meet the required length.
	for buffer.Len() < LengthPassword {
		switch rand.Intn(3) {
		case 0:
			buffer.WriteRune(lowercase[rand.Intn(len(lowercase))])
		case 1:
			buffer.WriteRune(uppercase[rand.Intn(len(uppercase))])
		case 2:
			buffer.WriteRune(special[rand.Intn(len(special))])
		}
	}

	return buffer.String()
	
}
