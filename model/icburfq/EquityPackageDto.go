package icburfq

// EquityPackageDto 
type EquityPackageDto struct {

    // 剩余权益
    
    EquityCount   int64 `json:"equity_count,omitempty" xml:"equity_count,omitempty"`
    

    // 过期时间
    
    ExpiredDate   string `json:"expired_date,omitempty" xml:"expired_date,omitempty"`
    

    // 市场表现分
    
    Score   int64 `json:"score,omitempty" xml:"score,omitempty"`
    

    // 击败其他供应商百分比
    
    BeatSupplierPercent   string `json:"beat_supplier_percent,omitempty" xml:"beat_supplier_percent,omitempty"`
    

    // 市场表现分统计开始时间
    
    StatisticStartDate   string `json:"statistic_start_date,omitempty" xml:"statistic_start_date,omitempty"`
    

    // 市场表现分统计结束时间
    
    StatisticEndDate   string `json:"statistic_end_date,omitempty" xml:"statistic_end_date,omitempty"`
    

    // 剩余置顶报价权益
    
    TopServiceCount   int64 `json:"top_service_count,omitempty" xml:"top_service_count,omitempty"`
    

}
