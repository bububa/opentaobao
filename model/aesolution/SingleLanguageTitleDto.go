package aesolution

// SingleLanguageTitleDto 结构体
type SingleLanguageTitleDto struct {
	// Support: en(English) ru(Russian) es(Spanish) fr(French) it(Italian) tr(Turkish) pt(Portuguese) de(German) nl(Dutch) in(Indonesian) ar(Arabic) ja(Japanese) ko(Korean) th(Thai) vi(Vietnamese) iw(Hebrew).Must contains the original locale.
	Language string `json:"language,omitempty" xml:"language,omitempty"`
	// subject, maximum 218 characters.
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
}
