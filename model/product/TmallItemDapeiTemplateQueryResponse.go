package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
搭配查询接口 APIResponse
tmall.item.dapei.template.query

根据条件获取搭配内容
*/
type TmallItemDapeiTemplateQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_item_dapei_template_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TmallItemDapeiTemplateQueryResultSet `json:"result,omitempty" xml:"