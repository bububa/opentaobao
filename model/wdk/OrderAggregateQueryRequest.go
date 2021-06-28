package wdk

// OrderAggregateQueryRequest 
type OrderAggregateQueryRequest struct {

    // 起始时间
    
    StartTime   string `json:"start_time,omitempty" xml:"start_time,omitempty"`
    

    // 结束时间
    
    EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
    

    // DRF1001
    
    StoreId   string `json:"store_id,omitempty" xml:"store_id,omitempty"`
    

    // 机台号
    
    PosNo   string `json:"pos_no,omitempty" xml:"pos_no,omitempty"`
    

    // 下单终端: APP / POS
    
    OrderClient   string `json:"order_client,omitempty" xml:"order_client,omitempty"`
    

    // 班次号
    
    DutyCode   string `json:"duty_code,omitempty" xml:"duty_code,omitempty"`
    

    // 收营员id
    
    OperatorId   string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
    

    // 订单状态: PAID / PACAKAGED / SUCCESS
    
    OrderStatus   string `json:"order_status,omitempty" xml:"order_status,omitempty"`
    

    // 分页序号
    
    PageIndex   int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
    

    // 分页size
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

}
