package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
组合货品关系查询接口 APIResponse
taobao.qimen.combineitem.query

组合货品关系查询
*/
type TaobaoQimenCombineitemQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_combineitem_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *Response `json:"response,omitempty" xml:"