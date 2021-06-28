package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫门店商品编辑 APIResponse
tmall.item.store.schema.update

天猫门店商品编辑
*/
type TmallItemStoreSchemaUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_item_store_schema_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 无
    
    ApiResult   *TmallItemStoreSchemaUpdateApiResult `json:"api_result,omitempty" xml:"