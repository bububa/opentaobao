package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
vip商家编辑商品的规则获取接口 APIResponse
tmall.item.vip.update.schema.get

获取vip商家编辑商品的规则
*/
type TmallItemVipUpdateSchemaGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_item_vip_update_schema_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 获取的编辑商品的规则
    
    UpdateGetResult   string `json:"update_get_result,omitempty" xml:"