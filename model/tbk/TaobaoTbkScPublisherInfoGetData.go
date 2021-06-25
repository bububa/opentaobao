package tbk

// TaobaoTbkScPublisherInfoGetData 
type TaobaoTbkScPublisherInfoGetData struct {

    // 渠道专属pidList
    RootPidChannelList   []String `json:"root_pid_channel_list,omitempty"`

    // 共享字段 - 总记录数
    TotalCount   int64 `json:"total_count,omitempty"`

    // 共享字段 - 渠道或会员列表
    InviterList   []MapData `json:"inviter_list,omitempty"`

}
