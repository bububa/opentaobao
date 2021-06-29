package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商确认工人签到成功 API请求
tmall.servicecenter.workcard.signin

服务商确认工人签到成功。需要服务商自己保证工人是在现场服务中。否则虚假回传签到而引起的后续问题全部由服务商自己承担
*/
type TmallServicecenterWorkcardSigninRequest struct {
    model.Params
    // 核销单外部id
    _outerId   string
    // 工单id
    _workcardId   int64
    // 核销单id
    _fulfilTaskId   int64
}

// 初始化TmallServicecenterWorkcardSigninRequest对象
func NewTmallServicecenterWorkcardSigninRequest() *TmallServicecenterWorkcardSigninRequest{
    return &TmallServicecenterWorkcardSigninRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardSigninRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.signin"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardSigninRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterId Setter
// 核销单外部id
func (r *TmallServicecenterWorkcardSigninRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TmallServicecenterWorkcardSigninRequest) GetOuterId() string {
    return r._outerId
}
// WorkcardId Setter
// 工单id
func (r *TmallServicecenterWorkcardSigninRequest) SetWorkcardId(_workcardId int64) error {
    r._workcardId = _workcardId
    r.Set("workcard_id", _workcardId)
    return nil
}

// WorkcardId Getter
func (r TmallServicecenterWorkcardSigninRequest) GetWorkcardId() int64 {
    return r._workcardId
}
// FulfilTaskId Setter
// 核销单id
func (r *TmallServicecenterWorkcardSigninRequest) SetFulfilTaskId(_fulfilTaskId int64) error {
    r._fulfilTaskId = _fulfilTaskId
    r.Set("fulfil_task_id", _fulfilTaskId)
    return nil
}

// FulfilTaskId Getter
func (r TmallServicecenterWorkcardSigninRequest) GetFulfilTaskId() int64 {
    return r._fulfilTaskId
}
