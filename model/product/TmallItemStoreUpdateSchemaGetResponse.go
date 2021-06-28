package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
天猫门店商品修改规则获取 APIResponse
tmall.item.store.update.schema.get

天猫门店商品修改规则获取
*/
type TmallItemStoreUpdateSchemaGetAPIResponse struct {
    model.CommonResponse
    // Response *TmallItemStoreUpdateSchemaGetResponse `json:"tmall_item_store_update_schema_get_response,omitempty"` 
    TmallItemStoreUpdateSchemaGetResponse
}

/* model for simplify = false
type TmallItemStoreUpdateSchemaGetResponse struct {

    // 无
    
    ApiResult  *struct {
        TmallItemStoreUpdateSchemaGetApiResult  *TmallItemStoreUpdateSchemaGetApiResult `json:"tmall_item_store_update_schema_get_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type TmallItemStoreUpdateSchemaGetResponse struct {

    // 无
    ApiResult   *TmallItemStoreUpdateSchemaGetApiResult `json:"api_result,omitempty"`

}
