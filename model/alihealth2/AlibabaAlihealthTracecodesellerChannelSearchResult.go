package alihealth2

import (
	"sync"
)

// AlibabaAlihealthTracecodesellerChannelSearchResult 结构体
type AlibabaAlihealthTracecodesellerChannelSearchResult struct {
	// 区县
	AreaDesc string `json:"area_desc,omitempty" xml:"area_desc,omitempty"`
	// cityDesc
	CityDesc string `json:"city_desc,omitempty" xml:"city_desc,omitempty"`
	// userName
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// 城市
	MerchantId string `json:"merchant_id,omitempty" xml:"merchant_id,omitempty"`
	// 自定义编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 自定义编码
	ProvDesc string `json:"prov_desc,omitempty" xml:"prov_desc,omitempty"`
}

var poolAlibabaAlihealthTracecodesellerChannelSearchResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthTracecodesellerChannelSearchResult)
	},
}

// GetAlibabaAlihealthTracecodesellerChannelSearchResult() 从对象池中获取AlibabaAlihealthTracecodesellerChannelSearchResult
func GetAlibabaAlihealthTracecodesellerChannelSearchResult() *AlibabaAlihealthTracecodesellerChannelSearchResult {
	return poolAlibabaAlihealthTracecodesellerChannelSearchResult.Get().(*AlibabaAlihealthTracecodesellerChannelSearchResult)
}

// ReleaseAlibabaAlihealthTracecodesellerChannelSearchResult 释放AlibabaAlihealthTracecodesellerChannelSearchResult
func ReleaseAlibabaAlihealthTracecodesellerChannelSearchResult(v *AlibabaAlihealthTracecodesellerChannelSearchResult) {
	v.AreaDesc = ""
	v.CityDesc = ""
	v.UserName = ""
	v.MerchantId = ""
	v.Code = ""
	v.ProvDesc = ""
	poolAlibabaAlihealthTracecodesellerChannelSearchResult.Put(v)
}
