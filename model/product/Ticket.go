package product

// Ticket 
/* model for simplify = false
type Ticket struct {

    // 产品规格ID
    
    SpecId   int64 `json:"spec_id,omitempty"`
    

    // 创建人ID
    
    CreateUserId   int64 `json:"create_user_id,omitempty"`
    

    // 如果产品规格，需要商家审核，为商家审核用户ID
    
    AuditSellerId   int64 `json:"audit_seller_id,omitempty"`
    

    // 审核原因
    
    Reason   string `json:"reason,omitempty"`
    

    // 关于审核原因，更详细的说明
    
    Memo   string `json:"memo,omitempty"`
    

    // 1, "商家确认"<br/>2, "商家拒绝"<br/>3, "小二确认"<br/>4, "小二拒绝"<br/>5, "待商家处理"<br/>6, "商家审核超时"<br/>7, "待小二审核"<br/>9, "品牌商确认"<br/>10, "免审通过"<br/>14, "免审拒绝"
    
    Status   int64 `json:"status,omitempty"`
    

    // 产品规格申请时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 产品规格审核单最后修改时间
    
    GmtModified   string `json:"gmt_modified,omitempty"`
    

}
*/

// Ticket 
type Ticket struct {

    // 产品规格ID
    SpecId   int64 `json:"spec_id,omitempty"`

    // 创建人ID
    CreateUserId   int64 `json:"create_user_id,omitempty"`

    // 如果产品规格，需要商家审核，为商家审核用户ID
    AuditSellerId   int64 `json:"audit_seller_id,omitempty"`

    // 审核原因
    Reason   string `json:"reason,omitempty"`

    // 关于审核原因，更详细的说明
    Memo   string `json:"memo,omitempty"`

    // 1, "商家确认"<br/>2, "商家拒绝"<br/>3, "小二确认"<br/>4, "小二拒绝"<br/>5, "待商家处理"<br/>6, "商家审核超时"<br/>7, "待小二审核"<br/>9, "品牌商确认"<br/>10, "免审通过"<br/>14, "免审拒绝"
    Status   int64 `json:"status,omitempty"`

    // 产品规格申请时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 产品规格审核单最后修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

}
