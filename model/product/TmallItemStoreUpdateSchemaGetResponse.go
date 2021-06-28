package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫门店商品修改规则获取 APIResponse
tmall.item.store.update.schema.get

天猫门店商品修改规则获取
*/
type TmallItemStoreUpdateSchemaGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_item_store_update_schema_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 无
    
    ApiResult   *TmallItemStoreUpdateSchemaGetApiResult `json:"api_result,omitempty" xml:"