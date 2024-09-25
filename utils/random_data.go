package utils

import (
	"fmt"
	"math/rand"
)

func generateRandomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s := make([]rune, length)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func generateRandomName() string {
	names := []string{"Алексей", "Мария", "Иван", "Елена", "Петр"}
	return names[rand.Intn(len(names))]
}

func generateRandomPhoneNumber() string {
	return "+972" + fmt.Sprintf("%7d", rand.Intn(10000000))
}

func generateRandomZip() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func generateRandomCity() string {
	cities := []string{"Москва", "Санкт-Петербург", "Новосибирск", "Екатеринбург", "Казань"}
	return cities[rand.Intn(len(cities))]
}

func generateRandomAddress() string {
	streets := []string{"Ленина", "Советская", "Мира", "Пушкина", "Кирова"}
	return fmt.Sprintf("%d %s", rand.Intn(1000), streets[rand.Intn(len(streets))])
}

func generateRandomRegion() string {
	regions := []string{"Центральный", "Северо-Западный", "Южный", "Уральский", "Сибирский"}
	return regions[rand.Intn(len(regions))]
}

func generateRandomEmail() string {
	return generateRandomString(10) + "@gmail.com"
}
