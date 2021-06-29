package cainiaocntec

// ActivityStatisticsDto 
type ActivityStatisticsDto struct {

    // 箱规
    
    InventoryCoefficient   string `json:"inventory_coefficient,omitempty" xml:"inventory_coefficient,omitempty"`
    

    // 总价
    
    TotoalPrice   string `json:"totoal_price,omitempty" xml:"totoal_price,omitempty"`
    

    // 购买数量
    
    BuyAmount   string `json:"buy_amount,omitempty" xml:"buy_amount,omitempty"`
    

    // 单价
    
    Price   string `json:"price,omitempty" xml:"price,omitempty"`
    

    // 商品标题
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 商品编码
    
    ProductCode   string `json:"product_code,omitempty" xml:"product_code,omitempty"`
    

    // 驿站名称
    
    StationName   string `json:"station_name,omitempty" xml:"station_name,omitempty"`
    

    // 驿站id
    
    StationId   string `json:"station_id,omitempty" xml:"station_id,omitempty"`
    

    // 门店名称
    
    StoreName   string `json:"store_name,omitempty" xml:"store_name,omitempty"`
    

    // 门店id
    
    StoreId   string `json:"store_id,omitempty" xml:"store_id,omitempty"`
    

    // 活动名称
    
    ActivityName   string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
    

    // 预计到货时间
    
    PredictArrivalTime   string `json:"predict_arrival_time,omitempty" xml:"predict_arrival_time,omitempty"`
    

    // 营销活动
    
    MktActivityType   string `json:"mkt_activity_type,omitempty" xml:"mkt_activity_type,omitempty"`
    

}
