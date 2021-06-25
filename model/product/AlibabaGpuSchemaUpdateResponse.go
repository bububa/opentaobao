package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
产品更新接口 APIResponse
alibaba.gpu.schema.update

产品更新接口
*/
type AlibabaGpuSchemaUpdateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaGpuSchemaUpdateResponse `json:"alibaba_gpu_schema_update_response,omitempty"`
}

type AlibabaGpuSchemaUpdateResponse struct {

    // 更新产品的结果
    UpdateProductResult   string `json:"update_product_result,omitempty"`

}
