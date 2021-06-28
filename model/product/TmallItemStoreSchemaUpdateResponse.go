package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
天猫门店商品编辑 APIResponse
tmall.item.store.schema.update

天猫门店商品编辑
*/
type TmallItemStoreSchemaUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TmallItemStoreSchemaUpdateResponse `json:"tmall_item_store_schema_update_response,omitempty"` 
    TmallItemStoreSchemaUpdateResponse
}

/* model for simplify = false
type TmallItemStoreSchemaUpdateResponse struct {

    // 无
    
    ApiResult  *struct {
        TmallItemStoreSchemaUpdateApiResult  *TmallItemStoreSchemaUpdateApiResult `json:"tmall_item_store_schema_update_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type TmallItemStoreSchemaUpdateResponse struct {

    // 无
    ApiResult   *TmallItemStoreSchemaUpdateApiResult `json:"api_result,omitempty"`

}
