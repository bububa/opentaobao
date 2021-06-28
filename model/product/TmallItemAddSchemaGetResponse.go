package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
天猫发布商品规则获取 APIResponse
tmall.item.add.schema.get

通过类目以及productId获取商品发布规则；
*/
type TmallItemAddSchemaGetAPIResponse struct {
    model.CommonResponse
    // Response *TmallItemAddSchemaGetResponse `json:"tmall_item_add_schema_get_response,omitempty"` 
    TmallItemAddSchemaGetResponse
}

/* model for simplify = false
type TmallItemAddSchemaGetResponse struct {

    // 返回发布商品的规则文档
    
    AddItemResult   string `json:"add_item_result,omitempty"`
    

}
*/

type TmallItemAddSchemaGetResponse struct {

    // 返回发布商品的规则文档
    AddItemResult   string `json:"add_item_result,omitempty"`

}
