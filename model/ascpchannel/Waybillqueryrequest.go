package ascpchannel

// Waybillqueryrequest 
type Waybillqueryrequest struct {
    // 发货LP单号列表
    ConsignLpOrderCodes   []Consignlpordercodes `json:"consign_lp_order_codes,omitempty" xml:"consign_lp_order_codes>consignlpordercodes,omitempty"`
    // 操作人id
    Operator   string `json:"operator,omitempty" xml:"operator,omitempty"`
    // 操作人名称
    OperatorName   string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
    // 配送公司Code
    CpCode   string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
    // 供应商id
    SupplierId   string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
    // 物流服务编码
    ServiceCode   string `json:"service_code,omitempty" xml:"service_code,omitempty"`
}
