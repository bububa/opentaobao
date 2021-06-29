package jstinteractive

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
互动积分查询接口 APIResponse
taobao.jst.interactive.point.query

查询用户的互动积分
*/
type TaobaoJstInteractivePointQueryAPIResponse struct {
    model.CommonResponse
    TaobaoJstInteractivePointQueryResponse
}

type TaobaoJstInteractivePointQueryResponse struct {
    XMLName xml.Name `xml:"jst_interactive_point_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 调用是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 查询结果
    
    Data   *InteractivePointQueryResponse `json:"data,omitempty" xml:"data,omitempty"`

    
}
