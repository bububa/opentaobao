package omniorder

// StoreConsignedResponse 结构体
type StoreConsignedResponse struct {
	// mailNo
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// shortId
	ShortId string `json:"short_id,omitempty" xml:"short_id,omitempty"`
	// gotCode
	GotCode string `json:"got_code,omitempty" xml:"got_code,omitempty"`
}
