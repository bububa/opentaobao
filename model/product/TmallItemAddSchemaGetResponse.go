package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫发布商品规则获取 APIResponse
tmall.item.add.schema.get

通过类目以及productId获取商品发布规则；
*/
type TmallItemAddSchemaGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_item_add_schema_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回发布商品的规则文档
    
    AddItemResult   string `json:"add_item_result,omitempty" xml:"