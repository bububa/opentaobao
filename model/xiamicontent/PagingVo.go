package xiamicontent

// PagingVo 
type PagingVo struct {

    // 页码
    
    Page   int64 `json:"page,omitempty" xml:"page,omitempty"`
    

    // 每页展示数
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // 总页数
    
    Pages   int64 `json:"pages,omitempty" xml:"pages,omitempty"`
    

    // 总歌曲数
    
    Count   int64 `json:"count,omitempty" xml:"count,omitempty"`
    

}
