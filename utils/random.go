package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//alphabets that contains all acceptable
const alphabets = "abcdefghijklmonpqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

//RandomInt generate a random integer between min and max
func RandomInt(min, max int32) int32 {
	return min + rand.Int31n(max-min+1) // return a random integer between 0- (max-min)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)
	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()

}

func GenerateAlphanumericCode(size int) string {
	b := make([]byte, size)
	rand.Read(b)
	result := ""
	for i := 0; i < size; i++ {
		curValue := b[i]
		for curValue < 48 || curValue > 57 && curValue < 65 || curValue > 90 && curValue < 97 || curValue > 122 {
			if curValue < 48 {
				curValue += 5
			}
			if curValue > 57 {
				curValue += 8
			}
			if curValue > 122 {
				curValue -= 5
			}
		}
		result += string(curValue)
	}
	return result
}

//Generate a random owner name
func RandomName() string {
	return RandomString(10)
}

//generate a random password
func RandomPassword() string {
	return GenerateAlphanumericCode(50)
}

//generate a random axxess token
func RandomAccessToken() string {
	return GenerateAlphanumericCode(150)
}

//generate a random refresh token
func RandomRefreshToken() string {
	return GenerateAlphanumericCode(150)
}

//Rancom role selects a random role
func RandomRole() string {
	roles := []string{"Owner", "Admin", "Master"}
	n := len(roles)
	return roles[rand.Intn(n)]
}

func RandomPhoneNumber() uint32 {
	return uint32(RandomInt(1, 99999999) * RandomInt(1, 999999999))
}

func RandomEmail() string {
	return fmt.Sprintf("%v@random.com", RandomString(10))
}

func RandomID() int32 {
	return RandomInt(0, 999)
}

func RandomUserName() string {
	return RandomString(3) + "_" + GenerateAlphanumericCode(5)
}
