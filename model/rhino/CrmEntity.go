package rhino

// CrmEntity 结构体
type CrmEntity struct {
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 创建人名称
	Creator string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 实例id
	InstanceId string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
	// 创建人dingId
	CreatorDingUserId string `json:"creator_ding_user_id,omitempty" xml:"creator_ding_user_id,omitempty"`
	// 数据json,修改时包括before data
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 修改人dingId
	ModifierDingUserId string `json:"modifier_ding_user_id,omitempty" xml:"modifier_ding_user_id,omitempty"`
	// 修改人名称
	Modifier string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
}
