package wms

// CainiaoBillQueryOrderinfolist 
/* model for simplify = false
type CainiaoBillQueryOrderinfolist struct {

    // 订单信息
    
    OrderInfo  *struct {
        CainiaoBillQueryOrderinfo  *CainiaoBillQueryOrderinfo `json:"cainiao_bill_query_orderinfo,omitempty"`
    } `json:"order_info,omitempty"`
    

}
*/

// CainiaoBillQueryOrderinfolist 
type CainiaoBillQueryOrderinfolist struct {

    // 订单信息
    OrderInfo   *CainiaoBillQueryOrderinfo `json:"order_info,omitempty"`

}
