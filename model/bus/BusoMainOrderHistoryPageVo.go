package bus

// BusoMainOrderHistoryPageVo 
type BusoMainOrderHistoryPageVo struct {
    // busoMainOrderHistoryVOList
    BusoMainOrderHistoryVOList   []Busomainorderhistoryvolist `json:"buso_main_order_history_v_o_list,omitempty" xml:"buso_main_order_history_v_o_list>busomainorderhistoryvolist,omitempty"`
    // currentPage 当前多少页
    CurrentPage   int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
    // totalSize 总条目
    TotalSize   int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`
}
