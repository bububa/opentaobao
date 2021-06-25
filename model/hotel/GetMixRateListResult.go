package hotel

// GetMixRateListResult 
type GetMixRateListResult struct {

    // 附加属性
    Attributes   *Attributes `json:"attributes,omitempty"`

    // 页面布局信息
    ConfigInfo   string `json:"config_info,omitempty"`

    // debugInfo
    DebugInfo   string `json:"debug_info,omitempty"`

    // degrade
    Degrade   bool `json:"degrade,omitempty"`

    // errCode
    ErrCode   string `json:"err_code,omitempty"`

    // errMsg
    ErrMsg   string `json:"err_msg,omitempty"`

    // 是否可以向下翻页，0不可以，1可以
    HasNextPage   int64 `json:"has_next_page,omitempty"`

    // 服务器主机名称，排查问题使用
    HostName   string `json:"host_name,omitempty"`

    // 统计项
    ItemStatistic   *ItemStatisticVo `json:"item_statistic,omitempty"`

    // 商品评论列表
    MixRates   []MixRateVo `json:"mix_rates,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

    // 总记录数
    TotalNum   int64 `json:"total_num,omitempty"`

}
