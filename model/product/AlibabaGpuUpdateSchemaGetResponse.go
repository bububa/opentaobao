package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取产品编辑schema规则的接口 APIResponse
alibaba.gpu.update.schema.get

获取产品编辑schema规则的接口
*/
type AlibabaGpuUpdateSchemaGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_gpu_update_schema_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 参数产品ID对应的产品更新规则
    
    UpdateProductRule   string `json:"update_product_rule,omitempty" xml:"