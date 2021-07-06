package wdk

// ItemPoolActivityOperateRequest 结构体
type ItemPoolActivityOperateRequest struct {
	// 阶梯分组
	StairGroups []StairGroupDto `json:"stair_groups,omitempty" xml:"stair_groups>stair_group_dto,omitempty"`
	// 支持的终端，1：APP,2:POS
	Terminals []int64 `json:"terminals,omitempty" xml:"terminals>int64,omitempty"`
	// 经营店
	StoreIds []string `json:"store_ids,omitempty" xml:"store_ids>string,omitempty"`
	// 限制人群,OLD_USER:老用户，NEW_USER:新用户，LIGHT_NEW_USER：闪购新客
	MemberCrowdCode []string `json:"member_crowd_code,omitempty" xml:"member_crowd_code>string,omitempty"`
	// 逻辑分组
	LogicGroups []LogicGroupDto `json:"logic_groups,omitempty" xml:"logic_groups>logic_group_dto,omitempty"`
	// 活动名称，最长15个字符
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 操作人id
	CreatorId string `json:"creator_id,omitempty" xml:"creator_id,omitempty"`
	// 操作人名称
	CreatorName string `json:"creator_name,omitempty" xml:"creator_name,omitempty"`
	// 活动描述，最长30个字符
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 外部商家erp活动id
	OutActId string `json:"out_act_id,omitempty" xml:"out_act_id,omitempty"`
	// 周期配置
	PeriodicConfig *PeriodicConfigDto `json:"periodic_config,omitempty" xml:"periodic_config,omitempty"`
	// 限购
	Limit *LimitDto `json:"limit,omitempty" xml:"limit,omitempty"`
	// 活动开始时间戳
	StartTime int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 活动结束时间戳
	EndTime int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 同城零售活动id
	ActId int64 `json:"act_id,omitempty" xml:"act_id,omitempty"`
	// 上不封顶标识，若设置{@link #enableMultiple}=true，则不支持多阶梯配置，且可以叠加优惠
	EnableMultiple bool `json:"enable_multiple,omitempty" xml:"enable_multiple,omitempty"`
}
