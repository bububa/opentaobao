package alicom

// PreAuthorizeModel 
/* model for simplify = false
type PreAuthorizeModel struct {

    // 扩展字段(json)
    
    Ext   string `json:"ext,omitempty"`
    

    // 0:业务办理成功，2:业务办理失败,全额解冻
    
    Status   int64 `json:"status,omitempty"`
    

    // 淘宝订单号
    
    TbOrderNo   string `json:"tb_order_no,omitempty"`
    

    // 外部流水号
    
    OutBizOrderNo   string `json:"out_biz_order_no,omitempty"`
    

    // 授权水号
    
    FundAuthNo   string `json:"fund_auth_no,omitempty"`
    

    // 描述
    
    Desc   string `json:"desc,omitempty"`
    

    // 履约日期(yyyy-MM-dd)
    
    FulfillmentDate   string `json:"fulfillment_date,omitempty"`
    

    // 0:解冻，1:违约 2：罚没
    
    OperationType   string `json:"operation_type,omitempty"`
    

    // 淘宝昵称
    
    UserNick   string `json:"user_nick,omitempty"`
    

    // 金额（单位：分）
    
    Amount   string `json:"amount,omitempty"`
    

    // 淘宝订单号
    
    Tid   string `json:"tid,omitempty"`
    

}
*/

// PreAuthorizeModel 
type PreAuthorizeModel struct {

    // 扩展字段(json)
    Ext   string `json:"ext,omitempty"`

    // 0:业务办理成功，2:业务办理失败,全额解冻
    Status   int64 `json:"status,omitempty"`

    // 淘宝订单号
    TbOrderNo   string `json:"tb_order_no,omitempty"`

    // 外部流水号
    OutBizOrderNo   string `json:"out_biz_order_no,omitempty"`

    // 授权水号
    FundAuthNo   string `json:"fund_auth_no,omitempty"`

    // 描述
    Desc   string `json:"desc,omitempty"`

    // 履约日期(yyyy-MM-dd)
    FulfillmentDate   string `json:"fulfillment_date,omitempty"`

    // 0:解冻，1:违约 2：罚没
    OperationType   string `json:"operation_type,omitempty"`

    // 淘宝昵称
    UserNick   string `json:"user_nick,omitempty"`

    // 金额（单位：分）
    Amount   string `json:"amount,omitempty"`

    // 淘宝订单号
    Tid   string `json:"tid,omitempty"`

}
