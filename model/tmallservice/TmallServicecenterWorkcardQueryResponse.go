package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
工单查询接口 APIResponse
tmall.servicecenter.workcard.query

工单查询接口
*/
type TmallServicecenterWorkcardQueryAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterWorkcardQueryResponse `json:"tmall_servicecenter_workcard_query_response,omitempty"`
}

type TmallServicecenterWorkcardQueryResponse struct {

    // 请求结果
    Result   *TmallServicecenterWorkcardQueryResult `json:"result,omitempty"`

}