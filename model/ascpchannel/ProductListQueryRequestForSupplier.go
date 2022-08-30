package ascpchannel

// ProductListQueryRequestForSupplier 结构体
type ProductListQueryRequestForSupplier struct {
	// 经营模式,agent 表示代销，dealer 表示经销
	SalesMode []string `json:"sales_mode,omitempty" xml:"sales_mode>string,omitempty"`
	// 二级渠道编码
	SubChannelCode string `json:"sub_channel_code,omitempty" xml:"sub_channel_code,omitempty"`
	// 渠道编码
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 页号，从1开始
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
