package hotel

// SellerSupplierPartnerMemberInfoVo 
/* model for simplify = false
type SellerSupplierPartnerMemberInfoVo struct {

    // 会员卡名称
    
    CardName   string `json:"card_name,omitempty"`
    

    // 会员详细描述
    
    DetailMemo   string `json:"detail_memo,omitempty"`
    

    // 用户是否在第三方系统中绑定了会员,true--已绑定
    
    HasBinded   bool `json:"has_binded,omitempty"`
    

    // 会员等级
    
    MemberLevel   int64 `json:"member_level,omitempty"`
    

    // 会员对应的商家
    
    SellerId   int64 `json:"seller_id,omitempty"`
    

    // 会员对应的供应商
    
    Supplier   string `json:"supplier,omitempty"`
    

}
*/

// SellerSupplierPartnerMemberInfoVo 
type SellerSupplierPartnerMemberInfoVo struct {

    // 会员卡名称
    CardName   string `json:"card_name,omitempty"`

    // 会员详细描述
    DetailMemo   string `json:"detail_memo,omitempty"`

    // 用户是否在第三方系统中绑定了会员,true--已绑定
    HasBinded   bool `json:"has_binded,omitempty"`

    // 会员等级
    MemberLevel   int64 `json:"member_level,omitempty"`

    // 会员对应的商家
    SellerId   int64 `json:"seller_id,omitempty"`

    // 会员对应的供应商
    Supplier   string `json:"supplier,omitempty"`

}
