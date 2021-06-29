package westcrm

// ThirdpartyPageData 
type ThirdpartyPageData struct {
    // 当前页
    Current   int64 `json:"current,omitempty" xml:"current,omitempty"`
    // 活动列表
    List   []ActivitiesListVo `json:"list,omitempty" xml:"list>activities_list_vo,omitempty"`
    // 总数
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
}
