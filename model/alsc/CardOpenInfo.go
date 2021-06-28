package alsc

// CardOpenInfo 
/* model for simplify = false
type CardOpenInfo struct {

    // 资产账户列表
    
    Accounts  struct {
        AccountOpenInfo  []AccountOpenInfo `json:"account_open_info,omitempty"`
    } `json:"accounts,omitempty"`
    

    // 激活操作人ID
    
    ActiveOperatorId   string `json:"active_operator_id,omitempty"`
    

    // 激活操作人名
    
    ActiveOperatorName   string `json:"active_operator_name,omitempty"`
    

    // 激活门店ID
    
    ActiveShopId   string `json:"active_shop_id,omitempty"`
    

    // 激活时间
    
    ActiveTime   string `json:"active_time,omitempty"`
    

    // 是否主动购买（是则赠送开卡礼包）
    
    Buy   bool `json:"buy,omitempty"`
    

    // 卡实例ID
    
    CardId   string `json:"card_id,omitempty"`
    

    // 卡模板ID
    
    CardTemplateId   string `json:"card_template_id,omitempty"`
    

    // 卡类型
    
    CardType   string `json:"card_type,omitempty"`
    

    // 创建人
    
    CreateBy   string `json:"create_by,omitempty"`
    

    // 顾客id
    
    CustomerId   string `json:"customer_id,omitempty"`
    

    // 逻辑删除
    
    Deleted   bool `json:"deleted,omitempty"`
    

    // 有效期结束时间
    
    ExpireEnd   string `json:"expire_end,omitempty"`
    

    // 有效期起始时间
    
    ExpireStart   string `json:"expire_start,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 修改时间
    
    GmtModified   string `json:"gmt_modified,omitempty"`
    

    // 开卡操作人ID
    
    OpenOperatorId   string `json:"open_operator_id,omitempty"`
    

    // 开卡操作人名
    
    OpenOperatorName   string `json:"open_operator_name,omitempty"`
    

    // 开卡门店ID
    
    OpenShopId   string `json:"open_shop_id,omitempty"`
    

    // 开卡门店名称
    
    OpenShopName   string `json:"open_shop_name,omitempty"`
    

    // 开卡时间
    
    OpenTime   string `json:"open_time,omitempty"`
    

    // 操作员id
    
    OperatorId   string `json:"operator_id,omitempty"`
    

    // 运营计划ID
    
    OptPlanId   string `json:"opt_plan_id,omitempty"`
    

    // 关联的物理卡列表
    
    PhysicalCards  struct {
        PhysicalCardOpenInfo  []PhysicalCardOpenInfo `json:"physical_card_open_info,omitempty"`
    } `json:"physical_cards,omitempty"`
    

    // 会员计划ID
    
    PlanId   string `json:"plan_id,omitempty"`
    

    // SOLD("SOLD", "已出售"),     ACTIVED("ACTIVED", "已激活"),     STOP("STOP", "已停用"),     INVALID("INVALID", "已作废"),     EXPIRED("EXPIRED", "已过期"),     REFUND("REFUND", "已退卡")
    
    Status   string `json:"status,omitempty"`
    

    // 修改人
    
    UpdateBy   string `json:"update_by,omitempty"`
    

    // 扩展信息
    
    ExtInfo  *struct {
        Extinfo  *Extinfo `json:"extinfo,omitempty"`
    } `json:"ext_info,omitempty"`
    

}
*/

// CardOpenInfo 
type CardOpenInfo struct {

    // 资产账户列表
    Accounts   []AccountOpenInfo `json:"accounts,omitempty"`

    // 激活操作人ID
    ActiveOperatorId   string `json:"active_operator_id,omitempty"`

    // 激活操作人名
    ActiveOperatorName   string `json:"active_operator_name,omitempty"`

    // 激活门店ID
    ActiveShopId   string `json:"active_shop_id,omitempty"`

    // 激活时间
    ActiveTime   string `json:"active_time,omitempty"`

    // 是否主动购买（是则赠送开卡礼包）
    Buy   bool `json:"buy,omitempty"`

    // 卡实例ID
    CardId   string `json:"card_id,omitempty"`

    // 卡模板ID
    CardTemplateId   string `json:"card_template_id,omitempty"`

    // 卡类型
    CardType   string `json:"card_type,omitempty"`

    // 创建人
    CreateBy   string `json:"create_by,omitempty"`

    // 顾客id
    CustomerId   string `json:"customer_id,omitempty"`

    // 逻辑删除
    Deleted   bool `json:"deleted,omitempty"`

    // 有效期结束时间
    ExpireEnd   string `json:"expire_end,omitempty"`

    // 有效期起始时间
    ExpireStart   string `json:"expire_start,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 开卡操作人ID
    OpenOperatorId   string `json:"open_operator_id,omitempty"`

    // 开卡操作人名
    OpenOperatorName   string `json:"open_operator_name,omitempty"`

    // 开卡门店ID
    OpenShopId   string `json:"open_shop_id,omitempty"`

    // 开卡门店名称
    OpenShopName   string `json:"open_shop_name,omitempty"`

    // 开卡时间
    OpenTime   string `json:"open_time,omitempty"`

    // 操作员id
    OperatorId   string `json:"operator_id,omitempty"`

    // 运营计划ID
    OptPlanId   string `json:"opt_plan_id,omitempty"`

    // 关联的物理卡列表
    PhysicalCards   []PhysicalCardOpenInfo `json:"physical_cards,omitempty"`

    // 会员计划ID
    PlanId   string `json:"plan_id,omitempty"`

    // SOLD("SOLD", "已出售"),     ACTIVED("ACTIVED", "已激活"),     STOP("STOP", "已停用"),     INVALID("INVALID", "已作废"),     EXPIRED("EXPIRED", "已过期"),     REFUND("REFUND", "已退卡")
    Status   string `json:"status,omitempty"`

    // 修改人
    UpdateBy   string `json:"update_by,omitempty"`

    // 扩展信息
    ExtInfo   *Extinfo `json:"ext_info,omitempty"`

}
