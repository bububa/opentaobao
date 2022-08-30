package waybill

// ClientInfoDto 结构体
type ClientInfoDto struct {
	// 调用时自定义描述信息
	Description string `json:"description,omitempty" xml:"description,omitempty"`
}
