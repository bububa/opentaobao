package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
产品发布规则获取接口 APIResponse
tmall.product.add.schema.get

获取用户发布产品的规则
*/
type TmallProductAddSchemaGetAPIResponse struct {
    model.CommonResponse
    // Response *TmallProductAddSchemaGetResponse `json:"tmall_product_add_schema_get_response,omitempty"` 
    TmallProductAddSchemaGetResponse
}

/* model for simplify = false
type TmallProductAddSchemaGetResponse struct {

    // 返回发布产品的规则文档
    
    AddProductRule   string `json:"add_product_rule,omitempty"`
    

}
*/

type TmallProductAddSchemaGetResponse struct {

    // 返回发布产品的规则文档
    AddProductRule   string `json:"add_product_rule,omitempty"`

}
