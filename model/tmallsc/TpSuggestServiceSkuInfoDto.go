package tmallsc

import (
	"sync"
)

// TpSuggestServiceSkuInfoDto 结构体
type TpSuggestServiceSkuInfoDto struct {
	// 服务项
	ServiceSkuCode string `json:"service_sku_code,omitempty" xml:"service_sku_code,omitempty"`
	// 问题图片链接
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 服务商调整后的价格
	ServiceSkuPrice string `json:"service_sku_price,omitempty" xml:"service_sku_price,omitempty"`
}

var poolTpSuggestServiceSkuInfoDto = sync.Pool{
	New: func() any {
		return new(TpSuggestServiceSkuInfoDto)
	},
}

// GetTpSuggestServiceSkuInfoDto() 从对象池中获取TpSuggestServiceSkuInfoDto
func GetTpSuggestServiceSkuInfoDto() *TpSuggestServiceSkuInfoDto {
	return poolTpSuggestServiceSkuInfoDto.Get().(*TpSuggestServiceSkuInfoDto)
}

// ReleaseTpSuggestServiceSkuInfoDto 释放TpSuggestServiceSkuInfoDto
func ReleaseTpSuggestServiceSkuInfoDto(v *TpSuggestServiceSkuInfoDto) {
	v.ServiceSkuCode = ""
	v.PicUrl = ""
	v.ServiceSkuPrice = ""
	poolTpSuggestServiceSkuInfoDto.Put(v)
}
