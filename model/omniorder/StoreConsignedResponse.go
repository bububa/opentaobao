package omniorder

// StoreConsignedResponse 结构体
type StoreConsignedResponse struct {
	// shortId
	ShortId string `json:"short_id,omitempty" xml:"short_id,omitempty"`
	// mailNo
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// gotCode
	GotCode string `json:"got_code,omitempty" xml:"got_code,omitempty"`
}
