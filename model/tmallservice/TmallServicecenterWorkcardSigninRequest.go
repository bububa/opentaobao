package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商确认工人签到成功 APIRequest
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

func NewTmallServicecenterWorkcardSigninRequest() *TmallServicecenterWorkcardSigninRequest{
    return &TmallServicecenterWorkcardSigninRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterWorkcardSigninRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.signin"
}

func (r TmallServicecenterWorkcardSigninRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterWorkcardSigninRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TmallServicecenterWorkcardSigninRequest) GetOuterId() string {
    return r.outerId
}

func (r *TmallServicecenterWorkcardSigninRequest) SetWorkcardId(workcardId int64) error {
    r.workcardId = workcardId
    r.Set("workcard_id", workcardId)
    return nil
}

func (r TmallServicecenterWorkcardSigninRequest) GetWorkcardId() int64 {
    return r.workcardId
}

func (r *TmallServicecenterWorkcardSigninRequest) SetFulfilTaskId(fulfilTaskId int64) error {
    r.fulfilTaskId = fulfilTaskId
    r.Set("fulfil_task_id", fulfilTaskId)
    return nil
}

func (r TmallServicecenterWorkcardSigninRequest) GetFulfilTaskId() int64 {
    return r.fulfilTaskId
}

