package tmallnr

// NrtCertificateSendDto 结构体
type NrtCertificateSendDto struct {
	// 淘系加密ID
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 业务编码
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 投放渠道
	Channel int64 `json:"channel,omitempty" xml:"channel,omitempty"`
	// 电子凭证模版id
	OutId int64 `json:"out_id,omitempty" xml:"out_id,omitempty"`
}
