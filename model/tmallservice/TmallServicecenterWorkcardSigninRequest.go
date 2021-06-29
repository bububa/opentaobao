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
    outerId   string
    // 工单id
    workcardId   int64
    // 核销单id
    fulfilTaskId   int64
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
func (r *TmallServicecenterWorkcardSigninRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TmallServicecenterWorkcardSigninRequest) GetOuterId() string {
    return r.outerId
}
// WorkcardId Setter
// 工单id
func (r *TmallServicecenterWorkcardSigninRequest) SetWorkcardId(workcardId int64) error {
    r.workcardId = workcardId
    r.Set("workcard_id", workcardId)
    return nil
}

// WorkcardId Getter
func (r TmallServicecenterWorkcardSigninRequest) GetWorkcardId() int64 {
    return r.workcardId
}
// FulfilTaskId Setter
// 核销单id
func (r *TmallServicecenterWorkcardSigninRequest) SetFulfilTaskId(fulfilTaskId int64) error {
    r.fulfilTaskId = fulfilTaskId
    r.Set("fulfil_task_id", fulfilTaskId)
    return nil
}

// FulfilTaskId Getter
func (r TmallServicecenterWorkcardSigninRequest) GetFulfilTaskId() int64 {
    return r.fulfilTaskId
}
