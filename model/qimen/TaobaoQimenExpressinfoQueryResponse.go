package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
配送公司信息查询接口 APIResponse
taobao.qimen.expressinfo.query

配送公司信息查询
*/
type TaobaoQimenExpressinfoQueryAPIResponse struct {
    model.CommonResponse
    TaobaoQimenExpressinfoQueryResponse
}

type TaobaoQimenExpressinfoQueryResponse struct {
    XMLName xml.Name `xml:"qimen_expressinfo_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 
    
    Response   *Response `json:"response,omitempty" xml:"response,omitempty"`

    
}
