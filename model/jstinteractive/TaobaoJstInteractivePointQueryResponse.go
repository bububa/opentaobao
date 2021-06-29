package jstinteractive

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
互动积分查询接口 API返回值 
taobao.jst.interactive.point.query

查询用户的互动积分
*/
type TaobaoJstInteractivePointQueryAPIResponse struct {
    model.CommonResponse
    TaobaoJstInteractivePointQueryResponse
}

// 互动积分查询接口 成功返回结果
type TaobaoJstInteractivePointQueryResponse struct {
    XMLName xml.Name `xml:"jst_interactive_point_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 调用是否成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    // 查询结果
    Data   *InteractivePointQueryResponse `json:"data,omitempty" xml:"data,omitempty"`
}
