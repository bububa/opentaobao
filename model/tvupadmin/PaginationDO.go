package tvupadmin

// PaginationDO 
type PaginationDO struct {

    // 总数
    
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
    

    // 页码
    
    PageNo   int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
    

    // 单页数量
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // 内容列表
    
    List   []AdvertScheduleDO `json:"list,omitempty" xml:"list,omitempty"`
    

}
