package mozi

// PageTenantSubAdminsRequest 
type PageTenantSubAdminsRequest struct {

    // 是否返回总数量
    
    ReturnTotalSize   bool `json:"return_total_size,omitempty" xml:"return_total_size,omitempty"`
    

    // 页数
    
    PageNo   int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
    

    // 租户id
    
    TenantId   int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
    

    // 一页的条目
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

}
