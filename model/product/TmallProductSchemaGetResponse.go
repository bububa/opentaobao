package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
产品信息获取schema获取 APIResponse
tmall.product.schema.get

产品信息获取接口schema形式返回
*/
type TmallProductSchemaGetAPIResponse struct {
    model.CommonResponse
    Response *TmallProductSchemaGetResponse `json:"tmall_product_schema_get_response,omitempty"`
}

type TmallProductSchemaGetResponse struct {

    // 产品信息数据。schema形式
    GetProductResult   string `json:"get_product_result,omitempty"`

}
