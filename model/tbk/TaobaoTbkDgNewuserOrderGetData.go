package tbk

// TaobaoTbkDgNewuserOrderGetData 
type TaobaoTbkDgNewuserOrderGetData struct {

    // resultList
    Results   []MapData `json:"results,omitempty"`

    // 页码
    PageNo   int64 `json:"page_no,omitempty"`

    // 每页大小
    PageSize   int64 `json:"page_size,omitempty"`

    // 是否有下一页
    HasNext   bool `json:"has_next,omitempty"`

}
