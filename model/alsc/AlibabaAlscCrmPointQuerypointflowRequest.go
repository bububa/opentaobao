package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询积分流水 APIRequest
alibaba.alsc.crm.point.querypointflow

分页查询积分流水
*/
type AlibabaAlscCrmPointQuerypointflowRequest struct {
    model.Params

    // 入参
    paramPageQueryPointFlowOpenReq   *PageQueryPointFlowOpenReq 

}

func NewAlibabaAlscCrmPointQuerypointflowRequest() *AlibabaAlscCrmPointQuerypointflowRequest{
    return &AlibabaAlscCrmPointQuerypointflowRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmPointQuerypointflowRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.point.querypointflow"
}

func (r AlibabaAlscCrmPointQuerypointflowRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmPointQuerypointflowRequest) SetParamPageQueryPointFlowOpenReq(paramPageQueryPointFlowOpenReq *PageQueryPointFlowOpenReq) error {
    r.paramPageQueryPointFlowOpenReq = paramPageQueryPointFlowOpenReq
    r.Set("param_page_query_point_flow_open_req", paramPageQueryPointFlowOpenReq)
    return nil
}

func (r AlibabaAlscCrmPointQuerypointflowRequest) GetParamPageQueryPointFlowOpenReq() *PageQueryPointFlowOpenReq {
    return r.paramPageQueryPointFlowOpenReq
}

