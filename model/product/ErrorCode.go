package product

// ErrorCode 结构体
type ErrorCode struct {
	// 错误码信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码标识
	MesCode string `json:"mes_code,omitempty" xml:"mes_code,omitempty"`
}
