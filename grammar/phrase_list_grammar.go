package grammar

import (
	"github.com/rnhdev2/cognitive-services-speech-sdk-go/speech"
)

// #include <speechapi_c_common.h>
// #include <speechapi_c_grammar.h>
import "C"

type PhraseListGrammar struct {
	handle C.SPXHANDLE
}

func GetPhraseListGrammarFromRecognizer(recognizer *speech.SpeechRecognizer) (*PhraseListGrammar, error) {
	hreco := uintptr2handle(recognizer.GetHandle())
	_ = hreco
	var hgrammar C.SPXGRAMMARHANDLE
	hgrammar = C.SPXHANDLE_INVALID
	n := C.CString("")

	C.phrase_list_grammar_from_recognizer_by_name(&hgrammar, hreco, n)

	return &PhraseListGrammar{handle: hgrammar}, nil
}

func (p *PhraseListGrammar) AddPhrase(value string) {
	phrase := C.CString(value)
	var hphrase C.SPXGRAMMARHANDLE
	hphrase = C.SPXHANDLE_INVALID

	C.grammar_phrase_create_from_text(&hphrase, phrase)
	C.phrase_list_grammar_add_phrase(p.handle, hphrase)
}

func (p *PhraseListGrammar) Clear() {
	C.phrase_list_grammar_clear(p.handle)
}
