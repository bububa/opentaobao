package security

// RealNameResult 
type RealNameResult struct {

    // checkCode
    CheckCode   string `json:"check_code,omitempty"`

    // checkMessage
    CheckMessage   string `json:"check_message,omitempty"`

    // match
    Match   bool `json:"match,omitempty"`

}
