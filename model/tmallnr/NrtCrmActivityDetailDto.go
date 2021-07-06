package tmallnr

// NrtCrmActivityDetailDto 结构体
type NrtCrmActivityDetailDto struct {
	// 状态
	StatusStr string `json:"status_str,omitempty" xml:"status_str,omitempty"`
	// 结束时间
	EndTimeStr string `json:"end_time_str,omitempty" xml:"end_time_str,omitempty"`
	// 开始时间
	StartTimeStr string `json:"start_time_str,omitempty" xml:"start_time_str,omitempty"`
	// 规则
	Rule string `json:"rule,omitempty" xml:"rule,omitempty"`
	// 摊位名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 导购员名称
	GuiderName string `json:"guider_name,omitempty" xml:"guider_name,omitempty"`
	// 活动标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 摊位地址
	StoreAddress string `json:"store_address,omitempty" xml:"store_address,omitempty"`
	// 留资人数
	GuiderCustomerNum int64 `json:"guider_customer_num,omitempty" xml:"guider_customer_num,omitempty"`
	// 总留资人数
	TotalCustomerNum int64 `json:"total_customer_num,omitempty" xml:"total_customer_num,omitempty"`
	// 总uv
	TotalUv int64 `json:"total_uv,omitempty" xml:"total_uv,omitempty"`
	// 总pv
	TotalPv int64 `json:"total_pv,omitempty" xml:"total_pv,omitempty"`
	// 摊位id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 导购员员工id
	GuiderId int64 `json:"guider_id,omitempty" xml:"guider_id,omitempty"`
	// 活动ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 100:已发布，-100:失效，200:已结束
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
