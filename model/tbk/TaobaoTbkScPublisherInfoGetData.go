package tbk

// TaobaoTbkScPublisherInfoGetData 结构体
type TaobaoTbkScPublisherInfoGetData struct {
	// 渠道专属pidList
	RootPidChannelList []string `json:"root_pid_channel_list,omitempty" xml:"root_pid_channel_list>string,omitempty"`
	// 共享字段 - 渠道或会员列表
	InviterList []TaobaoTbkScPublisherInfoGetMapData `json:"inviter_list,omitempty" xml:"inviter_list>taobao_tbk_sc_publisher_info_get_map_data,omitempty"`
	// 共享字段 - 总记录数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
