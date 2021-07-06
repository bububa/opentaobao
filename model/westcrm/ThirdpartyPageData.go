package westcrm

// ThirdpartyPageData 结构体
type ThirdpartyPageData struct {
	// 活动列表
	List []ActivitiesListVo `json:"list,omitempty" xml:"list>activities_list_vo,omitempty"`
	// 当前页
	Current int64 `json:"current,omitempty" xml:"current,omitempty"`
	// 总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}
