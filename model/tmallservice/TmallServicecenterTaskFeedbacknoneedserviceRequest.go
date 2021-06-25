package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商反馈无需安装工单接口 APIRequest
tmall.servicecenter.task.feedbacknoneedservice

服务商反馈无需安装工单接口
*/
type TmallServicecenterTaskFeedbacknoneedserviceRequest struct {
    model.Params

    // 入参对象
    param   *SuspendServiceDo 

}

func NewTmallServicecenterTaskFeedbacknoneedserviceRequest() *TmallServicecenterTaskFeedbacknoneedserviceRequest{
    return &TmallServicecenterTaskFeedbacknoneedserviceRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterTaskFeedbacknoneedserviceRequest) GetApiMethodName() string {
    return "tmall.servicecenter.task.feedbacknoneedservice"
}

func (r TmallServicecenterTaskFeedbacknoneedserviceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterTaskFeedbacknoneedserviceRequest) SetParam(param *SuspendServiceDo) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r TmallServicecenterTaskFeedbacknoneedserviceRequest) GetParam() *SuspendServiceDo {
    return r.param
}

