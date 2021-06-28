package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
天猫发布商品规则获取 APIResponse
tmall.item.add.simpleschema.get

通过商家信息获取商品发布字段和规则。
*/
type TmallItemAddSimpleschemaGetAPIResponse struct {
    model.CommonResponse
    // Response *TmallItemAddSimpleschemaGetResponse `json:"tmall_item_add_simpleschema_get_response,omitempty"` 
    TmallItemAddSimpleschemaGetResponse
}

/* model for simplify = false
type TmallItemAddSimpleschemaGetResponse struct {

    // 返回发布商品的规则文档
    
    Result   string `json:"result,omitempty"`
    

}
*/

type TmallItemAddSimpleschemaGetResponse struct {

    // 返回发布商品的规则文档
    Result   string `json:"result,omitempty"`

}
