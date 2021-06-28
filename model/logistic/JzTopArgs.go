package logistic

// JzTopArgs 
/* model for simplify = false
type JzTopArgs struct {

    // 运单号
    
    MailNo   string `json:"mail_no,omitempty"`
    

    // 包裹重量
    
    PackageWeight   string `json:"package_weight,omitempty"`
    

    // 自有物流发货时间
    
    ZyConsignTime   string `json:"zy_consign_time,omitempty"`
    

    // 包裹体积
    
    PackageVolume   string `json:"package_volume,omitempty"`
    

    // 自有物流公司名称
    
    ZyCompany   string `json:"zy_company,omitempty"`
    

    // 包裹备注
    
    PackageRemark   string `json:"package_remark,omitempty"`
    

    // 自有物流公司电话
    
    ZyPhoneNumber   string `json:"zy_phone_number,omitempty"`
    

    // 包裹数量
    
    PackageNumber   string `json:"package_number,omitempty"`
    

}
*/

// JzTopArgs 
type JzTopArgs struct {

    // 运单号
    MailNo   string `json:"mail_no,omitempty"`

    // 包裹重量
    PackageWeight   string `json:"package_weight,omitempty"`

    // 自有物流发货时间
    ZyConsignTime   string `json:"zy_consign_time,omitempty"`

    // 包裹体积
    PackageVolume   string `json:"package_volume,omitempty"`

    // 自有物流公司名称
    ZyCompany   string `json:"zy_company,omitempty"`

    // 包裹备注
    PackageRemark   string `json:"package_remark,omitempty"`

    // 自有物流公司电话
    ZyPhoneNumber   string `json:"zy_phone_number,omitempty"`

    // 包裹数量
    PackageNumber   string `json:"package_number,omitempty"`

}
