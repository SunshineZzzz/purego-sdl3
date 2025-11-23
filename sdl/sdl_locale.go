package sdl

import (
	"unsafe"

	"github.com/SunshineZzzz/purego-sdl3/internal/convert"
)

type Locale struct {
	Language, Country string
}

// GetPreferredLocales reports the user's preferred locale.
func GetPreferredLocales() []*Locale {
	var count int32
	locales := sdlGetPreferredLocales(&count)
	if locales == nil {
		return nil
	}
	defer Free(unsafe.Pointer(locales))

	result := make([]*Locale, count)

	for i, v := range unsafe.Slice(locales, count) {
		locale := new(Locale)
		locale.Language = convert.ToString(v.language)
		locale.Country = convert.ToString(v.country)
		result[i] = locale
	}

	return result
}
