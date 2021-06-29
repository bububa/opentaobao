package larkiot

// BizListResult 
type BizListResult struct {

    // 业务码
    
    BizCode   string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
    

    // 业务信息
    
    BizMsg   string `json:"biz_msg,omitempty" xml:"biz_msg,omitempty"`
    

    // 访问结果
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 影院列表
    
    DataList   []CinemaIotResponse `json:"data_list,omitempty" xml:"data_list,omitempty"`
    

}
