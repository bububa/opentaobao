package opentrade

// ItemBindResult 结构体
type ItemBindResult struct {
	// 绑定异常时的错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 是否绑定成功
	BindOk bool `json:"bind_ok,omitempty" xml:"bind_ok,omitempty"`
}
