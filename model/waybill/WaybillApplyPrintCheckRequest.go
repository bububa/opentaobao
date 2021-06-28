package waybill

// WaybillApplyPrintCheckRequest 
/* model for simplify = false
type WaybillApplyPrintCheckRequest struct {

    // 面单详情信息
    
    PrintCheckInfoCols  struct {
        PrintCheckInfo  []PrintCheckInfo `json:"print_check_info,omitempty"`
    } `json:"print_check_info_cols,omitempty"`
    

    // 物流服务商Code
    
    CpCode   string `json:"cp_code,omitempty"`
    

}
*/

// WaybillApplyPrintCheckRequest 
type WaybillApplyPrintCheckRequest struct {

    // 面单详情信息
    PrintCheckInfoCols   []PrintCheckInfo `json:"print_check_info_cols,omitempty"`

    // 物流服务商Code
    CpCode   string `json:"cp_code,omitempty"`

}
