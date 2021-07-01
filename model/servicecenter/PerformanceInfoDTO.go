package servicecenter

// PerformanceInfoDto 
type PerformanceInfoDto struct {
    // 首次相应时间
    FirstResponseTime   string `json:"first_response_time,omitempty" xml:"first_response_time,omitempty"`
    // 平均最后在线时间
    AvgLastOnlineTime   string `json:"avg_last_online_time,omitempty" xml:"avg_last_online_time,omitempty"`
    // 平均日在线时长
    AvgOnlineTimePerDay   string `json:"avg_online_time_per_day,omitempty" xml:"avg_online_time_per_day,omitempty"`
    // 客服转化率
    ConversionRateOfSer   string `json:"conversion_rate_of_ser,omitempty" xml:"conversion_rate_of_ser,omitempty"`
    // 销售量
    SalesQuantity   string `json:"sales_quantity,omitempty" xml:"sales_quantity,omitempty"`
    // 销售额
    SalesAmount   string `json:"sales_amount,omitempty" xml:"sales_amount,omitempty"`
    // 在线总时长
    TotalOnlineTime   string `json:"total_online_time,omitempty" xml:"total_online_time,omitempty"`
    // 上班天数
    OnlineDays   string `json:"online_days,omitempty" xml:"online_days,omitempty"`
    // 客服客单价
    UnitPriceOfSer   string `json:"unit_price_of_ser,omitempty" xml:"unit_price_of_ser,omitempty"`
    // 回复率
    ResponseRate   string `json:"response_rate,omitempty" xml:"response_rate,omitempty"`
    // 子账号名字
    SubAccountName   string `json:"sub_account_name,omitempty" xml:"sub_account_name,omitempty"`
    // 平均上线时间
    AvgOnlineTime   string `json:"avg_online_time,omitempty" xml:"avg_online_time,omitempty"`
    // 销售提成
    SalesBonus   string `json:"sales_bonus,omitempty" xml:"sales_bonus,omitempty"`
    // 平均响应时间
    AvgResponseTime   string `json:"avg_response_time,omitempty" xml:"avg_response_time,omitempty"`
}
