package icbu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuCategoryIdMappingAPIResponse 新旧属性的映射 API返回值
// alibaba.icbu.category.id.mapping
//
// 商品发布接口升级，需要传入新的类目。这个接口 根据旧的类目id，获取新的类目id
type AlibabaIcbuCategoryIdMappingAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuCategoryIdMappingAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuCategoryIdMappingAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuCategoryIdMappingAPIResponseModel).Reset()
}

// AlibabaIcbuCategoryIdMappingAPIResponseModel is 新旧属性的映射 成功返回结果
type AlibabaIcbuCategoryIdMappingAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_category_id_mapping_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 转化的类目id
	MappingResult int64 `json:"mapping_result,omitempty" xml:"mapping_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIcbuCategoryIdMappingAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MappingResult = 0
}

var poolAlibabaIcbuCategoryIdMappingAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuCategoryIdMappingAPIResponse)
	},
}

// GetAlibabaIcbuCategoryIdMappingAPIResponse 从 sync.Pool 获取 AlibabaIcbuCategoryIdMappingAPIResponse
func GetAlibabaIcbuCategoryIdMappingAPIResponse() *AlibabaIcbuCategoryIdMappingAPIResponse {
	return poolAlibabaIcbuCategoryIdMappingAPIResponse.Get().(*AlibabaIcbuCategoryIdMappingAPIResponse)
}

// ReleaseAlibabaIcbuCategoryIdMappingAPIResponse 将 AlibabaIcbuCategoryIdMappingAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuCategoryIdMappingAPIResponse(v *AlibabaIcbuCategoryIdMappingAPIResponse) {
	v.Reset()
	poolAlibabaIcbuCategoryIdMappingAPIResponse.Put(v)
}
