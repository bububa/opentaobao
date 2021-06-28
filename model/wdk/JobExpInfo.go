package wdk

// JobExpInfo 
type JobExpInfo struct {

    // 部门
    
    Department   string `json:"department,omitempty" xml:"department,omitempty"`
    

    // 离职原因
    
    DimissionReason   string `json:"dimission_reason,omitempty" xml:"dimission_reason,omitempty"`
    

    // 结束日期
    
    GmtEnd   string `json:"gmt_end,omitempty" xml:"gmt_end,omitempty"`
    

    // 开始日期
    
    GmtStart   string `json:"gmt_start,omitempty" xml:"gmt_start,omitempty"`
    

    // 职位
    
    Position   string `json:"position,omitempty" xml:"position,omitempty"`
    

    // 薪资（月）
    
    SalaryByMonth   string `json:"salary_by_month,omitempty" xml:"salary_by_month,omitempty"`
    

    // 工作单位
    
    Company   string `json:"company,omitempty" xml:"company,omitempty"`
    

}
