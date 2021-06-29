package tmallservice

// OnePagedDataList 
type OnePagedDataList struct {
    // 服务编码
    ServiceCode   string `json:"service_code,omitempty" xml:"service_code,omitempty"`
    // 服务名称
    DisplayName   string `json:"display_name,omitempty" xml:"display_name,omitempty"`
    // 用户ID
    UserId   int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
    // 用户昵称
    UserNick   string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
    // 状态
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 服务sku封装
    ServiceAbilityCodeList   []ServiceSkuPriceList `json:"service_ability_code_list,omitempty" xml:"service_ability_code_list>service_sku_price_list,omitempty"`
}
