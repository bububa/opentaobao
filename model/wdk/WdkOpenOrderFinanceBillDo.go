package wdk

// WdkOpenOrderFinanceBillDo 
/* model for simplify = false
type WdkOpenOrderFinanceBillDo struct {

    // alipay：支付宝
    
    PayChannel   string `json:"pay_channel,omitempty"`
    

    // 盒马订单号
    
    HmOrderId   string `json:"hm_order_id,omitempty"`
    

    // app：线上， pos：线下
    
    OrderChannel   string `json:"order_channel,omitempty"`
    

    // 经营店id
    
    StoreId   string `json:"store_id,omitempty"`
    

    // 商家编码
    
    MerchantCode   string `json:"merchant_code,omitempty"`
    

    // 淘系订单号
    
    TpOrderId   string `json:"tp_order_id,omitempty"`
    

    // 账单日期
    
    Dt   string `json:"dt,omitempty"`
    

    // 支付宝二级商户ID
    
    Smid   string `json:"smid,omitempty"`
    

    // 金额，正向交易时数值是正数，退款时数值是负数
    
    Amount   int64 `json:"amount,omitempty"`
    

    // 支付时间
    
    PayTime   string `json:"pay_time,omitempty"`
    

    // 业务类型。PAY：支付， REFUND：退款
    
    BizType   string `json:"biz_type,omitempty"`
    

    // 支付机构支付流水号
    
    TradeNo   string `json:"trade_no,omitempty"`
    

    // 财务组织名称
    
    FinanceOrganizationName   string `json:"finance_organization_name,omitempty"`
    

    // 财务组织code
    
    FinanceOrganizationCode   string `json:"finance_organization_code,omitempty"`
    

}
*/

// WdkOpenOrderFinanceBillDo 
type WdkOpenOrderFinanceBillDo struct {

    // alipay：支付宝
    PayChannel   string `json:"pay_channel,omitempty"`

    // 盒马订单号
    HmOrderId   string `json:"hm_order_id,omitempty"`

    // app：线上， pos：线下
    OrderChannel   string `json:"order_channel,omitempty"`

    // 经营店id
    StoreId   string `json:"store_id,omitempty"`

    // 商家编码
    MerchantCode   string `json:"merchant_code,omitempty"`

    // 淘系订单号
    TpOrderId   string `json:"tp_order_id,omitempty"`

    // 账单日期
    Dt   string `json:"dt,omitempty"`

    // 支付宝二级商户ID
    Smid   string `json:"smid,omitempty"`

    // 金额，正向交易时数值是正数，退款时数值是负数
    Amount   int64 `json:"amount,omitempty"`

    // 支付时间
    PayTime   string `json:"pay_time,omitempty"`

    // 业务类型。PAY：支付， REFUND：退款
    BizType   string `json:"biz_type,omitempty"`

    // 支付机构支付流水号
    TradeNo   string `json:"trade_no,omitempty"`

    // 财务组织名称
    FinanceOrganizationName   string `json:"finance_organization_name,omitempty"`

    // 财务组织code
    FinanceOrganizationCode   string `json:"finance_organization_code,omitempty"`

}
