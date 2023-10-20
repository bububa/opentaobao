package tmallservice

import (
	"sync"
)

// ServiceSkuPriceList 结构体
type ServiceSkuPriceList struct {
	// 服务sku
	ServiceAbilityCode string `json:"service_ability_code,omitempty" xml:"service_ability_code,omitempty"`
	// 价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
}

var poolServiceSkuPriceList = sync.Pool{
	New: func() any {
		return new(ServiceSkuPriceList)
	},
}

// GetServiceSkuPriceList() 从对象池中获取ServiceSkuPriceList
func GetServiceSkuPriceList() *ServiceSkuPriceList {
	return poolServiceSkuPriceList.Get().(*ServiceSkuPriceList)
}

// ReleaseServiceSkuPriceList 释放ServiceSkuPriceList
func ReleaseServiceSkuPriceList(v *ServiceSkuPriceList) {
	v.ServiceAbilityCode = ""
	v.Price = 0
	poolServiceSkuPriceList.Put(v)
}
