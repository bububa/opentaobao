package wms

// Invoinceconfirmlist 结构体
type Invoinceconfirmlist struct {
	// 发票确认信息
	InvoinceConfirm *Invoinceconfirm `json:"invoince_confirm,omitempty" xml:"invoince_confirm,omitempty"`
}
