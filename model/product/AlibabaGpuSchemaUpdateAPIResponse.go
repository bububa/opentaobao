package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaGpuSchemaUpdateAPIResponse 产品更新接口 API返回值
// alibaba.gpu.schema.update
//
// 产品更新接口
type AlibabaGpuSchemaUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaGpuSchemaUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaGpuSchemaUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaGpuSchemaUpdateAPIResponseModel).Reset()
}

// AlibabaGpuSchemaUpdateAPIResponseModel is 产品更新接口 成功返回结果
type AlibabaGpuSchemaUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_gpu_schema_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 更新产品的结果
	UpdateProductResult string `json:"update_product_result,omitempty" xml:"update_product_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaGpuSchemaUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.UpdateProductResult = ""
}

var poolAlibabaGpuSchemaUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaGpuSchemaUpdateAPIResponse)
	},
}

// GetAlibabaGpuSchemaUpdateAPIResponse 从 sync.Pool 获取 AlibabaGpuSchemaUpdateAPIResponse
func GetAlibabaGpuSchemaUpdateAPIResponse() *AlibabaGpuSchemaUpdateAPIResponse {
	return poolAlibabaGpuSchemaUpdateAPIResponse.Get().(*AlibabaGpuSchemaUpdateAPIResponse)
}

// ReleaseAlibabaGpuSchemaUpdateAPIResponse 将 AlibabaGpuSchemaUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaGpuSchemaUpdateAPIResponse(v *AlibabaGpuSchemaUpdateAPIResponse) {
	v.Reset()
	poolAlibabaGpuSchemaUpdateAPIResponse.Put(v)
}
