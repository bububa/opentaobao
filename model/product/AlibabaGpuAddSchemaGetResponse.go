package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取产品发布规则接口 APIResponse
alibaba.gpu.add.schema.get

获取产品发布规则接口
*/
type AlibabaGpuAddSchemaGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaGpuAddSchemaGetResponse `json:"alibaba_gpu_add_schema_get_response,omitempty"`
}

type AlibabaGpuAddSchemaGetResponse struct {

    // 返回产品发布规则
    AddProductRule   string `json:"add_product_rule,omitempty"`

}
