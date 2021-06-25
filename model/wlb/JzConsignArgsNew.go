package wlb

// JzConsignArgsNew 
type JzConsignArgsNew struct {

    // 快递运单号，serviceType=20 和serviceType=21时必填
    MailNo   string `json:"mail_no,omitempty"`

    // 包裹重量kg
    PackageWeight   string `json:"package_weight,omitempty"`

    // 包裹体积m3
    PackageVolume   string `json:"package_volume,omitempty"`

    // 包裹数目
    PackageNumber   string `json:"package_number,omitempty"`

    // 包裹备注信息
    PackageRemark   string `json:"package_remark,omitempty"`

    // 运单号，tmsPartner.virualType=true时必填
    ZyMailNo   string `json:"zy_mail_no,omitempty"`

    // 物流公司名称，tmsPartner.virualType=true时必填
    ZyCompany   string `json:"zy_company,omitempty"`

    // 物流公司电话，tmsPartner.virualType=true时必填
    ZyPhoneNumber   string `json:"zy_phone_number,omitempty"`

    // 发货时间，tmsPartner.virualType=true时必填
    ZyConsignTime   string `json:"zy_consign_time,omitempty"`

}
