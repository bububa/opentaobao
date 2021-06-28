package logistic

// ReachableDto 
type ReachableDto struct {

    // 阻断原因
    
    InterruptReason   string `json:"interrupt_reason,omitempty" xml:"interrupt_reason,omitempty"`
    

    // 是否阻断 true:阻断  false:可达
    
    InterruptApplyWaybillCode   bool `json:"interrupt_apply_waybill_code,omitempty" xml:"interrupt_apply_waybill_code,omitempty"`
    

}
