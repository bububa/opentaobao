package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
产品更新接口 APIResponse
alibaba.gpu.schema.update

产品更新接口
*/
type AlibabaGpuSchemaUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_gpu_schema_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 更新产品的结果
    
    UpdateProductResult   string `json:"update_product_result,omitempty" xml:"