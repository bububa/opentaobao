package yunos

// DpResult 
type DpResult struct {
    // code
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`
    // message
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
