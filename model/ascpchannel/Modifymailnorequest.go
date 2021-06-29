package ascpchannel

// Modifymailnorequest 
type Modifymailnorequest struct {
    // 供应商id
    SupplierId   string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
    // 原运单号
    OldTmsOrderCode   string `json:"old_tms_order_code,omitempty" xml:"old_tms_order_code,omitempty"`
    // 新运单号
    NewTmsOrderCode   string `json:"new_tms_order_code,omitempty" xml:"new_tms_order_code,omitempty"`
    // 原配送公司编码
    OldTmsServiceCode   string `json:"old_tms_service_code,omitempty" xml:"old_tms_service_code,omitempty"`
    // 发货LP
    BizOrderCode   string `json:"biz_order_code,omitempty" xml:"biz_order_code,omitempty"`
}
