package usergrowth2

// OfflineConvertionSummaryT1Dto 结构体
type OfflineConvertionSummaryT1Dto struct {
	// 结算新登数
	RegisterUserCntT1 int64 `json:"register_user_cnt_t1,omitempty" xml:"register_user_cnt_t1,omitempty"`
	// 购买首购数
	PurchaseUserCnt int64 `json:"purchase_user_cnt,omitempty" xml:"purchase_user_cnt,omitempty"`
	// 任务名称
	TaskName string `json:"task_name,omitempty" xml:"task_name,omitempty"`
	// 渠道名称
	Company string `json:"company,omitempty" xml:"company,omitempty"`
	// 归因新登数
	RegisterUserCnt int64 `json:"register_user_cnt,omitempty" xml:"register_user_cnt,omitempty"`
	// 结算首购数
	PurchaseUserCntT1 int64 `json:"purchase_user_cnt_t1,omitempty" xml:"purchase_user_cnt_t1,omitempty"`
	// 渠道id
	ChannelId string `json:"channel_id,omitempty" xml:"channel_id,omitempty"`
	// 任务id
	TaskId string `json:"task_id,omitempty" xml:"task_id,omitempty"`
	// 日期
	Ds int64 `json:"ds,omitempty" xml:"ds,omitempty"`
	// 推广码code
	PromoterCode string `json:"promoter_code,omitempty" xml:"promoter_code,omitempty"`
}
