package fundplatform

// AlibabaFundplatformCardorderMakeSuccessStruct 
type AlibabaFundplatformCardorderMakeSuccessStruct struct {

    // 制卡单号
    
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    

    // 制卡时传入的外部订单号
    
    OutBizId   string `json:"out_biz_id,omitempty" xml:"out_biz_id,omitempty"`
    

    // 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
    
    OwnSign   string `json:"own_sign,omitempty" xml:"own_sign,omitempty"`
    

}
