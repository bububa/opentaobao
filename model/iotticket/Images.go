package iotticket

// Images 
type Images struct {

    // 图片类型：service-上门服务图片;cuSendProof-客户邮寄凭证;spSendProof-服务商邮寄凭证;abnormalImage-异常信息;purchaseVoucher-用户购买凭证
    
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    

    // 图片url
    
    Url   string `json:"url,omitempty" xml:"url,omitempty"`
    

}
