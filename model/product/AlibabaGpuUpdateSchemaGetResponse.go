package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取产品编辑schema规则的接口 APIResponse
alibaba.gpu.update.schema.get

获取产品编辑schema规则的接口
*/
type AlibabaGpuUpdateSchemaGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaGpuUpdateSchemaGetResponse `json:"alibaba_gpu_update_schema_get_response,omitempty"`
}

type AlibabaGpuUpdateSchemaGetResponse struct {

    // 参数产品ID对应的产品更新规则
    UpdateProductRule   string `json:"update_product_rule,omitempty"`

}
