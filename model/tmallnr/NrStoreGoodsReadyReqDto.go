package tmallnr

// NrStoreGoodsReadyReqDto 结构体
type NrStoreGoodsReadyReqDto struct {
	// 配送人员的姓名
	PerformerName string `json:"performer_name,omitempty" xml:"performer_name,omitempty"`
	// 配送人员的电话
	PerformerPhone string `json:"performer_phone,omitempty" xml:"performer_phone,omitempty"`
	// 发货编码
	CompanyOrderNo string `json:"company_order_no,omitempty" xml:"company_order_no,omitempty"`
	// 发货公司
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 业务标识，dss标识定时送业务；jsd表示极速达业务
	BizIdentity string `json:"biz_identity,omitempty" xml:"biz_identity,omitempty"`
	// 发货公司编码
	CompanyCode string `json:"company_code,omitempty" xml:"company_code,omitempty"`
	// 交易主订单号
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
}
