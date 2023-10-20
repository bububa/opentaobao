package icbu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuCategorySchemaLevelGetAPIResponse (新)层级属性获取 API返回值
// alibaba.icbu.category.schema.level.get
//
// 将表单中层级属性的子属性返回
type AlibabaIcbuCategorySchemaLevelGetAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuCategorySchemaLevelGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuCategorySchemaLevelGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuCategorySchemaLevelGetAPIResponseModel).Reset()
}

// AlibabaIcbuCategorySchemaLevelGetAPIResponseModel is (新)层级属性获取 成功返回结果
type AlibabaIcbuCategorySchemaLevelGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_category_schema_level_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Top返回对象
	Result *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIcbuCategorySchemaLevelGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIcbuCategorySchemaLevelGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuCategorySchemaLevelGetAPIResponse)
	},
}

// GetAlibabaIcbuCategorySchemaLevelGetAPIResponse 从 sync.Pool 获取 AlibabaIcbuCategorySchemaLevelGetAPIResponse
func GetAlibabaIcbuCategorySchemaLevelGetAPIResponse() *AlibabaIcbuCategorySchemaLevelGetAPIResponse {
	return poolAlibabaIcbuCategorySchemaLevelGetAPIResponse.Get().(*AlibabaIcbuCategorySchemaLevelGetAPIResponse)
}

// ReleaseAlibabaIcbuCategorySchemaLevelGetAPIResponse 将 AlibabaIcbuCategorySchemaLevelGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuCategorySchemaLevelGetAPIResponse(v *AlibabaIcbuCategorySchemaLevelGetAPIResponse) {
	v.Reset()
	poolAlibabaIcbuCategorySchemaLevelGetAPIResponse.Put(v)
}
