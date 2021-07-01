package opentrade

// ItemUnBindResult 结构体
type ItemUnBindResult struct {
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 解绑异常时的错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 是否解绑成功
	BindOk bool `json:"bind_ok,omitempty" xml:"bind_ok,omitempty"`
}
