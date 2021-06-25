package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务单列表查询 APIResponse
tmall.servicecenter.spserviceorder.query

查询服务单列表
*/
type TmallServicecenterSpserviceorderQueryAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterSpserviceorderQueryResponse `json:"tmall_servicecenter_spserviceorder_query_response,omitempty"`
}

type TmallServicecenterSpserviceorderQueryResponse struct {

    // 返回参数
    Result   *FulfilplatformResult `json:"result,omitempty"`

}
