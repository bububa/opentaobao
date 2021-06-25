package wms

// Consignsendinfo 
type Consignsendinfo struct {

    // 订单信息
    OrderItemList   []Orderitemlist `json:"order_item_list,omitempty"`

    // 运单信息
    TmsOrderList   []Tmsorderlist `json:"tms_order_list,omitempty"`

    // 菜鸟订单编码
    CnOrderCode   string `json:"cn_order_code,omitempty"`

    // ERP订单编码
    OrderCode   string `json:"order_code,omitempty"`

    // 订单类型 201 销售出库 502 换货出库 503 补发出库
    OrderType   int64 `json:"order_type,omitempty"`

    // 仓储编码
    StoreCode   string `json:"store_code,omitempty"`

    // 订单状态 WMS_ACCEPT 接单成功 WMS_REJECT 接单失败 WMS_CONFIRMED 仓库生产完成
    Status   string `json:"status,omitempty"`

    // 仓库订单完成时间
    ConfirmTime   string `json:"confirm_time,omitempty"`

    // 发票确认信息列表
    InvoinceConfirmList   []Invoinceconfirmlist `json:"invoince_confirm_list,omitempty"`

}
