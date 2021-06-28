package iot

// AlibabaRetailDevicePayUrlGetResult 
/* model for simplify = false
type AlibabaRetailDevicePayUrlGetResult struct {

    // module
    
    Module   string `json:"module,omitempty"`
    

    // errorInfos
    
    ErrorInfos  struct {
        ErrorInfo  []ErrorInfo `json:"error_info,omitempty"`
    } `json:"error_infos,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// AlibabaRetailDevicePayUrlGetResult 
type AlibabaRetailDevicePayUrlGetResult struct {

    // module
    Module   string `json:"module,omitempty"`

    // errorInfos
    ErrorInfos   []ErrorInfo `json:"error_infos,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
