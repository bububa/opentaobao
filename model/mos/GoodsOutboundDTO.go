package mos

// GoodsOutboundDTO 
type GoodsOutboundDTO struct {

    // 配送员
    DelivererName   string `json:"deliverer_name,omitempty"`

    // 配送员电话
    DelivererPhone   string `json:"deliverer_phone,omitempty"`

    // 物流公司代码
    LogisticsCompanyCode   string `json:"logistics_company_code,omitempty"`

    // 物流单号
    LogisticsNo   string `json:"logistics_no,omitempty"`

    // 出库明细
    OutboundDetails   []OutboundDetailDTO `json:"outbound_details,omitempty"`

    // 发货时间
    SendOutTime   string `json:"send_out_time,omitempty"`

    // OC订单号
    TradeNo   string `json:"trade_no,omitempty"`

}
