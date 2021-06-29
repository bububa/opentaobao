package moscm

// DeliveryDTO 
type DeliveryDTO struct {
    // 承运公司名称
    CompanyName   string `json:"company_name,omitempty" xml:"company_name,omitempty"`
    // 订单号
    OrderNumber   string `json:"order_number,omitempty" xml:"order_number,omitempty"`
    // 承运公司编码
    CompanyCode   string `json:"company_code,omitempty" xml:"company_code,omitempty"`
    // 出库时间
    OutboundDate   string `json:"outbound_date,omitempty" xml:"outbound_date,omitempty"`
    // 运单号
    WaybillNumber   string `json:"waybill_number,omitempty" xml:"waybill_number,omitempty"`
    // 商品明细
    ShipItems   []ShipItemDTO `json:"ship_items,omitempty" xml:"ship_items>ship_item_dto,omitempty"`
    // 备注
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
}
