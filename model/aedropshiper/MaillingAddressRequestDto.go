package aedropshiper

// MaillingAddressRequestDto 
type MaillingAddressRequestDto struct {

    // 地址信息
    
    Address   string `json:"address,omitempty" xml:"address,omitempty"`
    

    // 地址扩展信息
    
    Address2   string `json:"address2,omitempty" xml:"address2,omitempty"`
    

    // 城市
    
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    

    // 联系人
    
    ContactPerson   string `json:"contact_person,omitempty" xml:"contact_person,omitempty"`
    

    // 国家
    
    Country   string `json:"country,omitempty" xml:"country,omitempty"`
    

    // cpf税号
    
    Cpf   string `json:"cpf,omitempty" xml:"cpf,omitempty"`
    

    // 收货人全称
    
    FullName   string `json:"full_name,omitempty" xml:"full_name,omitempty"`
    

    // 国际化locale
    
    Locale   string `json:"locale,omitempty" xml:"locale,omitempty"`
    

    // 电话号码
    
    MobileNo   string `json:"mobile_no,omitempty" xml:"mobile_no,omitempty"`
    

    // 护照号
    
    PassportNo   string `json:"passport_no,omitempty" xml:"passport_no,omitempty"`
    

    // 护照日期
    
    PassportNoDate   string `json:"passport_no_date,omitempty" xml:"passport_no_date,omitempty"`
    

    // 护照签发机构
    
    PassportOrganization   string `json:"passport_organization,omitempty" xml:"passport_organization,omitempty"`
    

    // 电话所在国家区号
    
    PhoneCountry   string `json:"phone_country,omitempty" xml:"phone_country,omitempty"`
    

    // 省份
    
    Province   string `json:"province,omitempty" xml:"province,omitempty"`
    

    // 税号
    
    TaxNumber   string `json:"tax_number,omitempty" xml:"tax_number,omitempty"`
    

    // 邮政编码
    
    Zip   string `json:"zip,omitempty" xml:"zip,omitempty"`
    

    // 智利税号
    
    RutNo   string `json:"rut_no,omitempty" xml:"rut_no,omitempty"`
    

    // 外籍税号(韩国外籍须填登陆证号或者护照号)
    
    ForeignerPassportNo   string `json:"foreigner_passport_no,omitempty" xml:"foreigner_passport_no,omitempty"`
    

    // 是否是外籍
    
    IsForeigner   string `json:"is_foreigner,omitempty" xml:"is_foreigner,omitempty"`
    

}
