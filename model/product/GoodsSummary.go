package product

// GoodsSummary 结构体
type GoodsSummary struct {
	// max_price_cent 产品最高价格
	MaxPriceCent int64 `json:"max_price_cent,omitempty" xml:"max_price_cent,omitempty"`
	// id 产品id
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// unit 产品单位
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// detail_url 产品详情url
	DetailUrl string `json:"detail_url,omitempty" xml:"detail_url,omitempty"`
	// supplier_id 产品对应供应商id
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// subject 产品标题
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
	// min_price_cent 产品最低价
	MinPriceCent int64 `json:"min_price_cent,omitempty" xml:"min_price_cent,omitempty"`
	// thumb_url 产品缩略图
	ThumbUrl string `json:"thumb_url,omitempty" xml:"thumb_url,omitempty"`
	// moq 产品最小起订量
	Moq int64 `json:"moq,omitempty" xml:"moq,omitempty"`
	// buy_now_url  产品下单链接
	BuyNowUrl string `json:"buy_now_url,omitempty" xml:"buy_now_url,omitempty"`
	// currency 产品价格货币单位
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
}
