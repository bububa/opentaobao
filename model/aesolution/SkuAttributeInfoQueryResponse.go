package aesolution

import (
	"sync"
)

// SkuAttributeInfoQueryResponse 结构体
type SkuAttributeInfoQueryResponse struct {
	// supported sku attribute lis
	SupportingSkuAttributeList []SupportedSkuAttributeDto `json:"supporting_sku_attribute_list,omitempty" xml:"supporting_sku_attribute_list>supported_sku_attribute_dto,omitempty"`
	// common attributes under a specific category
	SupportingCommonAttributeList []SupportedCommonAttributeDto `json:"supporting_common_attribute_list,omitempty" xml:"supporting_common_attribute_list>supported_common_attribute_dto,omitempty"`
}

var poolSkuAttributeInfoQueryResponse = sync.Pool{
	New: func() any {
		return new(SkuAttributeInfoQueryResponse)
	},
}

// GetSkuAttributeInfoQueryResponse() 从对象池中获取SkuAttributeInfoQueryResponse
func GetSkuAttributeInfoQueryResponse() *SkuAttributeInfoQueryResponse {
	return poolSkuAttributeInfoQueryResponse.Get().(*SkuAttributeInfoQueryResponse)
}

// ReleaseSkuAttributeInfoQueryResponse 释放SkuAttributeInfoQueryResponse
func ReleaseSkuAttributeInfoQueryResponse(v *SkuAttributeInfoQueryResponse) {
	v.SupportingSkuAttributeList = v.SupportingSkuAttributeList[:0]
	v.SupportingCommonAttributeList = v.SupportingCommonAttributeList[:0]
	poolSkuAttributeInfoQueryResponse.Put(v)
}
