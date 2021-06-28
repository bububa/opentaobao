package bus

// TaobaoBusBusnumberGetResultSet 
type TaobaoBusBusnumberGetResultSet struct {

    // errCode BUSNUMBER_SEARCH_NOBUS 找不到班次		POWER_D权限问题
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    

    // errMsg
    
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    

    // module
    
    Module   string `json:"module,omitempty" xml:"module,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
