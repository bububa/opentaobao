package hotelhstdf

// GetByTrdiDivisionIdParam 
type GetByTrdiDivisionIdParam struct {

    // 第1页
    
    Page   int64 `json:"page,omitempty" xml:"page,omitempty"`
    

    // 每页取100条数据
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // 平台行政区划id，北京市
    
    TrdiDivisionId   int64 `json:"trdi_division_id,omitempty" xml:"trdi_division_id,omitempty"`
    

}
