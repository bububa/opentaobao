package tmallgeniescp

// AlibabaTmallgenieScpLocationGetData 
type AlibabaTmallgenieScpLocationGetData struct {

    // 租户
    
    Tenant   string `json:"tenant,omitempty" xml:"tenant,omitempty"`
    

    // CDC或者RDC名称
    
    LocationName   string `json:"location_name,omitempty" xml:"location_name,omitempty"`
    

    // CDC或者RDC类型
    
    LocationType   string `json:"location_type,omitempty" xml:"location_type,omitempty"`
    

    // 地点编码
    
    LocationCode   string `json:"location_code,omitempty" xml:"location_code,omitempty"`
    

    // 扩展参数
    
    ExtendJson   string `json:"extend_json,omitempty" xml:"extend_json,omitempty"`
    

}
