package alsc

// OpenGiftVoucherOpenInfo 结构体
type OpenGiftVoucherOpenInfo struct {
	// 券模板ID
	VoucherId string `json:"voucher_id,omitempty" xml:"voucher_id,omitempty"`
	// 数量
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
}
