package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫发布商品规则获取 APIResponse
tmall.item.add.simpleschema.get

通过商家信息获取商品发布字段和规则。
*/
type TmallItemAddSimpleschemaGetAPIResponse struct {
    model.CommonResponse
    TmallItemAddSimpleschemaGetResponse
}

type TmallItemAddSimpleschemaGetResponse struct {
    XMLName xml.Name `xml:"tmall_item_add_simpleschema_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回发布商品的规则文档
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
