package axindata

// FscRouteScheduleDetailApiDto 结构体
type FscRouteScheduleDetailApiDto struct {
	// 日程明细类型 SCENIC景点;HOTEL酒店;TRAFFIC交通;BREAKFAST早餐;LUNCH午餐;DINNER晚餐;SHOPPING购物;ACTIVITY活动;OTHER其他;TRIP_DESCRIPTION行程描述
	ScheduleItem string `json:"schedule_item,omitempty" xml:"schedule_item,omitempty"`
	// 日程明细详情说明
	ScheduleDesc string `json:"schedule_desc,omitempty" xml:"schedule_desc,omitempty"`
	// 景点poiId
	PoiId string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	// 类型名称
	ScheduleName string `json:"schedule_name,omitempty" xml:"schedule_name,omitempty"`
	// 日程明细排序
	Sort int64 `json:"sort,omitempty" xml:"sort,omitempty"`
}
