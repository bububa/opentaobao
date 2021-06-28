package waybill

// WaybillDetailQueryByWaybillCodeRequest 
/* model for simplify = false
type WaybillDetailQueryByWaybillCodeRequest struct {

    // 快递公司code
    
    CpCode   string `json:"cp_code,omitempty"`
    

    // 请求id
    
    ObjectId   string `json:"object_id,omitempty"`
    

    // 电子面单号
    
    WaybillCode   string `json:"waybill_code,omitempty"`
    

}
*/

// WaybillDetailQueryByWaybillCodeRequest 
type WaybillDetailQueryByWaybillCodeRequest struct {

    // 快递公司code
    CpCode   string `json:"cp_code,omitempty"`

    // 请求id
    ObjectId   string `json:"object_id,omitempty"`

    // 电子面单号
    WaybillCode   string `json:"waybill_code,omitempty"`

}
