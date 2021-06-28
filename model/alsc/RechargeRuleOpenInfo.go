package alsc

// RechargeRuleOpenInfo 
/* model for simplify = false
type RechargeRuleOpenInfo struct {

    // 卡类型
    
    CardType   string `json:"card_type,omitempty"`
    

    // 创建人
    
    CreateBy   string `json:"create_by,omitempty"`
    

    // 扣减顺序：0-比例、1-先实储后赠送、2-先增储后实储
    
    DeductionOrder   string `json:"deduction_order,omitempty"`
    

    // 逻辑删除标志
    
    Deleted   bool `json:"deleted,omitempty"`
    

    // 扩展信息
    
    ExtInfo  *struct {
        Extinfo  *Extinfo `json:"extinfo,omitempty"`
    } `json:"ext_info,omitempty"`
    

    // 满赠开关：false-关，true-开，默认关闭
    
    FullGiftSwitch   bool `json:"full_gift_switch,omitempty"`
    

    // 满赠类型：0-按金额（默认），1-按百分比、
    
    GiftType   string `json:"gift_type,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 修改时间
    
    GmtModified   string `json:"gmt_modified,omitempty"`
    

    // 规则明细列表
    
    ListRechargeRuleDetailInfoList  struct {
        RechargeRuleDetailOpenInfo  []RechargeRuleDetailOpenInfo `json:"recharge_rule_detail_open_info,omitempty"`
    } `json:"list_recharge_rule_detail_info_list,omitempty"`
    

    // 免密额度:-1表示不限额，单位为分，默认200_00L
    
    PayNoPwdCredit   int64 `json:"pay_no_pwd_credit,omitempty"`
    

    // 免密开关:0-关闭（默认），1-开启
    
    PayNoPwdSwitch   bool `json:"pay_no_pwd_switch,omitempty"`
    

    // 规则id
    
    RuleId   string `json:"rule_id,omitempty"`
    

    // 规则名称
    
    RuleName   string `json:"rule_name,omitempty"`
    

    // 门店id
    
    ShopId   string `json:"shop_id,omitempty"`
    

    // 修改人
    
    UpdateBy   string `json:"update_by,omitempty"`
    

    // 外部门店ID
    
    OutShopId   string `json:"out_shop_id,omitempty"`
    

    // 规则明细列表
    
    ListRechargeRuleDetailInfos  struct {
        RechargeRuleDetailOpenInfo  []RechargeRuleDetailOpenInfo `json:"recharge_rule_detail_open_info,omitempty"`
    } `json:"list_recharge_rule_detail_infos,omitempty"`
    

}
*/

// RechargeRuleOpenInfo 
type RechargeRuleOpenInfo struct {

    // 卡类型
    CardType   string `json:"card_type,omitempty"`

    // 创建人
    CreateBy   string `json:"create_by,omitempty"`

    // 扣减顺序：0-比例、1-先实储后赠送、2-先增储后实储
    DeductionOrder   string `json:"deduction_order,omitempty"`

    // 逻辑删除标志
    Deleted   bool `json:"deleted,omitempty"`

    // 扩展信息
    ExtInfo   *Extinfo `json:"ext_info,omitempty"`

    // 满赠开关：false-关，true-开，默认关闭
    FullGiftSwitch   bool `json:"full_gift_switch,omitempty"`

    // 满赠类型：0-按金额（默认），1-按百分比、
    GiftType   string `json:"gift_type,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 规则明细列表
    ListRechargeRuleDetailInfoList   []RechargeRuleDetailOpenInfo `json:"list_recharge_rule_detail_info_list,omitempty"`

    // 免密额度:-1表示不限额，单位为分，默认200_00L
    PayNoPwdCredit   int64 `json:"pay_no_pwd_credit,omitempty"`

    // 免密开关:0-关闭（默认），1-开启
    PayNoPwdSwitch   bool `json:"pay_no_pwd_switch,omitempty"`

    // 规则id
    RuleId   string `json:"rule_id,omitempty"`

    // 规则名称
    RuleName   string `json:"rule_name,omitempty"`

    // 门店id
    ShopId   string `json:"shop_id,omitempty"`

    // 修改人
    UpdateBy   string `json:"update_by,omitempty"`

    // 外部门店ID
    OutShopId   string `json:"out_shop_id,omitempty"`

    // 规则明细列表
    ListRechargeRuleDetailInfos   []RechargeRuleDetailOpenInfo `json:"list_recharge_rule_detail_infos,omitempty"`

}
