package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫编辑商品规则获取 APIResponse
tmall.item.update.schema.get

Schema方式编辑天猫商品时，编辑商品规则获取
*/
type TmallItemUpdateSchemaGetAPIResponse struct {
    model.CommonResponse
    TmallItemUpdateSchemaGetResponse
}

type TmallItemUpdateSchemaGetResponse struct {
    XMLName xml.Name `xml:"tmall_item_update_schema_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回发布商品的规则文档
    
    UpdateItemResult   string `json:"update_item_result,omitempty" xml:"update_item_result,omitempty"`

    
}
