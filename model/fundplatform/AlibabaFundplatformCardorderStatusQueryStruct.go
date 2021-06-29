package fundplatform

// AlibabaFundplatformCardorderStatusQueryStruct 
type AlibabaFundplatformCardorderStatusQueryStruct struct {

    // 子制卡单ID
    
    CardOrderId   int64 `json:"card_order_id,omitempty" xml:"card_order_id,omitempty"`
    

    // 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
    
    OwnSign   string `json:"own_sign,omitempty" xml:"own_sign,omitempty"`
    

    // 制卡当前阶段 该字段是枚举值，有： REQUEST_RECEIVED（请求接收成功），REQUEST_NOT_EXIST（请求不存在），WAITING_DELIVERY（制卡完成，等待发货），DELIVERED（已发货）
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

    // 错误详情
    
    ResultMessage   string `json:"result_message,omitempty" xml:"result_message,omitempty"`
    

    // 错误CODE
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    

    // 是否调用成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 物流商名称，仅当制卡状态为发货完成时返回该值
    
    LogisticsCompany   string `json:"logistics_company,omitempty" xml:"logistics_company,omitempty"`
    

    // 物流单号，仅当制卡状态为发货完成时返回该值
    
    LogisticsOrderId   string `json:"logistics_order_id,omitempty" xml:"logistics_order_id,omitempty"`
    

}
