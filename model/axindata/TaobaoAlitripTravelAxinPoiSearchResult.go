package axindata

// TaobaoAlitripTravelAxinPoiSearchResult 
type TaobaoAlitripTravelAxinPoiSearchResult struct {

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 列表
    
    DataList   []PoiVo `json:"data_list,omitempty" xml:"data_list,omitempty"`
    

}
