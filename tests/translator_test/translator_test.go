package translator_test

import (
	"i18n-go/src/translator"
	"testing"
)

func TestTranslateReturnsCorrectTranslation(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"en", "hello"},
		{"pt-BR", "ola"},
		{"pt-PT", "oi"},
	}

	translator := translator.NewTranslator("en")
	translator.LoadTranslations("../mock_translation_files/en.json")
	translator.LoadTranslations("../mock_translation_files/pt-BR.json")
	translator.LoadTranslations("../mock_translation_files/pt-PT.json")

	for _, test := range tests {
		if got, _ := translator.Translate("greeting", test.input); string(got) != string(test.want) {
			t.Errorf("TestTranslateReturnsCorrectTranslation(%s) = %s, want %s", test.input, got, test.want)
		}
	}
}

func TestReturnDefaultLanguageTranslationWhenReceivedTranslationIsNotFound(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"en", "hello"},
		{"invalid", "hello"},
	}

	translator := translator.NewTranslator("en")
	translator.LoadTranslations("../mock_translation_files/en.json")

	for _, test := range tests {
		if got, _ := translator.Translate("greeting", test.input); string(got) != string(test.want) {
			t.Errorf("TestReturnDefaultLanguageTranslationWhenReceivedTranslationIsNotFound(%s) = %s, want %s", test.input, got, test.want)
		}
	}
}

func TestShouldReturnErrorWhenNoTranslationIsFound(t *testing.T) {
	var tests = []struct {
		inputLang   string
		inputTarget string
		want        translator.TranslationError
	}{
		{"pt-BR", "invalid", translator.TranslationError{Code: 1, Msg: "could not find translation for key invalid on language pt-BR"}},
		{"invalid", "greeting", translator.TranslationError{Code: 1, Msg: "could not find translation for key greeting on language invalid"}},
		{"invalid", "invalid", translator.TranslationError{Code: 1, Msg: "could not find translation for key invalid on language invalid"}},
	}

	translator := translator.NewTranslator("pt-BR")
	translator.LoadTranslations("../mock_translation_files/en.json")

	for _, test := range tests {
		if _, err := translator.Translate(test.inputTarget, test.inputLang); *err != test.want {
			t.Errorf("TestShouldReturnErrorWhenNoTranslationIsFound(%s, %s) = \n\t%v, \nwant \n\t%v", test.inputLang, test.inputTarget, err, &test.want)
		}
	}
}

func TestShouldThrowErrorWhenTranslationFileCantBeParsed(t *testing.T) {
	var tests = []struct {
		input string
		want  translator.TranslationError
	}{
		{"non-existing-file.json", translator.TranslationError{Code: 2, Msg: "could not open file at non-existing-file.json"}},
		{"../mock_translation_files/incorrect.json", translator.TranslationError{Code: 3, Msg: "an error occured while unmarshaling the JSON contents for file at ../mock_translation_files/incorrect.json"}},
	}

	translator := translator.NewTranslator("pt-BR")

	for _, test := range tests {
		if err := translator.LoadTranslations(test.input); *err != test.want {
			t.Errorf("TestShouldThrowErrorWhenTranslationFileCantBeParsed(%s) = \n\t%v, \nwant \n\t%v", test.input, err, &test.want)
		}
	}
}

func TestErrorFormating(t *testing.T) {
	var tests = []struct {
		input translator.TranslationError
		want  string
	}{
		{translator.TranslationError{Code: 0, Msg: "mock error"}, "Code: 0, Msg: mock error"},
	}

	for _, test := range tests {
		if errStr := test.input.Error(); errStr != test.want {
			t.Errorf("TestErrorFormating(%v) = %s, want %s", test.input, errStr, test.want)
		}
	}
}
