package tmallservice

// SpServiceorderDTO 
/* model for simplify = false
type SpServiceorderDTO struct {

    // 取消的份数
    
    RefundServiceCount   int64 `json:"refund_service_count,omitempty"`
    

    // 实物子订单信息
    
    MasterTradeOrder  *struct {
        MasterTradeOrderDTO  *MasterTradeOrderDTO `json:"master_trade_order_dto,omitempty"`
    } `json:"master_trade_order,omitempty"`
    

    // 服务定义
    
    ServiceDefinition  *struct {
        ServiceDefinitionDTO  *ServiceDefinitionDTO `json:"service_definition_dto,omitempty"`
    } `json:"service_definition,omitempty"`
    

    // 买家信息
    
    Buyer  *struct {
        BuyerDTO  *BuyerDTO `json:"buyer_dto,omitempty"`
    } `json:"buyer,omitempty"`
    

    // 已使用份数
    
    UsedServiceCount   int64 `json:"used_service_count,omitempty"`
    

    // 费用信息
    
    FeeList  struct {
        FeeInfo  []FeeInfo `json:"fee_info,omitempty"`
    } `json:"fee_list,omitempty"`
    

    // 剩余的份数
    
    LeftServiceCount   int64 `json:"left_service_count,omitempty"`
    

    // 服务子订单信息
    
    ServiceTradeOrder  *struct {
        ServiceTradeOrderDTO  *ServiceTradeOrderDTO `json:"service_trade_order_dto,omitempty"`
    } `json:"service_trade_order,omitempty"`
    

    // 服务的总份数
    
    ServiceCount   int64 `json:"service_count,omitempty"`
    

    // 服务过期时间
    
    GmtExpire   string `json:"gmt_expire,omitempty"`
    

    // 服务单修改时间
    
    GmtModified   string `json:"gmt_modified,omitempty"`
    

    // 服务单创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 服务单有效期开始时间
    
    GmtEffect   string `json:"gmt_effect,omitempty"`
    

    // 服务单id
    
    SpServiceorderId   int64 `json:"sp_serviceorder_id,omitempty"`
    

    // 正在使用中的次数
    
    UsingServiceCount   int64 `json:"using_service_count,omitempty"`
    

    // 状态编码：create(创建)、effect(生效)、closed(关闭)、finish(完成)
    
    StatusCode   string `json:"status_code,omitempty"`
    

}
*/

// SpServiceorderDTO 
type SpServiceorderDTO struct {

    // 取消的份数
    RefundServiceCount   int64 `json:"refund_service_count,omitempty"`

    // 实物子订单信息
    MasterTradeOrder   *MasterTradeOrderDTO `json:"master_trade_order,omitempty"`

    // 服务定义
    ServiceDefinition   *ServiceDefinitionDTO `json:"service_definition,omitempty"`

    // 买家信息
    Buyer   *BuyerDTO `json:"buyer,omitempty"`

    // 已使用份数
    UsedServiceCount   int64 `json:"used_service_count,omitempty"`

    // 费用信息
    FeeList   []FeeInfo `json:"fee_list,omitempty"`

    // 剩余的份数
    LeftServiceCount   int64 `json:"left_service_count,omitempty"`

    // 服务子订单信息
    ServiceTradeOrder   *ServiceTradeOrderDTO `json:"service_trade_order,omitempty"`

    // 服务的总份数
    ServiceCount   int64 `json:"service_count,omitempty"`

    // 服务过期时间
    GmtExpire   string `json:"gmt_expire,omitempty"`

    // 服务单修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 服务单创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 服务单有效期开始时间
    GmtEffect   string `json:"gmt_effect,omitempty"`

    // 服务单id
    SpServiceorderId   int64 `json:"sp_serviceorder_id,omitempty"`

    // 正在使用中的次数
    UsingServiceCount   int64 `json:"using_service_count,omitempty"`

    // 状态编码：create(创建)、effect(生效)、closed(关闭)、finish(完成)
    StatusCode   string `json:"status_code,omitempty"`

}
