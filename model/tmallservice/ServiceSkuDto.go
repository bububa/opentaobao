package tmallservice

import (
	"sync"
)

// ServiceSkuDto 结构体
type ServiceSkuDto struct {
	// 服务技能code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}

var poolServiceSkuDto = sync.Pool{
	New: func() any {
		return new(ServiceSkuDto)
	},
}

// GetServiceSkuDto() 从对象池中获取ServiceSkuDto
func GetServiceSkuDto() *ServiceSkuDto {
	return poolServiceSkuDto.Get().(*ServiceSkuDto)
}

// ReleaseServiceSkuDto 释放ServiceSkuDto
func ReleaseServiceSkuDto(v *ServiceSkuDto) {
	v.Code = ""
	poolServiceSkuDto.Put(v)
}
