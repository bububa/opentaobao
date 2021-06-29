package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫服务订购信息查询接口 APIResponse
tmall.mallitemcenter.subscribe.query

查询商家服务订购信息
*/
type TmallMallitemcenterSubscribeQueryAPIResponse struct {
    model.CommonResponse
    TmallMallitemcenterSubscribeQueryResponse
}

type TmallMallitemcenterSubscribeQueryResponse struct {
    XMLName xml.Name `xml:"tmall_mallitemcenter_subscribe_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TmallMallitemcenterSubscribeQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
