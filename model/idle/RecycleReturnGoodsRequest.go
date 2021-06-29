package idle

// RecycleReturnGoodsRequest 
type RecycleReturnGoodsRequest struct {

    // 快递公司
    
    LogisticsCompanyName   string `json:"logistics_company_name,omitempty" xml:"logistics_company_name,omitempty"`
    

    // 运单号
    
    LogisticsMailNo   string `json:"logistics_mail_no,omitempty" xml:"logistics_mail_no,omitempty"`
    

    // 订单号
    
    BizOrderId   int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
    

    // 卖家电话
    
    MobileNumber   string `json:"mobile_number,omitempty" xml:"mobile_number,omitempty"`
    

}
