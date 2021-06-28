package qimen

// SnReportRequest 
/* model for simplify = false
type SnReportRequest struct {

    // 总页数
    
    TotalPage   int64 `json:"totalPage,omitempty"`
    

    // 当前页(从1开始)
    
    CurrentPage   int64 `json:"currentPage,omitempty"`
    

    // 每页记录的条数
    
    PageSize   int64 `json:"pageSize,omitempty"`
    

    // 发货单信息
    
    DeliveryOrder  *struct {
        DeliveryOrder  *DeliveryOrder `json:"delivery_order,omitempty"`
    } `json:"deliveryOrder,omitempty"`
    

    // 商品列表
    
    Items  struct {
        Item  []Item `json:"item,omitempty"`
    } `json:"items,omitempty"`
    

    // 扩展属性
    
    ExtendProps  *struct {
        Map  *Map `json:"map,omitempty"`
    } `json:"extendProps,omitempty"`
    

}
*/

// SnReportRequest 
type SnReportRequest struct {

    // 总页数
    TotalPage   int64 `json:"totalPage,omitempty"`

    // 当前页(从1开始)
    CurrentPage   int64 `json:"currentPage,omitempty"`

    // 每页记录的条数
    PageSize   int64 `json:"pageSize,omitempty"`

    // 发货单信息
    DeliveryOrder   *DeliveryOrder `json:"deliveryOrder,omitempty"`

    // 商品列表
    Items   []Item `json:"items,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

}
