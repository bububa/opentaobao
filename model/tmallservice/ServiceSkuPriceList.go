package tmallservice

// ServiceSkuPriceList 
type ServiceSkuPriceList struct {

    // 价格
    Price   int64 `json:"price,omitempty"`

    // 服务sku
    ServiceAbilityCode   string `json:"service_ability_code,omitempty"`

}
