package maitix

// PerformInfoDto 结构体
type PerformInfoDto struct {
	// 价格
	PriceInfoList []PriceInfoDto `json:"price_info_list,omitempty" xml:"price_info_list>price_info_dto,omitempty"`
	// 子场次列表-暂时没用
	SubPerformList []string `json:"sub_perform_list,omitempty" xml:"sub_perform_list>string,omitempty"`
	// 场次有效期规则 每周周几可以
	WeeklyList []int64 `json:"weekly_list,omitempty" xml:"weekly_list>int64,omitempty"`
	// 延改期原因
	ChangeReason string `json:"change_reason,omitempty" xml:"change_reason,omitempty"`
	// 演出结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 场次名称
	PerformName string `json:"perform_name,omitempty" xml:"perform_name,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 演出开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 是否修改了场次时间 0或null：否 1：是
	IsChangePerformTime int64 `json:"is_change_perform_time,omitempty" xml:"is_change_perform_time,omitempty"`
	// 场次ID
	PerformId int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
	// 场次设置
	PerformSetting *PerformSettingDto `json:"perform_setting,omitempty" xml:"perform_setting,omitempty"`
	// 场次状态(0：创建中 10：已创建 20：待销售 30：销售中 40：场次取消 50：场次结束)一般不会透出30之前的状态给渠道
	PerformStatus int64 `json:"perform_status,omitempty" xml:"perform_status,omitempty"`
	// 场次类型(1 单场次，2 多次通票，3 单次通票)-暂时没用,可以认为都是普通场次
	PerformType int64 `json:"perform_type,omitempty" xml:"perform_type,omitempty"`
	// 是否对号入座 0：不对号入座 1：对号入座 2：对区入座
	ReserveSeat int64 `json:"reserve_seat,omitempty" xml:"reserve_seat,omitempty"`
}
