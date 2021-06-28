package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取产品发布规则接口 APIResponse
alibaba.gpu.add.schema.get

获取产品发布规则接口
*/
type AlibabaGpuAddSchemaGetAPIResponse struct {
    model.CommonResponse
    AlibabaGpuAddSchemaGetResponse
}

type AlibabaGpuAddSchemaGetResponse struct {
    XMLName xml.Name `xml:"alibaba_gpu_add_schema_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回产品发布规则
    
    AddProductRule   string `json:"add_product_rule,omitempty" xml:"add_product_rule,omitempty"`

    
}
