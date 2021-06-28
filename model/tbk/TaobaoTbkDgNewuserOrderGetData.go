package tbk

// TaobaoTbkDgNewuserOrderGetData 
/* model for simplify = false
type TaobaoTbkDgNewuserOrderGetData struct {

    // resultList
    
    Results  struct {
        TaobaoTbkDgNewuserOrderGetMapData  []TaobaoTbkDgNewuserOrderGetMapData `json:"taobao_tbk_dg_newuser_order_get_map_data,omitempty"`
    } `json:"results,omitempty"`
    

    // 页码
    
    PageNo   int64 `json:"page_no,omitempty"`
    

    // 每页大小
    
    PageSize   int64 `json:"page_size,omitempty"`
    

    // 是否有下一页
    
    HasNext   bool `json:"has_next,omitempty"`
    

}
*/

// TaobaoTbkDgNewuserOrderGetData 
type TaobaoTbkDgNewuserOrderGetData struct {

    // resultList
    Results   []TaobaoTbkDgNewuserOrderGetMapData `json:"results,omitempty"`

    // 页码
    PageNo   int64 `json:"page_no,omitempty"`

    // 每页大小
    PageSize   int64 `json:"page_size,omitempty"`

    // 是否有下一页
    HasNext   bool `json:"has_next,omitempty"`

}
