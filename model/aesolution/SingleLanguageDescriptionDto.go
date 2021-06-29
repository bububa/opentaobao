package aesolution

// SingleLanguageDescriptionDTO 
type SingleLanguageDescriptionDTO struct {
    // Support: en(English) ru(Russian) es(Spanish) fr(French) it(Italian) tr(Turkish) pt(Portuguese) de(German) nl(Dutch) in(Indonesian) ar(Arabic) ja(Japanese) ko(Korean) th(Thai) vi(Vietnamese) iw(Hebrew)
    Language   string `json:"language,omitempty" xml:"language,omitempty"`
    // mobile detail for  this language, please check the format here https://developers.aliexpress.com/en/doc.htm?docId=109534&docType=1
    MobileDetail   string `json:"mobile_detail,omitempty" xml:"mobile_detail,omitempty"`
    // web detail for this language, please check the format here: https://developers.aliexpress.com/en/doc.htm?docId=109534&docType=1
    WebDetail   string `json:"web_detail,omitempty" xml:"web_detail,omitempty"`
}
