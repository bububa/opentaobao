package tmallservice

// SpServiceOrderDTO 
type SpServiceOrderDTO struct {

    // 过期日期
    GmtExpire   string `json:"gmt_expire,omitempty"`

    // 服务定义
    ServiceDefinitionDTO   *ServiceDefinitionDTO `json:"service_definition_d_t_o,omitempty"`

    // 生效日期
    GmtEffect   string `json:"gmt_effect,omitempty"`

    // 服务单id
    Id   int64 `json:"id,omitempty"`

    // 服务交易订单
    ServiceTradeOrderDTO   *ServiceTradeOrderDTO `json:"service_trade_order_d_t_o,omitempty"`

    // 服务单申请工单的幂等键
    ServiceSequence   int64 `json:"service_sequence,omitempty"`

    // 履约类型:1, "到店"2, "到家"3, "寄送"
    FulfilTypeCode   string `json:"fulfil_type_code,omitempty"`

    // 服务提供者
    ServiceProviderDTO   *ServiceProviderDto `json:"service_provider_d_t_o,omitempty"`

}
