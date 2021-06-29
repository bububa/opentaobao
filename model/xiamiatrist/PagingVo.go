package xiamiatrist

// PagingVo 
type PagingVo struct {

    // 每页展示数
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // 页码
    
    Page   int64 `json:"page,omitempty" xml:"page,omitempty"`
    

}
