package logistic

// StoreInfo 
/* model for simplify = false
type StoreInfo struct {

    // 详细地址
    
    Address   string `json:"address,omitempty"`
    

    // XXX果园
    
    Name   string `json:"name,omitempty"`
    

    // 仓库真实名字
    
    RealName   string `json:"real_name,omitempty"`
    

    // 仓库服务代码
    
    ServiceCode   string `json:"service_code,omitempty"`
    

}
*/

// StoreInfo 
type StoreInfo struct {

    // 详细地址
    Address   string `json:"address,omitempty"`

    // XXX果园
    Name   string `json:"name,omitempty"`

    // 仓库真实名字
    RealName   string `json:"real_name,omitempty"`

    // 仓库服务代码
    ServiceCode   string `json:"service_code,omitempty"`

}
