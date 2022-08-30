package ascpchannel

// ProductList 结构体
type ProductList struct {
	// 经营模式
	SalesModeList []string `json:"sales_mode_list,omitempty" xml:"sales_mode_list>string,omitempty"`
	// 图片
	PictureList []string `json:"picture_list,omitempty" xml:"picture_list>string,omitempty"`
	// 产品标题名称
	ProductTitle string `json:"product_title,omitempty" xml:"product_title,omitempty"`
	// 外部商家编码
	OutNo string `json:"out_no,omitempty" xml:"out_no,omitempty"`
	// 二级渠道编码
	SubChannelCode string `json:"sub_channel_code,omitempty" xml:"sub_channel_code,omitempty"`
	// 渠道编码
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 产品 id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 产品类型
	ProductType int64 `json:"product_type,omitempty" xml:"product_type,omitempty"`
}
