package tmallservice

// Spb2bOderQuery 结构体
type Spb2bOderQuery struct {
	// 开始日期
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 结束日期
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 售卖商昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 服务code
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
	// 服务状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 每页个数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 最后条数
	EndRow int64 `json:"end_row,omitempty" xml:"end_row,omitempty"`
	// 开始条数
	StartRow int64 `json:"start_row,omitempty" xml:"start_row,omitempty"`
	// 售卖商id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 是否分页
	NeedByPage bool `json:"need_by_page,omitempty" xml:"need_by_page,omitempty"`
	// 是否查询新供给ssc订购数据
	NewSupplySubscriberData bool `json:"new_supply_subscriber_data,omitempty" xml:"new_supply_subscriber_data,omitempty"`
}
