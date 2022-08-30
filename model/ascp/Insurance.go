package ascp

// Insurance 结构体
type Insurance struct {
	// 保险类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 保险金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
}
