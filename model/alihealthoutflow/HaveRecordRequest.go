package alihealthoutflow

// HaveRecordRequest 
type HaveRecordRequest struct {

    // 疾病编码:(糖尿病:1)
    
    DiseaseCode   string `json:"disease_code,omitempty" xml:"disease_code,omitempty"`
    

    // 生活号渠道编码：（ 控糖家园:2019112569406507）
    
    PublicId   string `json:"public_id,omitempty" xml:"public_id,omitempty"`
    

    // 支付宝id
    
    AlipayUserId   string `json:"alipay_user_id,omitempty" xml:"alipay_user_id,omitempty"`
    

}
