package icburfq

// RfqRequestSearchCondDto 结构体
type RfqRequestSearchCondDto struct {
	// 关键词
	SearchText string `json:"search_text,omitempty" xml:"search_text,omitempty"`
	// 国家简称
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 类目
	CategoryId string `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 每页显示个数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 过滤RFQ发送时间秒级别的
	OpenTime int64 `json:"open_time,omitempty" xml:"open_time,omitempty"`
	// RFQ发布到现在的结束时间秒级别
	CloseTime int64 `json:"close_time,omitempty" xml:"close_time,omitempty"`
	// 最小量
	QuantityMin int64 `json:"quantity_min,omitempty" xml:"quantity_min,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 最大量
	QuantityMax int64 `json:"quantity_max,omitempty" xml:"quantity_max,omitempty"`
	// 是否有附件
	Attachment bool `json:"attachment,omitempty" xml:"attachment,omitempty"`
	// 是否有图片
	Photo bool `json:"photo,omitempty" xml:"photo,omitempty"`
	// 是否报满RFQ
	FullQuote bool `json:"full_quote,omitempty" xml:"full_quote,omitempty"`
	// 是否限免RFQ
	ZeroQuotation bool `json:"zero_quotation,omitempty" xml:"zero_quotation,omitempty"`
	// 是否过滤已报价
	FilterQuoted bool `json:"filter_quoted,omitempty" xml:"filter_quoted,omitempty"`
}
