package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaGpuUpdateSchemaGetAPIResponse 获取产品编辑schema规则的接口 API返回值
// alibaba.gpu.update.schema.get
//
// 获取产品编辑schema规则的接口
type AlibabaGpuUpdateSchemaGetAPIResponse struct {
	model.CommonResponse
	AlibabaGpuUpdateSchemaGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaGpuUpdateSchemaGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaGpuUpdateSchemaGetAPIResponseModel).Reset()
}

// AlibabaGpuUpdateSchemaGetAPIResponseModel is 获取产品编辑schema规则的接口 成功返回结果
type AlibabaGpuUpdateSchemaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_gpu_update_schema_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 参数产品ID对应的产品更新规则
	UpdateProductRule string `json:"update_product_rule,omitempty" xml:"update_product_rule,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaGpuUpdateSchemaGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.UpdateProductRule = ""
}

var poolAlibabaGpuUpdateSchemaGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaGpuUpdateSchemaGetAPIResponse)
	},
}

// GetAlibabaGpuUpdateSchemaGetAPIResponse 从 sync.Pool 获取 AlibabaGpuUpdateSchemaGetAPIResponse
func GetAlibabaGpuUpdateSchemaGetAPIResponse() *AlibabaGpuUpdateSchemaGetAPIResponse {
	return poolAlibabaGpuUpdateSchemaGetAPIResponse.Get().(*AlibabaGpuUpdateSchemaGetAPIResponse)
}

// ReleaseAlibabaGpuUpdateSchemaGetAPIResponse 将 AlibabaGpuUpdateSchemaGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaGpuUpdateSchemaGetAPIResponse(v *AlibabaGpuUpdateSchemaGetAPIResponse) {
	v.Reset()
	poolAlibabaGpuUpdateSchemaGetAPIResponse.Put(v)
}
