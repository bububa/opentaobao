package usergrowth2

// TaobaoUsergrowthOfflineConvertionDetailsEightGetE 结构体
type TaobaoUsergrowthOfflineConvertionDetailsEightGetE struct {
	// 结算系数
	AccountRatio string `json:"account_ratio,omitempty" xml:"account_ratio,omitempty"`
	// 渠道名称
	ChannelName string `json:"channel_name,omitempty" xml:"channel_name,omitempty"`
	// 任务名称
	TaskName string `json:"task_name,omitempty" xml:"task_name,omitempty"`
	// 渠道id
	ChannelId string `json:"channel_id,omitempty" xml:"channel_id,omitempty"`
	// 任务ID
	TaskId string `json:"task_id,omitempty" xml:"task_id,omitempty"`
	// 结算金额
	SettlementAccount string `json:"settlement_account,omitempty" xml:"settlement_account,omitempty"`
	// 推广码code
	PromoterCode string `json:"promoter_code,omitempty" xml:"promoter_code,omitempty"`
	// 归因首购数
	PurchaseUserCnt int64 `json:"purchase_user_cnt,omitempty" xml:"purchase_user_cnt,omitempty"`
	// 结算新登数
	AccountRegisterUserCnt int64 `json:"account_register_user_cnt,omitempty" xml:"account_register_user_cnt,omitempty"`
	// 归因新登数
	RegisterUserCnt int64 `json:"register_user_cnt,omitempty" xml:"register_user_cnt,omitempty"`
	// 日期
	Ds int64 `json:"ds,omitempty" xml:"ds,omitempty"`
	// 结算首购数
	AccountPurchaseUserCnt int64 `json:"account_purchase_user_cnt,omitempty" xml:"account_purchase_user_cnt,omitempty"`
}
