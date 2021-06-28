package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
使用schema文件发布产品 APIResponse
alibaba.gpu.schema.add

使用Schema文件发布一个产品
*/
type AlibabaGpuSchemaAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_gpu_schema_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 产品发布的结果
    
    AddProductResult   string `json:"add_product_result,omitempty" xml:"