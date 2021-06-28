package alicom

// IdentityInfo 
/* model for simplify = false
type IdentityInfo struct {

    // 身份证背面图片
    
    BackImageUrl   string `json:"back_image_url,omitempty"`
    

    // 订单id
    
    BizOrderId   int64 `json:"biz_order_id,omitempty"`
    

    // 身份证号码
    
    CardNum   string `json:"card_num,omitempty"`
    

    // 身份证正面图片
    
    FrontImageUrl   string `json:"front_image_url,omitempty"`
    

    // 手持照
    
    HoldImageUrl   string `json:"hold_image_url,omitempty"`
    

    // 姓名
    
    Name   string `json:"name,omitempty"`
    

    // 地址
    
    Address   string `json:"address,omitempty"`
    

    // 是否长期有效
    
    LongTerm   bool `json:"long_term,omitempty"`
    

    // 卡类型
    
    CardType   string `json:"card_type,omitempty"`
    

    // 失效时间
    
    CardExpireDate   string `json:"card_expire_date,omitempty"`
    

    // 身份证生效时间
    
    StartDate   string `json:"start_date,omitempty"`
    

    // 身份证失效时间
    
    EndDate   string `json:"end_date,omitempty"`
    

    // biometricSeq
    
    BiometricSeq   string `json:"biometric_seq,omitempty"`
    

}
*/

// IdentityInfo 
type IdentityInfo struct {

    // 身份证背面图片
    BackImageUrl   string `json:"back_image_url,omitempty"`

    // 订单id
    BizOrderId   int64 `json:"biz_order_id,omitempty"`

    // 身份证号码
    CardNum   string `json:"card_num,omitempty"`

    // 身份证正面图片
    FrontImageUrl   string `json:"front_image_url,omitempty"`

    // 手持照
    HoldImageUrl   string `json:"hold_image_url,omitempty"`

    // 姓名
    Name   string `json:"name,omitempty"`

    // 地址
    Address   string `json:"address,omitempty"`

    // 是否长期有效
    LongTerm   bool `json:"long_term,omitempty"`

    // 卡类型
    CardType   string `json:"card_type,omitempty"`

    // 失效时间
    CardExpireDate   string `json:"card_expire_date,omitempty"`

    // 身份证生效时间
    StartDate   string `json:"start_date,omitempty"`

    // 身份证失效时间
    EndDate   string `json:"end_date,omitempty"`

    // biometricSeq
    BiometricSeq   string `json:"biometric_seq,omitempty"`

}
