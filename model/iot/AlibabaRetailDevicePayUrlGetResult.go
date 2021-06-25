package iot

// AlibabaRetailDevicePayUrlGetResult 
type AlibabaRetailDevicePayUrlGetResult struct {

    // module
    Module   string `json:"module,omitempty"`

    // errorInfos
    ErrorInfos   []ErrorInfo `json:"error_infos,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
