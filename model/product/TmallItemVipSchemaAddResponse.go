package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大商家商品发布接口 APIResponse
tmall.item.vip.schema.add

大商家商品发布接口
*/
type TmallItemVipSchemaAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_item_vip_schema_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品发布成功_商品id
    
    AddItemResult   string `json:"add_item_result,omitempty" xml:"