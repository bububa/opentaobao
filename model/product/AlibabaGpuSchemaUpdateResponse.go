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
    AlibabaGpuSchemaUpdateResponse
}

type AlibabaGpuSchemaUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_gpu_schema_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 更新产品的结果
    
    UpdateProductResult   string `json:"update_product_result,omitempty" xml:"update_product_result,omitempty"`

    
}
