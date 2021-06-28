package bus

// TaobaoBusBusnumberGetResultSet 
/* model for simplify = false
type TaobaoBusBusnumberGetResultSet struct {

    // errCode BUSNUMBER_SEARCH_NOBUS 找不到班次		POWER_D权限问题
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // errMsg
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // module
    
    Module   string `json:"module,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// TaobaoBusBusnumberGetResultSet 
type TaobaoBusBusnumberGetResultSet struct {

    // errCode BUSNUMBER_SEARCH_NOBUS 找不到班次		POWER_D权限问题
    ErrCode   string `json:"err_code,omitempty"`

    // errMsg
    ErrMsg   string `json:"err_msg,omitempty"`

    // module
    Module   string `json:"module,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
