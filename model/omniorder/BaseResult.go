package omniorder

// BaseResult 
type BaseResult struct {

    // 错误信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

    // 返回的数据实体
    
    CommissionResultList   []CommissionResultDto `json:"commission_result_list,omitempty" xml:"commission_result_list,omitempty"`
    

    // 返回的执行状态吗
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

    // 是否执行成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 总记录数
    
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
    

    // 页大小
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // 当前页
    
    PageNo   int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
    

}
