package btrip

// OpenCostCenterSaveRq 
type OpenCostCenterSaveRq struct {

    // 绑定支付宝账号
    
    AlipayNo   string `json:"alipay_no,omitempty" xml:"alipay_no,omitempty"`
    

    // 成本中心名称
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 适用范围:1全员，2部分人员
    
    Scope   int64 `json:"scope,omitempty" xml:"scope,omitempty"`
    

    // 第三方成本中心id
    
    ThirdpartId   string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
    

    // 第三方成本中心编号
    
    Number   string `json:"number,omitempty" xml:"number,omitempty"`
    

    // 企业id
    
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
    

    // 商旅开放平台传2
    
    Version   int64 `json:"version,omitempty" xml:"version,omitempty"`
    

}
