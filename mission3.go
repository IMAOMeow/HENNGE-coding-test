package main

import (
	"io"
	"net/http"
	"strings"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"

	"encoding/base32"
	"encoding/json"
	"fmt"
	"time"
)

// Generates Passcode using a UTF-8 (not base32) secret and custom parameters
func GeneratePassCode(utf8string string) string {
	secret := base32.StdEncoding.EncodeToString([]byte(utf8string))
	passcode, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Period:    30,
		Skew:      0,
		Digits:    10,
		Algorithm: otp.AlgorithmSHA512,
	})
	if err != nil {
		panic(err)
	}

	return passcode
}

func main() {
	key := "<replace_to_ur_own>"
	result := GeneratePassCode(key)

	hc := http.Client{}
	req, _ := http.NewRequest("POST",
		"https://api.challenge.hennge.com/challenges/003",
		strings.NewReader(`{"github_url": "<ur_own_gist_url>","contact_email": "<email_adress>","solution_language": "golang"}`))

	req.SetBasicAuth("<email_adress>", result)
	req.Header.Add("Content-Type", "application/json")

	res, err := hc.Do(req)
	if err != nil {
		b, _ := json.Marshal(err)
		fmt.Println(string(b))
		return
	}

	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
}
