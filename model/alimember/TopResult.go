package alimember

// TopResult 
type TopResult struct {
    // code，返回码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // message，返回信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
}