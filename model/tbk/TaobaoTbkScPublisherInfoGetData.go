package tbk

// TaobaoTbkScPublisherInfoGetData 
/* model for simplify = false
type TaobaoTbkScPublisherInfoGetData struct {

    // 渠道专属pidList
    
    RootPidChannelList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"root_pid_channel_list,omitempty"`
    

    // 共享字段 - 总记录数
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

    // 共享字段 - 渠道或会员列表
    
    InviterList  struct {
        TaobaoTbkScPublisherInfoGetMapData  []TaobaoTbkScPublisherInfoGetMapData `json:"taobao_tbk_sc_publisher_info_get_map_data,omitempty"`
    } `json:"inviter_list,omitempty"`
    

}
*/

// TaobaoTbkScPublisherInfoGetData 
type TaobaoTbkScPublisherInfoGetData struct {

    // 渠道专属pidList
    RootPidChannelList   []string `json:"root_pid_channel_list,omitempty"`

    // 共享字段 - 总记录数
    TotalCount   int64 `json:"total_count,omitempty"`

    // 共享字段 - 渠道或会员列表
    InviterList   []TaobaoTbkScPublisherInfoGetMapData `json:"inviter_list,omitempty"`

}
