package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
vip商家编辑商品的规则获取接口 APIResponse
tmall.item.vip.update.schema.get

获取vip商家编辑商品的规则
*/
type TmallItemVipUpdateSchemaGetAPIResponse struct {
    model.CommonResponse
    // Response *TmallItemVipUpdateSchemaGetResponse `json:"tmall_item_vip_update_schema_get_response,omitempty"` 
    TmallItemVipUpdateSchemaGetResponse
}

/* model for simplify = false
type TmallItemVipUpdateSchemaGetResponse struct {

    // 获取的编辑商品的规则
    
    UpdateGetResult   string `json:"update_get_result,omitempty"`
    

}
*/

type TmallItemVipUpdateSchemaGetResponse struct {

    // 获取的编辑商品的规则
    UpdateGetResult   string `json:"update_get_result,omitempty"`

}
