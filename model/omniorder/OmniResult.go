package omniorder

import (
	"sync"
)

// OmniResult 结构体
type OmniResult struct {
	// 返回查询到的sku列表
	Datas []ItemLightPublishSkuDto `json:"datas,omitempty" xml:"datas>item_light_publish_sku_dto,omitempty"`
	// 错误码，0表示成功
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// data
	Data *ItemDeleteResult `json:"data,omitempty" xml:"data,omitempty"`
}

var poolOmniResult = sync.Pool{
	New: func() any {
		return new(OmniResult)
	},
}

// GetOmniResult() 从对象池中获取OmniResult
func GetOmniResult() *OmniResult {
	return poolOmniResult.Get().(*OmniResult)
}

// ReleaseOmniResult 释放OmniResult
func ReleaseOmniResult(v *OmniResult) {
	v.Datas = v.Datas[:0]
	v.Code = ""
	v.Message = ""
	v.Data = nil
	poolOmniResult.Put(v)
}
