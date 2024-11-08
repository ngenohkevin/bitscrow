package config

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func GenerateRandomID() uuid.UUID {
	return uuid.New()
}

func RandomBitcoinAddress() string {
	return "bc1" + RandomString(29)
}

func RandomStatus() pgtype.Text {
	status := []string{"pending", "confirmed", "disputed", "released"}
	n := len(status)

	// Randomly select a status
	selectedStatus := status[rand.Intn(n)]

	// Create a pgtype.Text instance with the selected status
	textStatus := pgtype.Text{
		String: selectedStatus,
		Valid:  true, // Mark as valid since it has a value
	}

	return textStatus
}

func RandomAmount() pgtype.Numeric {
	randomInt := int64(RandomInt(1000, 100000)) // Range between 1000 and 9999 cents

	var numeric pgtype.Numeric
	numeric.Int = big.NewInt(randomInt)
	numeric.Exp = -2 // Set scale to two decimal places
	numeric.Valid = true

	return numeric
}
