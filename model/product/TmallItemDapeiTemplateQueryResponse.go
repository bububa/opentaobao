package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
搭配查询接口 APIResponse
tmall.item.dapei.template.query

根据条件获取搭配内容
*/
type TmallItemDapeiTemplateQueryAPIResponse struct {
    model.CommonResponse
    // Response *TmallItemDapeiTemplateQueryResponse `json:"tmall_item_dapei_template_query_response,omitempty"` 
    TmallItemDapeiTemplateQueryResponse
}

/* model for simplify = false
type TmallItemDapeiTemplateQueryResponse struct {

    // result
    
    Result  *struct {
        TmallItemDapeiTemplateQueryResultSet  *TmallItemDapeiTemplateQueryResultSet `json:"tmall_item_dapei_template_query_result_set,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallItemDapeiTemplateQueryResponse struct {

    // result
    Result   *TmallItemDapeiTemplateQueryResultSet `json:"result,omitempty"`

}
