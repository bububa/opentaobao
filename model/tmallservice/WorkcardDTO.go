package tmallservice

// WorkcardDTO 
/* model for simplify = false
type WorkcardDTO struct {

    // 额外属性
    
    Attributes   string `json:"attributes,omitempty"`
    

    // 签到时间
    
    GmtSignIn   string `json:"gmt_sign_in,omitempty"`
    

    // 使用次数
    
    UsedServiceCount   int64 `json:"used_service_count,omitempty"`
    

    // 服务订单数据
    
    ServiceTradeOrder  *struct {
        ServiceTradeOrder  *ServiceTradeOrder `json:"service_trade_order,omitempty"`
    } `json:"service_trade_order,omitempty"`
    

    // 预约结束时间
    
    GmtReserveEnd   string `json:"gmt_reserve_end,omitempty"`
    

    // 修改时间
    
    GmtModify   string `json:"gmt_modify,omitempty"`
    

    // 工单服务总次数
    
    ServiceCount   int64 `json:"service_count,omitempty"`
    

    // 取消时间
    
    GmtCancel   string `json:"gmt_cancel,omitempty"`
    

    // 工单状态编码
    
    StatusCode   string `json:"status_code,omitempty"`
    

    // 工单id
    
    Id   int64 `json:"id,omitempty"`
    

    // 预约失败描述
    
    ReserveFailDesc   string `json:"reserve_fail_desc,omitempty"`
    

    // 剩余次数
    
    LeftServiceCount   int64 `json:"left_service_count,omitempty"`
    

    // 主订单号
    
    ParentBizOrderId   int64 `json:"parent_biz_order_id,omitempty"`
    

    // 预约失败编码
    
    ReserveFailCode   int64 `json:"reserve_fail_code,omitempty"`
    

    // 买家信息
    
    Buyer  *struct {
        Buyer  *Buyer `json:"buyer,omitempty"`
    } `json:"buyer,omitempty"`
    

    // 服务提供者
    
    ServiceProvider  *struct {
        ServiceProvider  *ServiceProvider `json:"service_provider,omitempty"`
    } `json:"service_provider,omitempty"`
    

    // 服务履约类型
    
    FulfilTypeCode   string `json:"fulfil_type_code,omitempty"`
    

    // 核销时间
    
    GmtIdentify   string `json:"gmt_identify,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 预约时间开始
    
    GmtReserveStart   string `json:"gmt_reserve_start,omitempty"`
    

    // 实物订单信息
    
    MasterTradeOrder  *struct {
        MasterTradeOrder  *MasterTradeOrder `json:"master_trade_order,omitempty"`
    } `json:"master_trade_order,omitempty"`
    

    // 服务定义
    
    ServiceDefinition  *struct {
        ServiceDefinition  *ServiceDefinition `json:"service_definition,omitempty"`
    } `json:"service_definition,omitempty"`
    

    // 买家期望服务时间
    
    BuyerExpectDate   string `json:"buyer_expect_date,omitempty"`
    

    // 费用信息
    
    FeeList  struct {
        FeeInfo  []FeeInfo `json:"fee_info,omitempty"`
    } `json:"fee_list,omitempty"`
    

    // 核销/工单完结请求中直接带回该字段
    
    Sequence   int64 `json:"sequence,omitempty"`
    

    // 服务单id
    
    SpServiceOrderId   int64 `json:"sp_service_order_id,omitempty"`
    

}
*/

// WorkcardDTO 
type WorkcardDTO struct {

    // 额外属性
    Attributes   string `json:"attributes,omitempty"`

    // 签到时间
    GmtSignIn   string `json:"gmt_sign_in,omitempty"`

    // 使用次数
    UsedServiceCount   int64 `json:"used_service_count,omitempty"`

    // 服务订单数据
    ServiceTradeOrder   *ServiceTradeOrder `json:"service_trade_order,omitempty"`

    // 预约结束时间
    GmtReserveEnd   string `json:"gmt_reserve_end,omitempty"`

    // 修改时间
    GmtModify   string `json:"gmt_modify,omitempty"`

    // 工单服务总次数
    ServiceCount   int64 `json:"service_count,omitempty"`

    // 取消时间
    GmtCancel   string `json:"gmt_cancel,omitempty"`

    // 工单状态编码
    StatusCode   string `json:"status_code,omitempty"`

    // 工单id
    Id   int64 `json:"id,omitempty"`

    // 预约失败描述
    ReserveFailDesc   string `json:"reserve_fail_desc,omitempty"`

    // 剩余次数
    LeftServiceCount   int64 `json:"left_service_count,omitempty"`

    // 主订单号
    ParentBizOrderId   int64 `json:"parent_biz_order_id,omitempty"`

    // 预约失败编码
    ReserveFailCode   int64 `json:"reserve_fail_code,omitempty"`

    // 买家信息
    Buyer   *Buyer `json:"buyer,omitempty"`

    // 服务提供者
    ServiceProvider   *ServiceProvider `json:"service_provider,omitempty"`

    // 服务履约类型
    FulfilTypeCode   string `json:"fulfil_type_code,omitempty"`

    // 核销时间
    GmtIdentify   string `json:"gmt_identify,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 预约时间开始
    GmtReserveStart   string `json:"gmt_reserve_start,omitempty"`

    // 实物订单信息
    MasterTradeOrder   *MasterTradeOrder `json:"master_trade_order,omitempty"`

    // 服务定义
    ServiceDefinition   *ServiceDefinition `json:"service_definition,omitempty"`

    // 买家期望服务时间
    BuyerExpectDate   string `json:"buyer_expect_date,omitempty"`

    // 费用信息
    FeeList   []FeeInfo `json:"fee_list,omitempty"`

    // 核销/工单完结请求中直接带回该字段
    Sequence   int64 `json:"sequence,omitempty"`

    // 服务单id
    SpServiceOrderId   int64 `json:"sp_service_order_id,omitempty"`

}
