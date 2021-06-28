package ticket

// ItemEleCertInfo 
/* model for simplify = false
type ItemEleCertInfo struct {

    // 有效期 过期类型
    
    ExpiryDateType   int64 `json:"expiry_date_type,omitempty"`
    

    // 电子凭证 有效期 开始时间
    
    ExpiryDateStart   string `json:"expiry_date_start,omitempty"`
    

    // 电子凭证 有效期 结束时间
    
    ExpiryDateEnd   string `json:"expiry_date_end,omitempty"`
    

    // 电子凭证 有效期 天数
    
    ExpiryDays   int64 `json:"expiry_days,omitempty"`
    

    // 核销门店库id
    
    PackageId   int64 `json:"package_id,omitempty"`
    

    // 售中自动退款比例，0~100
    
    AutoRefundRate   int64 `json:"auto_refund_rate,omitempty"`
    

    // 过期自动退款比例，0~100
    
    ExpiredRefundRate   int64 `json:"expired_refund_rate,omitempty"`
    

    // 门票商品电子凭证信息必填，店铺联系方式
    
    ShopTel   string `json:"shop_tel,omitempty"`
    

    // 核销服务提供商
    
    MerchantName   string `json:"merchant_name,omitempty"`
    

}
*/

// ItemEleCertInfo 
type ItemEleCertInfo struct {

    // 有效期 过期类型
    ExpiryDateType   int64 `json:"expiry_date_type,omitempty"`

    // 电子凭证 有效期 开始时间
    ExpiryDateStart   string `json:"expiry_date_start,omitempty"`

    // 电子凭证 有效期 结束时间
    ExpiryDateEnd   string `json:"expiry_date_end,omitempty"`

    // 电子凭证 有效期 天数
    ExpiryDays   int64 `json:"expiry_days,omitempty"`

    // 核销门店库id
    PackageId   int64 `json:"package_id,omitempty"`

    // 售中自动退款比例，0~100
    AutoRefundRate   int64 `json:"auto_refund_rate,omitempty"`

    // 过期自动退款比例，0~100
    ExpiredRefundRate   int64 `json:"expired_refund_rate,omitempty"`

    // 门票商品电子凭证信息必填，店铺联系方式
    ShopTel   string `json:"shop_tel,omitempty"`

    // 核销服务提供商
    MerchantName   string `json:"merchant_name,omitempty"`

}
