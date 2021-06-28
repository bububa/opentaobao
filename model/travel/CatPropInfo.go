package travel

// CatPropInfo 
/* model for simplify = false
type CatPropInfo struct {

    // 属性PID，调用taobao.itemprops.get取得
    
    Pid   string `json:"pid,omitempty"`
    

    // 属性VID，调用taobao.itempropvalues.get取得
    
    Vid   string `json:"vid,omitempty"`
    

}
*/

// CatPropInfo 
type CatPropInfo struct {

    // 属性PID，调用taobao.itemprops.get取得
    Pid   string `json:"pid,omitempty"`

    // 属性VID，调用taobao.itempropvalues.get取得
    Vid   string `json:"vid,omitempty"`

}
