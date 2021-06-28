package bus

// BusoMainOrderHistoryPageVo 
/* model for simplify = false
type BusoMainOrderHistoryPageVo struct {

    // busoMainOrderHistoryVOList
    
    BusoMainOrderHistoryVOList  struct {
        Busomainorderhistoryvolist  []Busomainorderhistoryvolist `json:"busomainorderhistoryvolist,omitempty"`
    } `json:"buso_main_order_history_v_o_list,omitempty"`
    

    // currentPage 当前多少页
    
    CurrentPage   int64 `json:"current_page,omitempty"`
    

    // totalSize 总条目
    
    TotalSize   int64 `json:"total_size,omitempty"`
    

}
*/

// BusoMainOrderHistoryPageVo 
type BusoMainOrderHistoryPageVo struct {

    // busoMainOrderHistoryVOList
    BusoMainOrderHistoryVOList   []Busomainorderhistoryvolist `json:"buso_main_order_history_v_o_list,omitempty"`

    // currentPage 当前多少页
    CurrentPage   int64 `json:"current_page,omitempty"`

    // totalSize 总条目
    TotalSize   int64 `json:"total_size,omitempty"`

}
