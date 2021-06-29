package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
前后端商品映射查询接口 APIResponse
taobao.qimen.itemmapping.query

前后端商品映射查询接口
*/
type TaobaoQimenItemmappingQueryAPIResponse struct {
    model.CommonResponse
    TaobaoQimenItemmappingQueryResponse
}

type TaobaoQimenItemmappingQueryResponse struct {
    XMLName xml.Name `xml:"qimen_itemmapping_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 
    
    Response   *Response `json:"response,omitempty" xml:"response,omitempty"`

    
}
