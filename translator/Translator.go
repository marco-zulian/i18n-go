package translator

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Translator struct {
	translations    map[string]map[string]string
	defaultLanguage string
}

func NewTranslator(defaultLanguage string) *Translator {
	return &Translator{
		translations:    make(map[string]map[string]string),
		defaultLanguage: defaultLanguage,
	}
}

func (translator *Translator) Translate(key string, lang string) (string, *TranslationError) {
	if translations, ok := translator.translations[lang]; ok {
		if translation, ok := translations[key]; ok {
			return translation, nil
		}
	}

	if translations, ok := translator.translations[translator.defaultLanguage]; ok {
		if translation, ok := translations[key]; ok {
			fmt.Printf("WARNING: No translations for %s on language %s was found. Falling back to default %s\n", key, lang, translator.defaultLanguage)
			return translation, nil
		}
	}

	return "", &TranslationError{
		Code: TranslationCode["TranslationNotFound"],
		Msg:  fmt.Sprintf("could not find translation for key %s on language %s", key, lang),
	}
}

func (translator *Translator) LoadTranslations(filePath string) *TranslationError {
	if data, err := os.ReadFile(filePath); err != nil {
		return &TranslationError{
			Code: TranslationCode["FileLoadingError"],
			Msg:  fmt.Sprintf("could not open file at %s", filePath),
		}
	} else {
		var translations map[string]string

		if err := json.Unmarshal([]byte(data), &translations); err != nil {
			return &TranslationError{
				Code: TranslationCode["FileUnmarshalingError"],
				Msg:  fmt.Sprintf("an error occured while unmarshaling the JSON contents for file at %s", filePath),
			}
		}

		fileSegments := strings.Split(filePath, "/")
		language := strings.Split(fileSegments[len(fileSegments)-1], ".")[0]

		translator.translations[language] = translations
		return nil
	}
}
