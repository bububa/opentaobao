package wdk

// StockShiftDetailDTO 
type StockShiftDetailDTO struct {
    // batchInfo
    BatchInfo   string `json:"batch_info,omitempty" xml:"batch_info,omitempty"`
    // outDeptCode
    OutDeptCode   string `json:"out_dept_code,omitempty" xml:"out_dept_code,omitempty"`
    // inDeptCode
    InDeptCode   string `json:"in_dept_code,omitempty" xml:"in_dept_code,omitempty"`
    // outCabinetCode
    OutCabinetCode   string `json:"out_cabinet_code,omitempty" xml:"out_cabinet_code,omitempty"`
    // inCabinetCode
    InCabinetCode   string `json:"in_cabinet_code,omitempty" xml:"in_cabinet_code,omitempty"`
    // quantity
    Quantity   string `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // itemCode
    ItemCode   string `json:"item_code,omitempty" xml:"item_code,omitempty"`
}
