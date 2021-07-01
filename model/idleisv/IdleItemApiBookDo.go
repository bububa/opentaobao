package idleisv

// IdleItemApiBookDo 结构体
type IdleItemApiBookDo struct {
	// 图书ISBN码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 图书ISBN码对应的书名等信息
	BarcodeName string `json:"barcode_name,omitempty" xml:"barcode_name,omitempty"`
}
