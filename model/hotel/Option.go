package hotel

// Option 结构体
type Option struct {
	// link
	Link string `json:"link,omitempty" xml:"link,omitempty"`
	// text
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// value
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// active
	Active bool `json:"active,omitempty" xml:"active,omitempty"`
}
