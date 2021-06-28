package eticket

// PackageBase 
/* model for simplify = false
type PackageBase struct {

    // 包分账类型    0：不分账，1：门店分账，2：宝贝分账，3：账号分账。
    
    BillType   int64 `json:"bill_type,omitempty"`
    

    // 备注
    
    Memo   string `json:"memo,omitempty"`
    

    // 包名
    
    PackageName   string `json:"package_name,omitempty"`
    

    // 卖家ID
    
    SellerId   int64 `json:"seller_id,omitempty"`
    

    // 发码类型 0 不发码，1 淘宝发码， 2 信任卖家发码， 3 码商发码， 4 码库发码
    
    SendType   int64 `json:"send_type,omitempty"`
    

    // 是否核销放行   0：不核销放行，1：核销放行
    
    IsConsumePass   int64 `json:"is_consume_pass,omitempty"`
    

    // 是否关联门店：0:不关联，1:关联
    
    HasPos   int64 `json:"has_pos,omitempty"`
    

    // 系统自动生成，传入无效
    
    Version   int64 `json:"version,omitempty"`
    

    // 系统自动生成，传入无效
    
    BillVersion   int64 `json:"bill_version,omitempty"`
    

    // 发码方   0：淘宝，码商userId：码商，poolId：码池
    
    SendId   int64 `json:"send_id,omitempty"`
    

    // 如果是宝贝分账，即billType为2，则必须填写宝贝分账模板 (key-value 格式, 以 ; 分隔)
    
    AccountToBillMapStr   string `json:"account_to_bill_map_str,omitempty"`
    

    // 是否支持子账号核销：0不支持，1支持
    
    IsSubaccount   int64 `json:"is_subaccount,omitempty"`
    

    // 是否支持身份证核销：0:不支持，1:支持
    
    IsIdCard   int64 `json:"is_id_card,omitempty"`
    

    // 卖家昵称
    
    SellerNick   string `json:"seller_nick,omitempty"`
    

    // 系统自动生成，传入无效
    
    PackageId   int64 `json:"package_id,omitempty"`
    

    // 发码码商名字
    
    SendMerchantName   string `json:"send_merchant_name,omitempty"`
    

    // 核销码商id
    
    ConsumeMerchantId   string `json:"consume_merchant_id,omitempty"`
    

    // 核销码商名字
    
    ConsumeMerchantName   string `json:"consume_merchant_name,omitempty"`
    

}
*/

// PackageBase 
type PackageBase struct {

    // 包分账类型    0：不分账，1：门店分账，2：宝贝分账，3：账号分账。
    BillType   int64 `json:"bill_type,omitempty"`

    // 备注
    Memo   string `json:"memo,omitempty"`

    // 包名
    PackageName   string `json:"package_name,omitempty"`

    // 卖家ID
    SellerId   int64 `json:"seller_id,omitempty"`

    // 发码类型 0 不发码，1 淘宝发码， 2 信任卖家发码， 3 码商发码， 4 码库发码
    SendType   int64 `json:"send_type,omitempty"`

    // 是否核销放行   0：不核销放行，1：核销放行
    IsConsumePass   int64 `json:"is_consume_pass,omitempty"`

    // 是否关联门店：0:不关联，1:关联
    HasPos   int64 `json:"has_pos,omitempty"`

    // 系统自动生成，传入无效
    Version   int64 `json:"version,omitempty"`

    // 系统自动生成，传入无效
    BillVersion   int64 `json:"bill_version,omitempty"`

    // 发码方   0：淘宝，码商userId：码商，poolId：码池
    SendId   int64 `json:"send_id,omitempty"`

    // 如果是宝贝分账，即billType为2，则必须填写宝贝分账模板 (key-value 格式, 以 ; 分隔)
    AccountToBillMapStr   string `json:"account_to_bill_map_str,omitempty"`

    // 是否支持子账号核销：0不支持，1支持
    IsSubaccount   int64 `json:"is_subaccount,omitempty"`

    // 是否支持身份证核销：0:不支持，1:支持
    IsIdCard   int64 `json:"is_id_card,omitempty"`

    // 卖家昵称
    SellerNick   string `json:"seller_nick,omitempty"`

    // 系统自动生成，传入无效
    PackageId   int64 `json:"package_id,omitempty"`

    // 发码码商名字
    SendMerchantName   string `json:"send_merchant_name,omitempty"`

    // 核销码商id
    ConsumeMerchantId   string `json:"consume_merchant_id,omitempty"`

    // 核销码商名字
    ConsumeMerchantName   string `json:"consume_merchant_name,omitempty"`

}
