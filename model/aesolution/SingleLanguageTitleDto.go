package aesolution

// SingleLanguageTitleDto 
type SingleLanguageTitleDto struct {

    // Support: en(English) ru(Russian) es(Spanish) fr(French) it(Italian) tr(Turkish) pt(Portuguese) de(German) nl(Dutch) in(Indonesian) ar(Arabic) ja(Japanese) ko(Korean) th(Thai) vi(Vietnamese) iw(Hebrew)
    
    Language   string `json:"language,omitempty" xml:"language,omitempty"`
    

    // subject, maximum 128 characters.
    
    Subject   string `json:"subject,omitempty" xml:"subject,omitempty"`
    

}
