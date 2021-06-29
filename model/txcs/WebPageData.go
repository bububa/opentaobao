package txcs

// WebPageData 
type WebPageData struct {

    // 页码信息
    
    Pagination   *Pagination `json:"pagination,omitempty" xml:"pagination,omitempty"`
    

    // 结果数据
    
    List   []WebPageDataList `json:"list,omitempty" xml:"list,omitempty"`
    

}
