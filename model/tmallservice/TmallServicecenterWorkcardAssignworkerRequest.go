package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商分派工人 APIRequest
tmall.servicecenter.workcard.assignworker

服务商调用该接口分派工人给具体的工单
*/
type TmallServicecenterWorkcardAssignworkerRequest struct {
    model.Params

    // 需要指派的工人id
    targetWorkerId   int64 

    // 需要指派的工人手机
    targetWorkerMobile   string 

    // 需要派工人的工单id
    workcardId   int64 

    // 需要指派的工人姓名
    targetWorkerName   string 

    // 核销单外部id
    outerId   string 

    // 不检查订单类型的原因ID由运营提供，服务商不可自由传
    stopOrderTypeCheckReason   int64 

    // 核销单id
    fulfilTaskId   int64 

    // 扩展信息
    extJson   string 

}

func NewTmallServicecenterWorkcardAssignworkerRequest() *TmallServicecenterWorkcardAssignworkerRequest{
    return &TmallServicecenterWorkcardAssignworkerRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterWorkcardAssignworkerRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.assignworker"
}

func (r TmallServicecenterWorkcardAssignworkerRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterWorkcardAssignworkerRequest) SetTargetWorkerId(targetWorkerId int64) error {
    r.targetWorkerId = targetWorkerId
    r.Set("target_worker_id", targetWorkerId)
    return nil
}

func (r TmallServicecenterWorkcardAssignworkerRequest) GetTargetWorkerId() int64 {
    return r.targetWorkerId
}

func (r *TmallServicecenterWorkcardAssignworkerRequest) SetTargetWorkerMobile(targetWorkerMobile string) error {
    r.targetWorkerMobile = targetWorkerMobile
    r.Set("target_worker_mobile", targetWorkerMobile)
    return nil
}

func (r TmallServicecenterWorkcardAssignworkerRequest) GetTargetWorkerMobile() string {
    return r.targetWorkerMobile
}

func (r *TmallServicecenterWorkcardAssignworkerRequest) SetWorkcardId(workcardId int64) error {
    r.workcardId = workcardId
    r.Set("workcard_id", workcardId)
    return nil
}

func (r TmallServicecenterWorkcardAssignworkerRequest) GetWorkcardId() int64 {
    return r.workcardId
}

func (r *TmallServicecenterWorkcardAssignworkerRequest) SetTargetWorkerName(targetWorkerName string) error {
    r.targetWorkerName = targetWorkerName
    r.Set("target_worker_name", targetWorkerName)
    return nil
}

func (r TmallServicecenterWorkcardAssignworkerRequest) GetTargetWorkerName() string {
    return r.targetWorkerName
}

func (r *TmallServicecenterWorkcardAssignworkerRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TmallServicecenterWorkcardAssignworkerRequest) GetOuterId() string {
    return r.outerId
}

func (r *TmallServicecenterWorkcardAssignworkerRequest) SetStopOrderTypeCheckReason(stopOrderTypeCheckReason int64) error {
    r.stopOrderTypeCheckReason = stopOrderTypeCheckReason
    r.Set("stop_order_type_check_reason", stopOrderTypeCheckReason)
    return nil
}

func (r TmallServicecenterWorkcardAssignworkerRequest) GetStopOrderTypeCheckReason() int64 {
    return r.stopOrderTypeCheckReason
}

func (r *TmallServicecenterWorkcardAssignworkerRequest) SetFulfilTaskId(fulfilTaskId int64) error {
    r.fulfilTaskId = fulfilTaskId
    r.Set("fulfil_task_id", fulfilTaskId)
    return nil
}

func (r TmallServicecenterWorkcardAssignworkerRequest) GetFulfilTaskId() int64 {
    return r.fulfilTaskId
}

func (r *TmallServicecenterWorkcardAssignworkerRequest) SetExtJson(extJson string) error {
    r.extJson = extJson
    r.Set("ext_json", extJson)
    return nil
}

func (r TmallServicecenterWorkcardAssignworkerRequest) GetExtJson() string {
    return r.extJson
}

