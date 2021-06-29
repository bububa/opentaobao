package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
配送公司信息查询接口 API返回值 
taobao.qimen.expressinfo.query

配送公司信息查询
*/
type TaobaoQimenExpressinfoQueryAPIResponse struct {
    model.CommonResponse
    TaobaoQimenExpressinfoQueryResponse
}

// 配送公司信息查询接口 成功返回结果
type TaobaoQimenExpressinfoQueryResponse struct {
    XMLName xml.Name `xml:"qimen_expressinfo_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 
    Response   *Response `json:"response,omitempty" xml:"response,omitempty"`
}
