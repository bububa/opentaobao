package b2bcert

// Response 
type Response struct {
    // data
    Data   string `json:"data,omitempty" xml:"data,omitempty"`
    // code
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // msg
    Msg   string `json:"msg,omitempty" xml:"msg,omitempty"`
}
