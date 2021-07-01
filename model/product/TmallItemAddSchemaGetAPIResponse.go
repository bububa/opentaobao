package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫发布商品规则获取 API返回值 
tmall.item.add.schema.get

通过类目以及productId获取商品发布规则；
*/
type TmallItemAddSchemaGetAPIResponse struct {
    model.CommonResponse
    TmallItemAddSchemaGetAPIResponseModel
}

// 天猫发布商品规则获取 成功返回结果
type TmallItemAddSchemaGetAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_item_add_schema_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回发布商品的规则文档
    AddItemResult   string `json:"add_item_result,omitempty" xml:"add_item_result,omitempty"`
}
