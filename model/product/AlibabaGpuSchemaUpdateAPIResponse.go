package product

import (
	"encoding/xml"

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

// AlibabaGpuSchemaUpdateAPIResponseModel is 产品更新接口 成功返回结果
type AlibabaGpuSchemaUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_gpu_schema_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 更新产品的结果
	UpdateProductResult string `json:"update_product_result,omitempty" xml:"update_product_result,omitempty"`
}
