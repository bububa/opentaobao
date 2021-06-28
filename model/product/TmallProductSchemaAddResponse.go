package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
使用Schema文件发布一个产品 APIResponse
tmall.product.schema.add

Schema体系发布一个产品
*/
type TmallProductSchemaAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_product_schema_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 新发产品结果
    
    AddProductResult   string `json:"add_product_result,omitempty" xml:"