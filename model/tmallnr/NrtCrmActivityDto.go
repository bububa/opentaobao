package tmallnr

// NrtCrmActivityDto 结构体
type NrtCrmActivityDto struct {
	// 头图
	BannerUrl string `json:"banner_url,omitempty" xml:"banner_url,omitempty"`
	// 同城站ID
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 活动状态
	StatusStr string `json:"status_str,omitempty" xml:"status_str,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 活动状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 卖家ID
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 活动ID
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 活动描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 活动规则
	Rule string `json:"rule,omitempty" xml:"rule,omitempty"`
	// 有价券DTO
	SceneActivityList []NrtCrmSceneActivityDto `json:"scene_activity_list,omitempty" xml:"scene_activity_list>nrt_crm_scene_activity_dto,omitempty"`
	// 站外页面ID
	PageId int64 `json:"page_id,omitempty" xml:"page_id,omitempty"`
}
