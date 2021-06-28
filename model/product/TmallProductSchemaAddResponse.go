package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
使用Schema文件发布一个产品 APIResponse
tmall.product.schema.add

Schema体系发布一个产品
*/
type TmallProductSchemaAddAPIResponse struct {
    model.CommonResponse
    // Response *TmallProductSchemaAddResponse `json:"tmall_product_schema_add_response,omitempty"` 
    TmallProductSchemaAddResponse
}

/* model for simplify = false
type TmallProductSchemaAddResponse struct {

    // 新发产品结果
    
    AddProductResult   string `json:"add_product_result,omitempty"`
    

}
*/

type TmallProductSchemaAddResponse struct {

    // 新发产品结果
    AddProductResult   string `json:"add_product_result,omitempty"`

}
