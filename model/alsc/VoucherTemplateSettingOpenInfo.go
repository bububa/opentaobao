package alsc

// VoucherTemplateSettingOpenInfo 
/* model for simplify = false
type VoucherTemplateSettingOpenInfo struct {

    // 渠道POS,移动门店
    
    OrderChannel   string `json:"order_channel,omitempty"`
    

    // 每天限制
    
    PerDay   int64 `json:"per_day,omitempty"`
    

    // 每单限制
    
    PerOrder   int64 `json:"per_order,omitempty"`
    

    // 适配的券模板id列表，为空代表适配全部
    
    VoucherTemplateIds  struct {
        String  []string `json:"string,omitempty"`
    } `json:"voucher_template_ids,omitempty"`
    

}
*/

// VoucherTemplateSettingOpenInfo 
type VoucherTemplateSettingOpenInfo struct {

    // 渠道POS,移动门店
    OrderChannel   string `json:"order_channel,omitempty"`

    // 每天限制
    PerDay   int64 `json:"per_day,omitempty"`

    // 每单限制
    PerOrder   int64 `json:"per_order,omitempty"`

    // 适配的券模板id列表，为空代表适配全部
    VoucherTemplateIds   []string `json:"voucher_template_ids,omitempty"`

}
