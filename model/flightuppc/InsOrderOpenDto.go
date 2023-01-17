package flightuppc

// InsOrderOpenDto 结构体
type InsOrderOpenDto struct {
	// 订单详情列表
	InsOrderDetailList []InsOrderOpenDetailDto `json:"ins_order_detail_list,omitempty" xml:"ins_order_detail_list>ins_order_open_detail_dto,omitempty"`
	// 卖家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 保单详情页地址
	PolicyDetailUrl string `json:"policy_detail_url,omitempty" xml:"policy_detail_url,omitempty"`
	// 保单号
	PolicyNo string `json:"policy_no,omitempty" xml:"policy_no,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 保险产品名称
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// features
	Features string `json:"features,omitempty" xml:"features,omitempty"`
	// 保险产品编号
	ProductNo string `json:"product_no,omitempty" xml:"product_no,omitempty"`
	// 保险订单号
	TcOrderId int64 `json:"tc_order_id,omitempty" xml:"tc_order_id,omitempty"`
	// 保险供应方id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 保险份数
	Copies int64 `json:"copies,omitempty" xml:"copies,omitempty"`
	// 保险价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 外部订单号
	OutOrderId int64 `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 保险订单状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 保险产品id
	PremiumId int64 `json:"premium_id,omitempty" xml:"premium_id,omitempty"`
}
