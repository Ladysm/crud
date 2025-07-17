package validators

import "unicode"

// validador de género
func IsValidGenre(genre string) bool {
	// aqui colocco en un array los generos
	//map es una estructura de datos que asocia claves (keys) con valores (values).
	validGenres := map[string]bool{
		"Fantasy": true,
		"Horror":  true,
		"Mystery": true,
		"Romance": true,
	}
	return validGenres[genre]

}

func IsValidLanguage(language string) bool {
	validLanguage := map[string]bool{
		"English":  true,
		"Spanish":  true,
		"French":   true,
		"German":   true,
		"Italian":  true,
		"Chinese":  true,
		"Japanese": true,
		"Other":    true,
	}
	return validLanguage[language]
}
func IsValidPublishedYear(publishedYear int) bool {
	return publishedYear >= 1000 && publishedYear <= 2100

}
func IsValidPageCount(pageCount int) bool {
	return pageCount > 0
}
func IsValidAvailableCopies(copies int) bool {
	return copies >= 0
}
func IsValidISBN(isbn string) bool {
	if len(isbn) < 10 || len(isbn) > 20 {
		return false
	}
	for _, c := range isbn {
		if unicode.IsDigit(c) && c != '-' {
			return false
		}
	}
	return true
}
