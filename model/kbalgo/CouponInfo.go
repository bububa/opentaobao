package kbalgo

// CouponInfo 结构体
type CouponInfo struct {
	// schema
	Schema *Schema `json:"schema,omitempty" xml:"schema,omitempty"`
	// ext
	Ext string `json:"ext,omitempty" xml:"ext,omitempty"`
	// 开始时间
	GmtStart string `json:"gmt_start,omitempty" xml:"gmt_start,omitempty"`
	// 图片url
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// 价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 销量
	SalesNum string `json:"sales_num,omitempty" xml:"sales_num,omitempty"`
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 原价
	OriginPrice string `json:"origin_price,omitempty" xml:"origin_price,omitempty"`
	// 结束时间
	GmtEnd string `json:"gmt_end,omitempty" xml:"gmt_end,omitempty"`
	// 商品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 商品id
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
}
