package tmallnr

// PageResultDo 
type PageResultDo struct {
    // 数据总数目
    TotalNum   int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
    // 错误码
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 门店集合
    DataList   []NrtStoreDto `json:"data_list,omitempty" xml:"data_list>nrt_store_dto,omitempty"`
    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    // 每页记录数
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 总页数
    TotalPageNum   int64 `json:"total_page_num,omitempty" xml:"total_page_num,omitempty"`
    // 当前页
    PageNum   int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
}
