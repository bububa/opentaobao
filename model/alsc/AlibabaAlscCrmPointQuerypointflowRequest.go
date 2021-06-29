package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询积分流水 API请求
alibaba.alsc.crm.point.querypointflow

分页查询积分流水
*/
type AlibabaAlscCrmPointQuerypointflowRequest struct {
    model.Params
    // 入参
    _paramPageQueryPointFlowOpenReq   *PageQueryPointFlowOpenReq
}

// 初始化AlibabaAlscCrmPointQuerypointflowRequest对象
func NewAlibabaAlscCrmPointQuerypointflowRequest() *AlibabaAlscCrmPointQuerypointflowRequest{
    return &AlibabaAlscCrmPointQuerypointflowRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmPointQuerypointflowRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.point.querypointflow"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmPointQuerypointflowRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamPageQueryPointFlowOpenReq Setter
// 入参
func (r *AlibabaAlscCrmPointQuerypointflowRequest) SetParamPageQueryPointFlowOpenReq(_paramPageQueryPointFlowOpenReq *PageQueryPointFlowOpenReq) error {
    r._paramPageQueryPointFlowOpenReq = _paramPageQueryPointFlowOpenReq
    r.Set("param_page_query_point_flow_open_req", _paramPageQueryPointFlowOpenReq)
    return nil
}

// ParamPageQueryPointFlowOpenReq Getter
func (r AlibabaAlscCrmPointQuerypointflowRequest) GetParamPageQueryPointFlowOpenReq() *PageQueryPointFlowOpenReq {
    return r._paramPageQueryPointFlowOpenReq
}
