package tmallcar

// ServiceReceiptInfo4IsvDto 
type ServiceReceiptInfo4IsvDto struct {

    // 具体地址
    
    Address   string `json:"address,omitempty" xml:"address,omitempty"`
    

    // 市
    
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    

    // 区
    
    County   string `json:"county,omitempty" xml:"county,omitempty"`
    

    // 手机号
    
    Mobile   string `json:"mobile,omitempty" xml:"mobile,omitempty"`
    

    // 省
    
    Province   string `json:"province,omitempty" xml:"province,omitempty"`
    

    // 店铺名称
    
    StoreName   string `json:"store_name,omitempty" xml:"store_name,omitempty"`
    

    // 古荡街道
    
    Town   string `json:"town,omitempty" xml:"town,omitempty"`
    

    // 消费者名称
    
    UserName   string `json:"user_name,omitempty" xml:"user_name,omitempty"`
    

    // 工单号
    
    ReceiptId   int64 `json:"receipt_id,omitempty" xml:"receipt_id,omitempty"`
    

    // 品牌
    
    Brand   string `json:"brand,omitempty" xml:"brand,omitempty"`
    

    // 选择问题列表，多个问题以英文逗号分隔
    
    ChooseProblems   string `json:"choose_problems,omitempty" xml:"choose_problems,omitempty"`
    

    // 问题描述图片列表，多个图片以英文逗号分隔
    
    ProblemDescPics   string `json:"problem_desc_pics,omitempty" xml:"problem_desc_pics,omitempty"`
    

    // 用户备注
    
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    

    // 工单状态： SUBMIT:预约成功，待服务 CANCEL:已取消 FINISH:服务已完成
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

    // 工单附加信息
    
    Extension   string `json:"extension,omitempty" xml:"extension,omitempty"`
    

}
