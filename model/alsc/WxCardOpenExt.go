package alsc

// WxCardOpenExt 
/* model for simplify = false
type WxCardOpenExt struct {

    // 品牌logoURL
    
    BrandLogo   string `json:"brand_logo,omitempty"`
    

    // 统一卡面URL
    
    GeneralBgLogo   string `json:"general_bg_logo,omitempty"`
    

    // 是否统一卡面
    
    GeneralBgSwitch   bool `json:"general_bg_switch,omitempty"`
    

    // 支付后是否可领取
    
    PaidGetSwitch   bool `json:"paid_get_switch,omitempty"`
    

    // 等级卡面列表
    
    WxLevelBgs  struct {
        WxLevelBgOpenInfo  []WxLevelBgOpenInfo `json:"wx_level_bg_open_info,omitempty"`
    } `json:"wx_level_bgs,omitempty"`
    

}
*/

// WxCardOpenExt 
type WxCardOpenExt struct {

    // 品牌logoURL
    BrandLogo   string `json:"brand_logo,omitempty"`

    // 统一卡面URL
    GeneralBgLogo   string `json:"general_bg_logo,omitempty"`

    // 是否统一卡面
    GeneralBgSwitch   bool `json:"general_bg_switch,omitempty"`

    // 支付后是否可领取
    PaidGetSwitch   bool `json:"paid_get_switch,omitempty"`

    // 等级卡面列表
    WxLevelBgs   []WxLevelBgOpenInfo `json:"wx_level_bgs,omitempty"`

}
