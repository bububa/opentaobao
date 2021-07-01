package icbudropshipping

// Phone 结构体
type Phone struct {
	// fax area
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// fax country
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// fax number
	Number string `json:"number,omitempty" xml:"number,omitempty"`
}
