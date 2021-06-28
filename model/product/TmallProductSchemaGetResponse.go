package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
产品信息获取schema获取 APIResponse
tmall.product.schema.get

产品信息获取接口schema形式返回
*/
type TmallProductSchemaGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_product_schema_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 产品信息数据。schema形式
    
    GetProductResult   string `json:"get_product_result,omitempty" xml:"