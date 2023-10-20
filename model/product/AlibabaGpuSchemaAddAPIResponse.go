package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabagpuschemaaddAPIResponse 使用schema文件发布产品 API返回值
// alibaba.gpu.schema.add
//
// 使用Schema文件发布一个产品
type AlibabagpuschemaaddAPIResponse struct {
	model.CommonResponse
	AlibabagpuschemaaddAPIResponseModel
}

// AlibabagpuschemaaddAPIResponseModel is 使用schema文件发布产品 成功返回结果
type AlibabagpuschemaaddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_gpu_schema_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 产品发布的结果
	AddProductResult string `json:"add_product_result,omitempty" xml:"add_product_result,omitempty"`
}
