package randlang

import (
	"math/rand"
	"time"
)

// RandomLanguageCode генерирует случайное обозначение языка из заданного списка
func RandomLanguageCode() string {
	languages := []string{"ru", "eng", "es", "fr", "de", "it", "jp"}
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(languages))
	return languages[randomIndex]
}
