package aesolution

// GlobalDescription 
type GlobalDescription struct {
    // locale of the descripiton
    Locale   string `json:"locale,omitempty" xml:"locale,omitempty"`
    // mobile detail
    MobileDetail   string `json:"mobile_detail,omitempty" xml:"mobile_detail,omitempty"`
    // web detail
    WebDetail   string `json:"web_detail,omitempty" xml:"web_detail,omitempty"`
}
