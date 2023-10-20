package tmallservice

import (
	"sync"
)

// EnableServiceStoreRequestDto 结构体
type EnableServiceStoreRequestDto struct {
	// 服务sku列表
	ServiceSkuList []string `json:"service_sku_list,omitempty" xml:"service_sku_list>string,omitempty"`
	// 门店名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 服务编码
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
	// 门店id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 是否启用
	Enable bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

var poolEnableServiceStoreRequestDto = sync.Pool{
	New: func() any {
		return new(EnableServiceStoreRequestDto)
	},
}

// GetEnableServiceStoreRequestDto() 从对象池中获取EnableServiceStoreRequestDto
func GetEnableServiceStoreRequestDto() *EnableServiceStoreRequestDto {
	return poolEnableServiceStoreRequestDto.Get().(*EnableServiceStoreRequestDto)
}

// ReleaseEnableServiceStoreRequestDto 释放EnableServiceStoreRequestDto
func ReleaseEnableServiceStoreRequestDto(v *EnableServiceStoreRequestDto) {
	v.ServiceSkuList = v.ServiceSkuList[:0]
	v.StoreName = ""
	v.ServiceCode = ""
	v.StoreId = 0
	v.Enable = false
	poolEnableServiceStoreRequestDto.Put(v)
}
