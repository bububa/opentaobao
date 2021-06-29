package axintrade

// AxinPayRegisterCreateDTO 
type AxinPayRegisterCreateDTO struct {
    // 法人身份证
    LegalCertType   string `json:"legal_cert_type,omitempty" xml:"legal_cert_type,omitempty"`
    // 站点信息
    Sites   []AxinPayRegisterSiteInfo `json:"sites,omitempty" xml:"sites>axin_pay_register_site_info,omitempty"`
    // 结算支付宝账号
    AlipayLogonId   string `json:"alipay_logon_id,omitempty" xml:"alipay_logon_id,omitempty"`
    // 商户使用服务
    Service   []string `json:"service,omitempty" xml:"service>string,omitempty"`
    // 联系人电话
    ContactPhone   string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
    // 联系人名字
    ContactName   string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
    // 客服电话
    ServicePhone   string `json:"service_phone,omitempty" xml:"service_phone,omitempty"`
    // 营业地址
    BusinessAddress   *AxinPayRegisterAddress `json:"business_address,omitempty" xml:"business_address,omitempty"`
    // 法人身份证反面url
    LegalCertFrontImage   string `json:"legal_cert_front_image,omitempty" xml:"legal_cert_front_image,omitempty"`
    // 法人身份证件号
    LegalCertNo   string `json:"legal_cert_no,omitempty" xml:"legal_cert_no,omitempty"`
    // 法人名称
    LegalName   string `json:"legal_name,omitempty" xml:"legal_name,omitempty"`
    // 商户证件图片url
    CertImage   string `json:"cert_image,omitempty" xml:"cert_image,omitempty"`
    // 商户证件类型
    CertType   string `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
    // 商户证件编号
    CertNo   string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
    // 商户类别码mcc
    Mcc   string `json:"mcc,omitempty" xml:"mcc,omitempty"`
    // 商户类型
    MerchantType   string `json:"merchant_type,omitempty" xml:"merchant_type,omitempty"`
    // 商户别名
    AliasName   string `json:"alias_name,omitempty" xml:"alias_name,omitempty"`
    // 商户名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 商户编号
    ExternalId   string `json:"external_id,omitempty" xml:"external_id,omitempty"`
    // 法人身份证反面url
    LegalCertBackImage   string `json:"legal_cert_back_image,omitempty" xml:"legal_cert_back_image,omitempty"`
    // 行业资格证
    Qualifications   []AxinPayRegisterQualification `json:"qualifications,omitempty" xml:"qualifications>axin_pay_register_qualification,omitempty"`
    // 支付宝签约账号
    BindingAlipayLogonId   string `json:"binding_alipay_logon_id,omitempty" xml:"binding_alipay_logon_id,omitempty"`
}
