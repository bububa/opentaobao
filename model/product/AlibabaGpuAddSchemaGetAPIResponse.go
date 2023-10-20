package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaGpuAddSchemaGetAPIResponse 获取产品发布规则接口 API返回值
// alibaba.gpu.add.schema.get
//
// 获取产品发布规则接口
type AlibabaGpuAddSchemaGetAPIResponse struct {
	model.CommonResponse
	AlibabaGpuAddSchemaGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaGpuAddSchemaGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaGpuAddSchemaGetAPIResponseModel).Reset()
}

// AlibabaGpuAddSchemaGetAPIResponseModel is 获取产品发布规则接口 成功返回结果
type AlibabaGpuAddSchemaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_gpu_add_schema_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回产品发布规则
	AddProductRule string `json:"add_product_rule,omitempty" xml:"add_product_rule,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaGpuAddSchemaGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AddProductRule = ""
}

var poolAlibabaGpuAddSchemaGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaGpuAddSchemaGetAPIResponse)
	},
}

// GetAlibabaGpuAddSchemaGetAPIResponse 从 sync.Pool 获取 AlibabaGpuAddSchemaGetAPIResponse
func GetAlibabaGpuAddSchemaGetAPIResponse() *AlibabaGpuAddSchemaGetAPIResponse {
	return poolAlibabaGpuAddSchemaGetAPIResponse.Get().(*AlibabaGpuAddSchemaGetAPIResponse)
}

// ReleaseAlibabaGpuAddSchemaGetAPIResponse 将 AlibabaGpuAddSchemaGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaGpuAddSchemaGetAPIResponse(v *AlibabaGpuAddSchemaGetAPIResponse) {
	v.Reset()
	poolAlibabaGpuAddSchemaGetAPIResponse.Put(v)
}
