package alsc

// CardTemplateOpenInfo 
/* model for simplify = false
type CardTemplateOpenInfo struct {

    // 卡模板ID
    
    CardTemplateId   string `json:"card_template_id,omitempty"`
    

    // 卡类型
    
    CardType   string `json:"card_type,omitempty"`
    

    // 创建者
    
    CreateBy   string `json:"create_by,omitempty"`
    

    // 逻辑删除标志
    
    Deleted   bool `json:"deleted,omitempty"`
    

    // 有效期类型
    
    ExpireType   string `json:"expire_type,omitempty"`
    

    // 有效期值
    
    ExpireValue   string `json:"expire_value,omitempty"`
    

    // 扩展信息
    
    ExtInfo  *struct {
        Extinfo  *Extinfo `json:"extinfo,omitempty"`
    } `json:"ext_info,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 修改时间
    
    GmtModified   string `json:"gmt_modified,omitempty"`
    

    // 特价菜单ID
    
    MenuId   string `json:"menu_id,omitempty"`
    

    // 特价菜限制
    
    MenuLimitType   string `json:"menu_limit_type,omitempty"`
    

    // 特价菜单开关
    
    MenuSwitch   bool `json:"menu_switch,omitempty"`
    

    // 卡模板名称
    
    Name   string `json:"name,omitempty"`
    

    // 开卡礼包开关
    
    OpenGiftSwitch   bool `json:"open_gift_switch,omitempty"`
    

    // 运营计划ID
    
    OptPlanId   string `json:"opt_plan_id,omitempty"`
    

    // 工本费金额
    
    PhyCardFeeAmount   int64 `json:"phy_card_fee_amount,omitempty"`
    

    // 是否可退卡费
    
    PhyCardFeeBack   bool `json:"phy_card_fee_back,omitempty"`
    

    // 工本费开关
    
    PhyCardFeeSwitch   bool `json:"phy_card_fee_switch,omitempty"`
    

    // 开卡礼包-赠送预储金额
    
    PreRechargeAmount   int64 `json:"pre_recharge_amount,omitempty"`
    

    // 开卡礼包-有效期结束是否清零
    
    RechargeClear   bool `json:"recharge_clear,omitempty"`
    

    // 储值规则信息对象
    
    RechargeRuleOpenInfo  *struct {
        RechargeRuleOpenInfo  *RechargeRuleOpenInfo `json:"recharge_rule_open_info,omitempty"`
    } `json:"recharge_rule_open_info,omitempty"`
    

    // 储值门店列表
    
    RechargeShopIds  struct {
        String  []string `json:"string,omitempty"`
    } `json:"recharge_shop_ids,omitempty"`
    

    // 储值开关
    
    RechargeSwitch   bool `json:"recharge_switch,omitempty"`
    

    // 售价
    
    SellPrice   int64 `json:"sell_price,omitempty"`
    

    // 售卡门店分组ID
    
    SellShopGroupId   string `json:"sell_shop_group_id,omitempty"`
    

    // 售卡门店列表
    
    SellShopIds  struct {
        String  []string `json:"string,omitempty"`
    } `json:"sell_shop_ids,omitempty"`
    

    // 何时开始计算有效期
    
    StartexpireType   string `json:"startexpire_type,omitempty"`
    

    // 卡模板状态
    
    Status   string `json:"status,omitempty"`
    

    // 更新者
    
    UpdateBy   string `json:"update_by,omitempty"`
    

    // 使用限制
    
    UseLimitType   string `json:"use_limit_type,omitempty"`
    

    // 使用门店列表
    
    UseShopIds  struct {
        String  []string `json:"string,omitempty"`
    } `json:"use_shop_ids,omitempty"`
    

    // 开卡礼包-优惠券模板列表
    
    Vouchers  struct {
        OpenGiftVoucherOpenInfo  []OpenGiftVoucherOpenInfo `json:"open_gift_voucher_open_info,omitempty"`
    } `json:"vouchers,omitempty"`
    

    // 微信卡扩展
    
    WxCardExt  *struct {
        WxCardOpenExt  *WxCardOpenExt `json:"wx_card_open_ext,omitempty"`
    } `json:"wx_card_ext,omitempty"`
    

    // 微信卡开关
    
    WxCardSwitch   bool `json:"wx_card_switch,omitempty"`
    

    // 外部储值门店
    
    OutRechargeShopIds  struct {
        String  []string `json:"string,omitempty"`
    } `json:"out_recharge_shop_ids,omitempty"`
    

    // 外部出售门店
    
    OutSellShopIds  struct {
        String  []string `json:"string,omitempty"`
    } `json:"out_sell_shop_ids,omitempty"`
    

    // 外部适用门店
    
    OutUseShopIds  struct {
        String  []string `json:"string,omitempty"`
    } `json:"out_use_shop_ids,omitempty"`
    

    // 是否已经制卡
    
    Publish   bool `json:"publish,omitempty"`
    

}
*/

// CardTemplateOpenInfo 
type CardTemplateOpenInfo struct {

    // 卡模板ID
    CardTemplateId   string `json:"card_template_id,omitempty"`

    // 卡类型
    CardType   string `json:"card_type,omitempty"`

    // 创建者
    CreateBy   string `json:"create_by,omitempty"`

    // 逻辑删除标志
    Deleted   bool `json:"deleted,omitempty"`

    // 有效期类型
    ExpireType   string `json:"expire_type,omitempty"`

    // 有效期值
    ExpireValue   string `json:"expire_value,omitempty"`

    // 扩展信息
    ExtInfo   *Extinfo `json:"ext_info,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 特价菜单ID
    MenuId   string `json:"menu_id,omitempty"`

    // 特价菜限制
    MenuLimitType   string `json:"menu_limit_type,omitempty"`

    // 特价菜单开关
    MenuSwitch   bool `json:"menu_switch,omitempty"`

    // 卡模板名称
    Name   string `json:"name,omitempty"`

    // 开卡礼包开关
    OpenGiftSwitch   bool `json:"open_gift_switch,omitempty"`

    // 运营计划ID
    OptPlanId   string `json:"opt_plan_id,omitempty"`

    // 工本费金额
    PhyCardFeeAmount   int64 `json:"phy_card_fee_amount,omitempty"`

    // 是否可退卡费
    PhyCardFeeBack   bool `json:"phy_card_fee_back,omitempty"`

    // 工本费开关
    PhyCardFeeSwitch   bool `json:"phy_card_fee_switch,omitempty"`

    // 开卡礼包-赠送预储金额
    PreRechargeAmount   int64 `json:"pre_recharge_amount,omitempty"`

    // 开卡礼包-有效期结束是否清零
    RechargeClear   bool `json:"recharge_clear,omitempty"`

    // 储值规则信息对象
    RechargeRuleOpenInfo   *RechargeRuleOpenInfo `json:"recharge_rule_open_info,omitempty"`

    // 储值门店列表
    RechargeShopIds   []string `json:"recharge_shop_ids,omitempty"`

    // 储值开关
    RechargeSwitch   bool `json:"recharge_switch,omitempty"`

    // 售价
    SellPrice   int64 `json:"sell_price,omitempty"`

    // 售卡门店分组ID
    SellShopGroupId   string `json:"sell_shop_group_id,omitempty"`

    // 售卡门店列表
    SellShopIds   []string `json:"sell_shop_ids,omitempty"`

    // 何时开始计算有效期
    StartexpireType   string `json:"startexpire_type,omitempty"`

    // 卡模板状态
    Status   string `json:"status,omitempty"`

    // 更新者
    UpdateBy   string `json:"update_by,omitempty"`

    // 使用限制
    UseLimitType   string `json:"use_limit_type,omitempty"`

    // 使用门店列表
    UseShopIds   []string `json:"use_shop_ids,omitempty"`

    // 开卡礼包-优惠券模板列表
    Vouchers   []OpenGiftVoucherOpenInfo `json:"vouchers,omitempty"`

    // 微信卡扩展
    WxCardExt   *WxCardOpenExt `json:"wx_card_ext,omitempty"`

    // 微信卡开关
    WxCardSwitch   bool `json:"wx_card_switch,omitempty"`

    // 外部储值门店
    OutRechargeShopIds   []string `json:"out_recharge_shop_ids,omitempty"`

    // 外部出售门店
    OutSellShopIds   []string `json:"out_sell_shop_ids,omitempty"`

    // 外部适用门店
    OutUseShopIds   []string `json:"out_use_shop_ids,omitempty"`

    // 是否已经制卡
    Publish   bool `json:"publish,omitempty"`

}
