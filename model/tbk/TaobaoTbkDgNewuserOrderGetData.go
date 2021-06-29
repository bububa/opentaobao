package tbk

// TaobaoTbkDgNewuserOrderGetData 
type TaobaoTbkDgNewuserOrderGetData struct {
    // resultList
    Results   []TaobaoTbkDgNewuserOrderGetMapData `json:"results,omitempty" xml:"results>taobao_tbk_dg_newuser_order_get_map_data,omitempty"`
    // 页码
    PageNo   int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
    // 每页大小
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 是否有下一页
    HasNext   bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}
