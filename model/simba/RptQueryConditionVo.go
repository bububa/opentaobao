package simba

// RptQueryConditionVo 结构体
type RptQueryConditionVo struct {
	// 流量来源(资源位包id),可为空
	AdzonePkgIdList []string `json:"adzone_pkg_id_list,omitempty" xml:"adzone_pkg_id_list>string,omitempty"`
	// 归因口径
	UnifyType string `json:"unify_type,omitempty" xml:"unify_type,omitempty"`
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 是否查询实时报表，true查实时、false查离线，不传则默认查离线
	IsRt bool `json:"is_rt,omitempty" xml:"is_rt,omitempty"`
}
