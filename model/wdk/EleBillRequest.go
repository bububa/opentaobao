package wdk

// EleBillRequest 
/* model for simplify = false
type EleBillRequest struct {

    // 查询页码,默认查询第一页,默认每页 20 条
    
    Page   string `json:"page,omitempty"`
    

    // 查询日期，时间戳格式(2019-06-10=1560124800)
    
    Date   string `json:"date,omitempty"`
    

    // 渠道店id
    
    ShopId   string `json:"shop_id,omitempty"`
    

}
*/

// EleBillRequest 
type EleBillRequest struct {

    // 查询页码,默认查询第一页,默认每页 20 条
    Page   string `json:"page,omitempty"`

    // 查询日期，时间戳格式(2019-06-10=1560124800)
    Date   string `json:"date,omitempty"`

    // 渠道店id
    ShopId   string `json:"shop_id,omitempty"`

}
