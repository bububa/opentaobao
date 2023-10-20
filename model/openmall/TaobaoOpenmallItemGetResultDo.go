package openmall

// TaobaoopenmallitemgetResultDo 结构体
type TaobaoopenmallitemgetResultDo struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 商品
	Item *TopItemVo `json:"item,omitempty" xml:"item,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
