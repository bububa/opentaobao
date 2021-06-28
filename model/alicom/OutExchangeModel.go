package alicom

// OutExchangeModel 
/* model for simplify = false
type OutExchangeModel struct {

    // 单位为分
    
    Amount   string `json:"amount,omitempty"`
    

    // 淘宝nick
    
    Nick   string `json:"nick,omitempty"`
    

    // 外部单号
    
    OutId   string `json:"out_id,omitempty"`
    

    // 单位为分
    
    Account   string `json:"account,omitempty"`
    

    // 固定值填15
    
    BizType   string `json:"biz_type,omitempty"`
    

    // 商户合作id
    
    Partner   string `json:"partner,omitempty"`
    

    // 来源
    
    From   string `json:"from,omitempty"`
    

    // 兑换的购物券面额（单位分）
    
    CouponFace   string `json:"coupon_face,omitempty"`
    

    // 扩展属性
    
    Ext   string `json:"ext,omitempty"`
    

}
*/

// OutExchangeModel 
type OutExchangeModel struct {

    // 单位为分
    Amount   string `json:"amount,omitempty"`

    // 淘宝nick
    Nick   string `json:"nick,omitempty"`

    // 外部单号
    OutId   string `json:"out_id,omitempty"`

    // 单位为分
    Account   string `json:"account,omitempty"`

    // 固定值填15
    BizType   string `json:"biz_type,omitempty"`

    // 商户合作id
    Partner   string `json:"partner,omitempty"`

    // 来源
    From   string `json:"from,omitempty"`

    // 兑换的购物券面额（单位分）
    CouponFace   string `json:"coupon_face,omitempty"`

    // 扩展属性
    Ext   string `json:"ext,omitempty"`

}
