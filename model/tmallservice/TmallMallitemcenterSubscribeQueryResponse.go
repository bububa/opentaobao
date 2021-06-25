package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
天猫服务订购信息查询接口 APIResponse
tmall.mallitemcenter.subscribe.query

查询商家服务订购信息
*/
type TmallMallitemcenterSubscribeQueryAPIResponse struct {
    model.CommonResponse
    Response *TmallMallitemcenterSubscribeQueryResponse `json:"tmall_mallitemcenter_subscribe_query_response,omitempty"`
}

type TmallMallitemcenterSubscribeQueryResponse struct {

    // 接口返回model
    Result   *TmallMallitemcenterSubscribeQueryResult `json:"result,omitempty"`

}
