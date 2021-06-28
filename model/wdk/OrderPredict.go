package wdk

// OrderPredict 
type OrderPredict struct {

    // 实际总单量
    
    ActualOrderQuantity   int64 `json:"actual_order_quantity,omitempty" xml:"actual_order_quantity,omitempty"`
    

    // 分渠道实际单量，JSON结构
    
    ChannelActualQuantity   string `json:"channel_actual_quantity,omitempty" xml:"channel_actual_quantity,omitempty"`
    

    // 分渠道预测单量，JSON结构
    
    ChannelPredictQuantity   string `json:"channel_predict_quantity,omitempty" xml:"channel_predict_quantity,omitempty"`
    

    // 配送站编码
    
    DeliveryStationCode   string `json:"delivery_station_code,omitempty" xml:"delivery_station_code,omitempty"`
    

    // 配送站名称
    
    DeliveryStationName   string `json:"delivery_station_name,omitempty" xml:"delivery_station_name,omitempty"`
    

    // 预测日期
    
    PredictDate   string `json:"predict_date,omitempty" xml:"predict_date,omitempty"`
    

    // 预测总单量
    
    PredictOrderQuantity   int64 `json:"predict_order_quantity,omitempty" xml:"predict_order_quantity,omitempty"`
    

    // 子公司编码
    
    SubCompanyCode   string `json:"sub_company_code,omitempty" xml:"sub_company_code,omitempty"`
    

    // 子公司名称
    
    SubCompanyName   string `json:"sub_company_name,omitempty" xml:"sub_company_name,omitempty"`
    

    // 时间段
    
    TimeRange   string `json:"time_range,omitempty" xml:"time_range,omitempty"`
    

    // 门店编码
    
    WarehouseCode   string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
    

    // 门店名称
    
    WarehouseName   string `json:"warehouse_name,omitempty" xml:"warehouse_name,omitempty"`
    

}
