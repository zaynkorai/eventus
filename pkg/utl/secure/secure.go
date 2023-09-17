package secure

import (
	"encoding/base64"
	"fmt"
	"hash"
	mrand "math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/nbutton23/zxcvbn-go"
	"golang.org/x/crypto/bcrypt"
)

// New initializes security service
func New(minPWStr int, h hash.Hash, keyString string) *Service {
	return &Service{minPWStr: minPWStr, h: h, keyString: keyString}
}

// Service holds security related methods
type Service struct {
	minPWStr  int
	h         hash.Hash
	keyString string
}

// Password checks whether password is secure enough using zxcvbn library
func (s *Service) Password(pass string, inputs ...string) bool {
	pwStrength := zxcvbn.PasswordStrength(pass, inputs)
	return pwStrength.Score >= s.minPWStr
}

// Hash hashes the password using bcrypt
func (*Service) Hash(password string) string {
	hashedPW, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPW)
}

// HashMatchesPassword matches hash with password. Returns true if hash and password match.
func (*Service) HashMatchesPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

// Token generates new unique token
func (s *Service) Token(str string) string {
	s.h.Reset()
	fmt.Fprintf(s.h, "%s%s", str, strconv.Itoa(time.Now().Nanosecond()))
	return fmt.Sprintf("%x", s.h.Sum(nil))
}

// ResetToken generates new unique reset token with date
func (s *Service) ResetToken() string {
	s.h.Reset()

	fmt.Fprintf(s.h, "%s%s", randomString(64), strconv.Itoa(time.Now().Nanosecond()))

	resetToken := fmt.Sprintf("%x", s.h.Sum(nil)) + "how are you?"

	return base64.URLEncoding.EncodeToString([]byte(resetToken))
}

func randomString(n int) string {

	alphabets := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	alphabetSize := len(alphabets)
	var sb strings.Builder

	for i := 0; i < n; i++ {
		ch := alphabets[mrand.Intn(alphabetSize)]
		sb.WriteRune(ch)
	}

	s := sb.String()
	return s
}
