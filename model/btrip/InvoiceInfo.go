package btrip

// InvoiceInfo 结构体
type InvoiceInfo struct {
	// 发票抬头
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 发票id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
