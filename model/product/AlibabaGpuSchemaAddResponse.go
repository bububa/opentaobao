package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
使用schema文件发布产品 APIResponse
alibaba.gpu.schema.add

使用Schema文件发布一个产品
*/
type AlibabaGpuSchemaAddAPIResponse struct {
    model.CommonResponse
    Response *AlibabaGpuSchemaAddResponse `json:"alibaba_gpu_schema_add_response,omitempty"`
}

type AlibabaGpuSchemaAddResponse struct {

    // 产品发布的结果
    AddProductResult   string `json:"add_product_result,omitempty"`

}
