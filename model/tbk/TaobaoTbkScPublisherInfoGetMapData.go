package tbk

// TaobaoTbkScPublisherInfoGetMapData 结构体
type TaobaoTbkScPublisherInfoGetMapData struct {
	// 共享字段 - 备案场景：common（通用备案），etao（一淘备案），minietao（一淘小程序备案），offlineShop（线下门店备案），offlinePerson（线下个人备案）
	RelationApp string `json:"relation_app,omitempty" xml:"relation_app,omitempty"`
	// 共享字段 - 备案日期
	CreateDate string `json:"create_date,omitempty" xml:"create_date,omitempty"`
	// 渠道独有 - 渠道昵称
	AccountName string `json:"account_name,omitempty" xml:"account_name,omitempty"`
	// 渠道独有 - 渠道姓名
	RealName string `json:"real_name,omitempty" xml:"real_name,omitempty"`
	// 渠道独有 - 渠道关系ID
	RelationId int64 `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
	// 渠道独有 - 线下场景信息，1 - 门店，2- 学校，3 - 工厂，4 - 其他
	OfflineScene string `json:"offline_scene,omitempty" xml:"offline_scene,omitempty"`
	// 渠道独有 - 线上场景信息，1 - 微信群，2- QQ群，3 - 其他
	OnlineScene string `json:"online_scene,omitempty" xml:"online_scene,omitempty"`
	// 渠道独有 - 媒体侧渠道备注信息
	Note string `json:"note,omitempty" xml:"note,omitempty"`
	// 共享字段 - 渠道/会员专属pid
	RootPid string `json:"root_pid,omitempty" xml:"root_pid,omitempty"`
	// 共享字段 - 渠道/会员原始身份信息
	Rtag string `json:"rtag,omitempty" xml:"rtag,omitempty"`
	// 线下备案专属信息
	OfflineInfo *RegisterInfoDto `json:"offline_info,omitempty" xml:"offline_info,omitempty"`
	// 会员独有 - 会员运营ID
	SpecialId int64 `json:"special_id,omitempty" xml:"special_id,omitempty"`
	// 渠道独有 - 处罚状态
	PunishStatus string `json:"punish_status,omitempty" xml:"punish_status,omitempty"`
	// 淘宝客外部用户标记，如自身系统账户ID；微信ID等
	ExternalId string `json:"external_id,omitempty" xml:"external_id,omitempty"`
	// 1-微信、2-微博、3-抖音、4-快手、5-QQ，0-其他
	ExternalType string `json:"external_type,omitempty" xml:"external_type,omitempty"`
}
