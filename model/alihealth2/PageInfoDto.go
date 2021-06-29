package alihealth2

// PageInfoDto 
type PageInfoDto struct {

    // 总数
    
    TotalNum   int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
    

    // result
    
    Results   []AlibabaAlihealthTracecodesellerBillResultSearchResult `json:"results,omitempty" xml:"results,omitempty"`
    

    // page
    
    Page   int64 `json:"page,omitempty" xml:"page,omitempty"`
    

    // pageSize
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

}
