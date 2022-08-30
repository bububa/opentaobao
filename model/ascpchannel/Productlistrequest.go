package ascpchannel

// Productlistrequest 结构体
type Productlistrequest struct {
	// 经营模式
	SalesMode []string `json:"sales_mode,omitempty" xml:"sales_mode>string,omitempty"`
	// 渠道编码
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 二级渠道编码
	SubChannelCode string `json:"sub_channel_code,omitempty" xml:"sub_channel_code,omitempty"`
	// 商家编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 页码
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
