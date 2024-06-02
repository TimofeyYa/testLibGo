package randcolor

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomColorHex генерирует случайный цвет в формате HEX
func RandomColorHex() string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(256)
	g := rand.Intn(256)
	b := rand.Intn(256)
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}
