package tmallnr

// NrtCrmActivityStatisticsDataReq 结构体
type NrtCrmActivityStatisticsDataReq struct {
	// 数据上传时间
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 业务身份
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 浏览人数
	Uv int64 `json:"uv,omitempty" xml:"uv,omitempty"`
	// 活动id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 页面访问数
	Pv int64 `json:"pv,omitempty" xml:"pv,omitempty"`
	// 摊位id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 导购员id
	GuiderId int64 `json:"guider_id,omitempty" xml:"guider_id,omitempty"`
}
