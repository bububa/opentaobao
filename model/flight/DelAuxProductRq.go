package flight

// DelAuxProductRq 
type DelAuxProductRq struct {

    // 渠道id.校验请求商家身份，channelId非空，则cid必须与之匹配。 当channelId与外部id都为空，表明删除此店铺全部渠道全部辅营产品
    
    ChannelId   int64 `json:"channel_id,omitempty" xml:"channel_id,omitempty"`
    

    // 外部ID列表，外部ID是辅营报价的唯一标识，后续用于校验生单；只允许数字字母组合，最大允许32个字符。 不允许包含空格、换行、|这类特殊符号
    
    OuterIds   []string `json:"outer_ids,omitempty" xml:"outer_ids>string,omitempty"`
    

    // 代理商ID
    
    AgentId   int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
    

    // 接口身份标识用户名（渠道唯一标识）。校验请求商家身份，channelId非空，则cid必须与之匹配。 当channelId与外部id都为空，表明删除此店铺全部渠道全部辅营产品
    
    Cid   string `json:"cid,omitempty" xml:"cid,omitempty"`
    

}
