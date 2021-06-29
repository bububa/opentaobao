package waybill

// WaybillDetailQueryByWaybillCodeRequest 
type WaybillDetailQueryByWaybillCodeRequest struct {
    // 快递公司code
    CpCode   string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
    // 请求id
    ObjectId   string `json:"object_id,omitempty" xml:"object_id,omitempty"`
    // 电子面单号
    WaybillCode   string `json:"waybill_code,omitempty" xml:"waybill_code,omitempty"`
}
