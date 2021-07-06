package servicecenter

// SubscInfo 结构体
type SubscInfo struct {
	// 订单创建时间
	SubscCreatedTime string `json:"subsc_created_time,omitempty" xml:"subsc_created_time,omitempty"`
	// 服务商名称
	SpName string `json:"sp_name,omitempty" xml:"sp_name,omitempty"`
	// 子账号名称列表
	SubAccountList string `json:"sub_account_list,omitempty" xml:"sub_account_list,omitempty"`
	// 服务开结束时间
	ServiceEndTime string `json:"service_end_time,omitempty" xml:"service_end_time,omitempty"`
	// 商家名称
	SellerName string `json:"seller_name,omitempty" xml:"seller_name,omitempty"`
	// 服务开始时间
	ServiceStartTime string `json:"service_start_time,omitempty" xml:"service_start_time,omitempty"`
	// 订单修改时间
	SubscModifiedTime string `json:"subsc_modified_time,omitempty" xml:"subsc_modified_time,omitempty"`
	// 销售提成
	SaleBonus string `json:"sale_bonus,omitempty" xml:"sale_bonus,omitempty"`
	// 订单id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 订单状态
	ProcessStatus int64 `json:"process_status,omitempty" xml:"process_status,omitempty"`
}
