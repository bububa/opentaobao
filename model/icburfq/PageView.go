package icburfq

// PageView 
type PageView struct {

    // 当前页面
    
    Current   int64 `json:"current,omitempty" xml:"current,omitempty"`
    

    // 页面大小
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // 推荐数量
    
    TotalItem   int64 `json:"total_item,omitempty" xml:"total_item,omitempty"`
    

    // 总页数
    
    TotalPages   int64 `json:"total_pages,omitempty" xml:"total_pages,omitempty"`
    

}
