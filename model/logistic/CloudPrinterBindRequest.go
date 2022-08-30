package logistic

// CloudPrinterBindRequest 结构体
type CloudPrinterBindRequest struct {
	// 打印机 mac 地址
	Uid string `json:"uid,omitempty" xml:"uid,omitempty"`
	//  验证码
	VerifyCode string `json:"verify_code,omitempty" xml:"verify_code,omitempty"`
}
