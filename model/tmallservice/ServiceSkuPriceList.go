package tmallservice

// ServiceSkuPriceList 结构体
type ServiceSkuPriceList struct {
	// 服务sku
	ServiceAbilityCode string `json:"service_ability_code,omitempty" xml:"service_ability_code,omitempty"`
	// 价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
}
