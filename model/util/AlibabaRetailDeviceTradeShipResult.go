package util

// AlibabaRetailDeviceTradeShipResult 
/* model for simplify = false
type AlibabaRetailDeviceTradeShipResult struct {

    // errorInfos
    
    ErrorInfos  struct {
        ErrorInfo  []ErrorInfo `json:"error_info,omitempty"`
    } `json:"error_infos,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// AlibabaRetailDeviceTradeShipResult 
type AlibabaRetailDeviceTradeShipResult struct {

    // errorInfos
    ErrorInfos   []ErrorInfo `json:"error_infos,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
