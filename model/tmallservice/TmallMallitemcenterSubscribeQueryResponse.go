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
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_mallitemcenter_subscribe_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *TmallMallitemcenterSubscribeQueryResult `json:"result,omitempty" xml:"