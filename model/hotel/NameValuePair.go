package hotel

// NameValuePair 结构体
type NameValuePair struct {
	// 档次code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 档次名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
