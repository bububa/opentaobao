package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大商家商品编辑接口 APIResponse
tmall.item.vip.schema.update

大商家编辑商品
*/
type TmallItemVipSchemaUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_item_vip_schema_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 编辑商品的id
    
    UpdateItemResult   string `json:"update_item_result,omitempty" xml:"