package tmallservice

// ServiceStoreSaveForTopReqDTO 
/* model for simplify = false
type ServiceStoreSaveForTopReqDTO struct {

    // 支付宝账号
    
    AlipayAccountIdNumber   string `json:"alipay_account_id_number,omitempty"`
    

    // 法人身份证反面l-使用阿里图片服务
    
    LegalPersonIdCardPicBack   string `json:"legal_person_id_card_pic_back,omitempty"`
    

    // 网点名称-唯一
    
    ServiceStoreName   string `json:"service_store_name,omitempty"`
    

    // 维度，必须是数字
    
    Latitude   string `json:"latitude,omitempty"`
    

    // 公司名称
    
    CompanyName   string `json:"company_name,omitempty"`
    

    // 扩展属性，底层必须是map结构
    
    Attributes   string `json:"attributes,omitempty"`
    

    // 法人身份证
    
    LegalPersonIdNumber   string `json:"legal_person_id_number,omitempty"`
    

    // 网点负责人姓名
    
    ManagerName   string `json:"manager_name,omitempty"`
    

    // 法人身份证照片l-使用阿里图片服务
    
    LegalPersonIdCardPic   string `json:"legal_person_id_card_pic,omitempty"`
    

    // 统一社会信用代码
    
    SocialCreditCode   string `json:"social_credit_code,omitempty"`
    

    // 网点地址编码经纬度和详细地址id二选一
    
    AddressId   int64 `json:"address_id,omitempty"`
    

    // 法人姓名
    
    LegalPersonName   string `json:"legal_person_name,omitempty"`
    

    // 品牌认证图片url-使用阿里图片服务
    
    BrandCertification   string `json:"brand_certification,omitempty"`
    

    // 网点邮箱
    
    ServiceStoreMail   string `json:"service_store_mail,omitempty"`
    

    // 网点负责人电话
    
    ManagerPhone   string `json:"manager_phone,omitempty"`
    

    // 精度，必须是数字
    
    Longitude   string `json:"longitude,omitempty"`
    

    // 门头照片-使用阿里图片服务
    
    FrontPhoto   string `json:"front_photo,omitempty"`
    

    // 详细地址
    
    Address   string `json:"address,omitempty"`
    

    // 认证的天猫品牌id集合-找运营支持
    
    CertificatedBrandIds   string `json:"certificated_brand_ids,omitempty"`
    

    // 营业时间,小时维度
    
    BusinessHours   string `json:"business_hours,omitempty"`
    

    // 支付宝账号id
    
    AlipayAccountId   string `json:"alipay_account_id,omitempty"`
    

    // 支付宝账号注册的信息--如用手机号或者邮箱
    
    AlipayAccount   string `json:"alipay_account,omitempty"`
    

    // 网点code和网点id选其一
    
    ServiceStoreCode   string `json:"service_store_code,omitempty"`
    

    // 营业执照照片-使用阿里图片服务
    
    LicensePhoto   string `json:"license_photo,omitempty"`
    

    // 支付宝账号真实姓名
    
    AlipayAccountName   string `json:"alipay_account_name,omitempty"`
    

    // 网点id和网点code选其一
    
    ServiceStoreId   int64 `json:"service_store_id,omitempty"`
    

    // 对外服务电话
    
    Phone   string `json:"phone,omitempty"`
    

}
*/

// ServiceStoreSaveForTopReqDTO 
type ServiceStoreSaveForTopReqDTO struct {

    // 支付宝账号
    AlipayAccountIdNumber   string `json:"alipay_account_id_number,omitempty"`

    // 法人身份证反面l-使用阿里图片服务
    LegalPersonIdCardPicBack   string `json:"legal_person_id_card_pic_back,omitempty"`

    // 网点名称-唯一
    ServiceStoreName   string `json:"service_store_name,omitempty"`

    // 维度，必须是数字
    Latitude   string `json:"latitude,omitempty"`

    // 公司名称
    CompanyName   string `json:"company_name,omitempty"`

    // 扩展属性，底层必须是map结构
    Attributes   string `json:"attributes,omitempty"`

    // 法人身份证
    LegalPersonIdNumber   string `json:"legal_person_id_number,omitempty"`

    // 网点负责人姓名
    ManagerName   string `json:"manager_name,omitempty"`

    // 法人身份证照片l-使用阿里图片服务
    LegalPersonIdCardPic   string `json:"legal_person_id_card_pic,omitempty"`

    // 统一社会信用代码
    SocialCreditCode   string `json:"social_credit_code,omitempty"`

    // 网点地址编码经纬度和详细地址id二选一
    AddressId   int64 `json:"address_id,omitempty"`

    // 法人姓名
    LegalPersonName   string `json:"legal_person_name,omitempty"`

    // 品牌认证图片url-使用阿里图片服务
    BrandCertification   string `json:"brand_certification,omitempty"`

    // 网点邮箱
    ServiceStoreMail   string `json:"service_store_mail,omitempty"`

    // 网点负责人电话
    ManagerPhone   string `json:"manager_phone,omitempty"`

    // 精度，必须是数字
    Longitude   string `json:"longitude,omitempty"`

    // 门头照片-使用阿里图片服务
    FrontPhoto   string `json:"front_photo,omitempty"`

    // 详细地址
    Address   string `json:"address,omitempty"`

    // 认证的天猫品牌id集合-找运营支持
    CertificatedBrandIds   string `json:"certificated_brand_ids,omitempty"`

    // 营业时间,小时维度
    BusinessHours   string `json:"business_hours,omitempty"`

    // 支付宝账号id
    AlipayAccountId   string `json:"alipay_account_id,omitempty"`

    // 支付宝账号注册的信息--如用手机号或者邮箱
    AlipayAccount   string `json:"alipay_account,omitempty"`

    // 网点code和网点id选其一
    ServiceStoreCode   string `json:"service_store_code,omitempty"`

    // 营业执照照片-使用阿里图片服务
    LicensePhoto   string `json:"license_photo,omitempty"`

    // 支付宝账号真实姓名
    AlipayAccountName   string `json:"alipay_account_name,omitempty"`

    // 网点id和网点code选其一
    ServiceStoreId   int64 `json:"service_store_id,omitempty"`

    // 对外服务电话
    Phone   string `json:"phone,omitempty"`

}
