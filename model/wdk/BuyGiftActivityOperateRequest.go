package wdk

// BuyGiftActivityOperateRequest 结构体
type BuyGiftActivityOperateRequest struct {
	// 活动终端：1-APP
	Terminals []string `json:"terminals,omitempty" xml:"terminals>string,omitempty"`
	// 活动生效的经营店ID
	StoreIds []string `json:"store_ids,omitempty" xml:"store_ids>string,omitempty"`
	// 活动人群编码，NEW_USER：新用户，OLD_USER：老用户，LIGHT_NEW_USER：闪购新客，MERCHANT_NEW_USER：商家新用户，MERCHANT_OLD_USER：商家老用户
	MemberCrowdCode []string `json:"member_crowd_code,omitempty" xml:"member_crowd_code>string,omitempty"`
	// 枚举： 2 美团 3 饿了么 26 京东到家 31 翱象淘鲜达 32 翱象共享库存
	Channels []string `json:"channels,omitempty" xml:"channels>string,omitempty"`
	// 活动名称，最长20个字符
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 活动创建者ID（数字类型）
	CreatorId string `json:"creator_id,omitempty" xml:"creator_id,omitempty"`
	// 活动创建者name
	CreatorName string `json:"creator_name,omitempty" xml:"creator_name,omitempty"`
	// 活动描述，最长15个字符
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 外部活动ID（商家自定义）
	OutActId string `json:"out_act_id,omitempty" xml:"out_act_id,omitempty"`
	// 扩展信息，json串
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 活动开始时间（毫秒时间戳）
	StartTime int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 活动结束时间（毫秒时间戳）
	EndTime int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 活动ID
	ActId int64 `json:"act_id,omitempty" xml:"act_id,omitempty"`
}
