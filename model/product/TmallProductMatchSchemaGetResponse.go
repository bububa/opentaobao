package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取匹配产品规则 APIResponse
tmall.product.match.schema.get

ISV发布商品前，需要先查找到产品ID，这个接口返回查找产品规则入参规则
*/
type TmallProductMatchSchemaGetAPIResponse struct {
    model.CommonResponse
    Response *TmallProductMatchSchemaGetResponse `json:"tmall_product_match_schema_get_response,omitempty"`
}

type TmallProductMatchSchemaGetResponse struct {

    // 返回匹配product的规则文档
    MatchResult   string `json:"match_result,omitempty"`

}
