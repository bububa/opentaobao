package icbu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuCategoryLevelAttrGetAPIResponse 层级属性的子属性获取 API返回值
// alibaba.icbu.category.level.attr.get
//
// 用于获取层级属性（车型库）的子属性和属性值
type AlibabaIcbuCategoryLevelAttrGetAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuCategoryLevelAttrGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuCategoryLevelAttrGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuCategoryLevelAttrGetAPIResponseModel).Reset()
}

// AlibabaIcbuCategoryLevelAttrGetAPIResponseModel is 层级属性的子属性获取 成功返回结果
type AlibabaIcbuCategoryLevelAttrGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_category_level_attr_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	ResultList *AlibabaIcbuCategoryLevelAttrGetResult `json:"result_list,omitempty" xml:"result_list,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIcbuCategoryLevelAttrGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = nil
}

var poolAlibabaIcbuCategoryLevelAttrGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuCategoryLevelAttrGetAPIResponse)
	},
}

// GetAlibabaIcbuCategoryLevelAttrGetAPIResponse 从 sync.Pool 获取 AlibabaIcbuCategoryLevelAttrGetAPIResponse
func GetAlibabaIcbuCategoryLevelAttrGetAPIResponse() *AlibabaIcbuCategoryLevelAttrGetAPIResponse {
	return poolAlibabaIcbuCategoryLevelAttrGetAPIResponse.Get().(*AlibabaIcbuCategoryLevelAttrGetAPIResponse)
}

// ReleaseAlibabaIcbuCategoryLevelAttrGetAPIResponse 将 AlibabaIcbuCategoryLevelAttrGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuCategoryLevelAttrGetAPIResponse(v *AlibabaIcbuCategoryLevelAttrGetAPIResponse) {
	v.Reset()
	poolAlibabaIcbuCategoryLevelAttrGetAPIResponse.Put(v)
}
