package aecreatives

// AliexpressAffiliateCategoryGetResult 
type AliexpressAffiliateCategoryGetResult struct {

    // 类目信息
    
    Categories   []Category `json:"categories,omitempty" xml:"categories,omitempty"`
    

    // 返回结果数量
    
    TotalResultCount   int64 `json:"total_result_count,omitempty" xml:"total_result_count,omitempty"`
    

}
