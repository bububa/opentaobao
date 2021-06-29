package alihealth2

// AlibabaAlihealthDentalStoreAuditQueryMtopResult 
type AlibabaAlihealthDentalStoreAuditQueryMtopResult struct {
    // model
    StoreAuditVoList   []StoreAuditVo `json:"store_audit_vo_list,omitempty" xml:"store_audit_vo_list>store_audit_vo,omitempty"`
    // msg_code
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // msg_info
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // true
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
