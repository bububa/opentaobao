package tmallservice

// ServiceStoreDTO 
type ServiceStoreDTO struct {

    // 城市
    AddressCity   string `json:"address_city,omitempty"`

    // 标准地址编码
    AddressCode   int64 `json:"address_code,omitempty"`

    // 详细地址
    AddressDetail   string `json:"address_detail,omitempty"`

    // 地址-区
    AddressDistrict   string `json:"address_district,omitempty"`

    // 地址-省
    AddressProvince   string `json:"address_province,omitempty"`

    // 标准地址-街道
    AddressTown   string `json:"address_town,omitempty"`

    // 支付宝账号
    AlipayAccount   string `json:"alipay_account,omitempty"`

    // 支付宝账号姓名
    AlipayAccountName   string `json:"alipay_account_name,omitempty"`

    // 额外属性。json
    Attributes   string `json:"attributes,omitempty"`

    // 业务身份。汽车传car_accessory_install
    BizType   string `json:"biz_type,omitempty"`

    // 品牌认证资料url列表,英文逗号分隔
    BrandCertification   string `json:"brand_certification,omitempty"`

    // 每日营业时间。字符串
    BusinessHours   string `json:"business_hours,omitempty"`

    // 认证的天猫品牌id列表
    CertificatedBrandIds   string `json:"certificated_brand_ids,omitempty"`

    // 公司名
    CompanyName   string `json:"company_name,omitempty"`

    // 纬度
    Latitude   string `json:"latitude,omitempty"`

    // 法人身份证号
    LegalPersonIdNumber   string `json:"legal_person_id_number,omitempty"`

    // 法人姓名
    LegalPersonName   string `json:"legal_person_name,omitempty"`

    // 经度
    Longitude   string `json:"longitude,omitempty"`

    // 经理人/联系人姓名
    ManagerName   string `json:"manager_name,omitempty"`

    // 经理人/联系人电话
    ManagerPhone   string `json:"manager_phone,omitempty"`

    // 对外服务座机号
    Phone   string `json:"phone,omitempty"`

    // 对外服务手机号
    ServiceMobile   string `json:"service_mobile,omitempty"`

    // 网点/门店编码
    ServiceStoreCode   string `json:"service_store_code,omitempty"`

    // 网点名称
    ServiceStoreName   string `json:"service_store_name,omitempty"`

    // 统一社会信用代码
    SocialCreditCode   string `json:"social_credit_code,omitempty"`

    // 门店类型
    Type   int64 `json:"type,omitempty"`

    // 别名
    ServiceStoreAlias   string `json:"service_store_alias,omitempty"`

    // 支付宝账户身份证号
    AlipayAccountIdNumber   string `json:"alipay_account_id_number,omitempty"`

    // 法人身份证正面
    LegalPersonIdCardPic   string `json:"legal_person_id_card_pic,omitempty"`

    // 法人身份证反面
    LegalPersonIdCardPicBack   string `json:"legal_person_id_card_pic_back,omitempty"`

    // 门头照片
    FrontPhoto   string `json:"front_photo,omitempty"`

    // 营业执照照片
    LicensePhoto   string `json:"license_photo,omitempty"`

}
