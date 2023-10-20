package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaGpuSchemaAddAPIResponse 使用schema文件发布产品 API返回值
// alibaba.gpu.schema.add
//
// 使用Schema文件发布一个产品
type AlibabaGpuSchemaAddAPIResponse struct {
	model.CommonResponse
	AlibabaGpuSchemaAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaGpuSchemaAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaGpuSchemaAddAPIResponseModel).Reset()
}

// AlibabaGpuSchemaAddAPIResponseModel is 使用schema文件发布产品 成功返回结果
type AlibabaGpuSchemaAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_gpu_schema_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 产品发布的结果
	AddProductResult string `json:"add_product_result,omitempty" xml:"add_product_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaGpuSchemaAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AddProductResult = ""
}

var poolAlibabaGpuSchemaAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaGpuSchemaAddAPIResponse)
	},
}

// GetAlibabaGpuSchemaAddAPIResponse 从 sync.Pool 获取 AlibabaGpuSchemaAddAPIResponse
func GetAlibabaGpuSchemaAddAPIResponse() *AlibabaGpuSchemaAddAPIResponse {
	return poolAlibabaGpuSchemaAddAPIResponse.Get().(*AlibabaGpuSchemaAddAPIResponse)
}

// ReleaseAlibabaGpuSchemaAddAPIResponse 将 AlibabaGpuSchemaAddAPIResponse 保存到 sync.Pool
func ReleaseAlibabaGpuSchemaAddAPIResponse(v *AlibabaGpuSchemaAddAPIResponse) {
	v.Reset()
	poolAlibabaGpuSchemaAddAPIResponse.Put(v)
}
