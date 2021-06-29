package ascpchannel

// Consignordershiprequest 
type Consignordershiprequest struct {
    // 供应商id
    SupplierId   string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
    // 外部业务号，幂等控制使用
    OutBizId   string `json:"out_biz_id,omitempty" xml:"out_biz_id,omitempty"`
    // 履约单号
    BizOrderCode   string `json:"biz_order_code,omitempty" xml:"biz_order_code,omitempty"`
    // 履约子单明细
    OrderItems   []Orderitems `json:"order_items,omitempty" xml:"order_items>orderitems,omitempty"`
    // 发货仓编码
    StoreCode   string `json:"store_code,omitempty" xml:"store_code,omitempty"`
    // 发货仓名称
    StoreName   string `json:"store_name,omitempty" xml:"store_name,omitempty"`
    // 包裹列表
    TmsOrders   []Tmsorders `json:"tms_orders,omitempty" xml:"tms_orders>tmsorders,omitempty"`
    // 发件人信息
    SenderInfo   *Senderinfo `json:"sender_info,omitempty" xml:"sender_info,omitempty"`
    // 是否整单发货,目前只支持履约单整单发货回传
    WholeSheetConsigned   bool `json:"whole_sheet_consigned,omitempty" xml:"whole_sheet_consigned,omitempty"`
}
