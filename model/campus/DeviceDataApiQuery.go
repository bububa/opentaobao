package campus

// DeviceDataApiQuery 
type DeviceDataApiQuery struct {

    // 分页大小(最大500)
    
    Limit   int64 `json:"limit,omitempty" xml:"limit,omitempty"`
    

    // 启始时间(防止接口超时,建议不要传入时间跨度过大,如查询一个月内的数据)
    
    StartDate   string `json:"start_date,omitempty" xml:"start_date,omitempty"`
    

    // 参数点code(详细请查阅‘平台技术’下‘设备详细信息开发文档’。)
    
    PropertyCode   string `json:"property_code,omitempty" xml:"property_code,omitempty"`
    

    // 终止时间
    
    EndDate   string `json:"end_date,omitempty" xml:"end_date,omitempty"`
    

    // 当前页
    
    CurrentPage   int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
    

    // 设备id
    
    Uuid   string `json:"uuid,omitempty" xml:"uuid,omitempty"`
    

}
