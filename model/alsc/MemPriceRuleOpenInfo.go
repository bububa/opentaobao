package alsc

// MemPriceRuleOpenInfo 
/* model for simplify = false
type MemPriceRuleOpenInfo struct {

    // 是否已删除
    
    Deleted   bool `json:"deleted,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 修改时间
    
    GmtModified   string `json:"gmt_modified,omitempty"`
    

    // 会员等级和特价菜单的关系
    
    LevelMenuList  struct {
        LevelMenuOpenInfo  []LevelMenuOpenInfo `json:"level_menu_open_info,omitempty"`
    } `json:"level_menu_list,omitempty"`
    

    // 支付方式限制，UNLIMITED_PAY：无限制，RECHARGE_PAY：储值整单支付
    
    PayType   string `json:"pay_type,omitempty"`
    

    // 扩展信息
    
    ExtInfo   string `json:"ext_info,omitempty"`
    

}
*/

// MemPriceRuleOpenInfo 
type MemPriceRuleOpenInfo struct {

    // 是否已删除
    Deleted   bool `json:"deleted,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 会员等级和特价菜单的关系
    LevelMenuList   []LevelMenuOpenInfo `json:"level_menu_list,omitempty"`

    // 支付方式限制，UNLIMITED_PAY：无限制，RECHARGE_PAY：储值整单支付
    PayType   string `json:"pay_type,omitempty"`

    // 扩展信息
    ExtInfo   string `json:"ext_info,omitempty"`

}
