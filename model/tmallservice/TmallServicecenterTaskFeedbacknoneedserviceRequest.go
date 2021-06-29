package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商反馈无需安装工单接口 API请求
tmall.servicecenter.task.feedbacknoneedservice

服务商反馈无需安装工单接口
*/
type TmallServicecenterTaskFeedbacknoneedserviceRequest struct {
    model.Params
    // 入参对象
    _param   *SuspendServiceDo
}

// 初始化TmallServicecenterTaskFeedbacknoneedserviceRequest对象
func NewTmallServicecenterTaskFeedbacknoneedserviceRequest() *TmallServicecenterTaskFeedbacknoneedserviceRequest{
    return &TmallServicecenterTaskFeedbacknoneedserviceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterTaskFeedbacknoneedserviceRequest) GetApiMethodName() string {
    return "tmall.servicecenter.task.feedbacknoneedservice"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterTaskFeedbacknoneedserviceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参对象
func (r *TmallServicecenterTaskFeedbacknoneedserviceRequest) SetParam(_param *SuspendServiceDo) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r TmallServicecenterTaskFeedbacknoneedserviceRequest) GetParam() *SuspendServiceDo {
    return r._param
}
