package product

// TmallItemStoreSchemaUpdateApiResult 结构体
type TmallItemStoreSchemaUpdateApiResult struct {
	// 错误信息
	ErMsg string `json:"er_msg,omitempty" xml:"er_msg,omitempty"`
	// 错误码
	ErCode string `json:"er_code,omitempty" xml:"er_code,omitempty"`
	// 错误信息
	MappedErrorMessages string `json:"mapped_error_messages,omitempty" xml:"mapped_error_messages,omitempty"`
	// 商品ID
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
	// 成功
	Error bool `json:"error,omitempty" xml:"error,omitempty"`
}
