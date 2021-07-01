package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取产品编辑schema规则的接口 API返回值 
alibaba.gpu.update.schema.get

获取产品编辑schema规则的接口
*/
type AlibabaGpuUpdateSchemaGetAPIResponse struct {
    model.CommonResponse
    AlibabaGpuUpdateSchemaGetAPIResponseModel
}

// 获取产品编辑schema规则的接口 成功返回结果
type AlibabaGpuUpdateSchemaGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_gpu_update_schema_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 参数产品ID对应的产品更新规则
    UpdateProductRule   string `json:"update_product_rule,omitempty" xml:"update_product_rule,omitempty"`
}
