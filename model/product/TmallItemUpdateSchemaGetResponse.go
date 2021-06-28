package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
天猫编辑商品规则获取 APIResponse
tmall.item.update.schema.get

Schema方式编辑天猫商品时，编辑商品规则获取
*/
type TmallItemUpdateSchemaGetAPIResponse struct {
    model.CommonResponse
    // Response *TmallItemUpdateSchemaGetResponse `json:"tmall_item_update_schema_get_response,omitempty"` 
    TmallItemUpdateSchemaGetResponse
}

/* model for simplify = false
type TmallItemUpdateSchemaGetResponse struct {

    // 返回发布商品的规则文档
    
    UpdateItemResult   string `json:"update_item_result,omitempty"`
    

}
*/

type TmallItemUpdateSchemaGetResponse struct {

    // 返回发布商品的规则文档
    UpdateItemResult   string `json:"update_item_result,omitempty"`

}
