package tbk

// PageResult 
/* model for simplify = false
type PageResult struct {

    // 翻页的pageno
    
    PageNo   int64 `json:"page_no,omitempty"`
    

    // 翻页的pagesie
    
    PageSize   int64 `json:"page_size,omitempty"`
    

    // 一共能查询出来的结果总数
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

    // 处罚订单列表
    
    Results  struct {
        TaobaoTbkDgPunishOrderGetResult  []TaobaoTbkDgPunishOrderGetResult `json:"taobao_tbk_dg_punish_order_get_result,omitempty"`
    } `json:"results,omitempty"`
    

}
*/

// PageResult 
type PageResult struct {

    // 翻页的pageno
    PageNo   int64 `json:"page_no,omitempty"`

    // 翻页的pagesie
    PageSize   int64 `json:"page_size,omitempty"`

    // 一共能查询出来的结果总数
    TotalCount   int64 `json:"total_count,omitempty"`

    // 处罚订单列表
    Results   []TaobaoTbkDgPunishOrderGetResult `json:"results,omitempty"`

}
