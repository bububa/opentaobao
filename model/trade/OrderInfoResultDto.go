package trade

// OrderInfoResultDto 
/* model for simplify = false
type OrderInfoResultDto struct {

    // 当前页
    
    CurPageNo   int64 `json:"cur_page_no,omitempty"`
    

    // 是否下一页
    
    HasNextPage   bool `json:"has_next_page,omitempty"`
    

    // 每页大小
    
    PageSize   int64 `json:"page_size,omitempty"`
    

    // 订单信息list
    
    OrderInfoList  struct {
        OrderInfoDto  []OrderInfoDto `json:"order_info_dto,omitempty"`
    } `json:"order_info_list,omitempty"`
    

}
*/

// OrderInfoResultDto 
type OrderInfoResultDto struct {

    // 当前页
    CurPageNo   int64 `json:"cur_page_no,omitempty"`

    // 是否下一页
    HasNextPage   bool `json:"has_next_page,omitempty"`

    // 每页大小
    PageSize   int64 `json:"page_size,omitempty"`

    // 订单信息list
    OrderInfoList   []OrderInfoDto `json:"order_info_list,omitempty"`

}
