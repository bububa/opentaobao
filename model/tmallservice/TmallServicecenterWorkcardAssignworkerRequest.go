package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商分派工人 API请求
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

// 初始化TmallServicecenterWorkcardAssignworkerRequest对象
func NewTmallServicecenterWorkcardAssignworkerRequest() *TmallServicecenterWorkcardAssignworkerRequest{
    return &TmallServicecenterWorkcardAssignworkerRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardAssignworkerRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.assignworker"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardAssignworkerRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TargetWorkerId Setter
// 需要指派的工人id
func (r *TmallServicecenterWorkcardAssignworkerRequest) SetTargetWorkerId(targetWorkerId int64) error {
    r.targetWorkerId = targetWorkerId
    r.Set("target_worker_id", targetWorkerId)
    return nil
}

// TargetWorkerId Getter
func (r TmallServicecenterWorkcardAssignworkerRequest) GetTargetWorkerId() int64 {
    return r.targetWorkerId
}
// TargetWorkerMobile Setter
// 需要指派的工人手机
func (r *TmallServicecenterWorkcardAssignworkerRequest) SetTargetWorkerMobile(targetWorkerMobile string) error {
    r.targetWorkerMobile = targetWorkerMobile
    r.Set("target_worker_mobile", targetWorkerMobile)
    return nil
}

// TargetWorkerMobile Getter
func (r TmallServicecenterWorkcardAssignworkerRequest) GetTargetWorkerMobile() string {
    return r.targetWorkerMobile
}
// WorkcardId Setter
// 需要派工人的工单id
func (r *TmallServicecenterWorkcardAssignworkerRequest) SetWorkcardId(workcardId int64) error {
    r.workcardId = workcardId
    r.Set("workcard_id", workcardId)
    return nil
}

// WorkcardId Getter
func (r TmallServicecenterWorkcardAssignworkerRequest) GetWorkcardId() int64 {
    return r.workcardId
}
// TargetWorkerName Setter
// 需要指派的工人姓名
func (r *TmallServicecenterWorkcardAssignworkerRequest) SetTargetWorkerName(targetWorkerName string) error {
    r.targetWorkerName = targetWorkerName
    r.Set("target_worker_name", targetWorkerName)
    return nil
}

// TargetWorkerName Getter
func (r TmallServicecenterWorkcardAssignworkerRequest) GetTargetWorkerName() string {
    return r.targetWorkerName
}
// OuterId Setter
// 核销单外部id
func (r *TmallServicecenterWorkcardAssignworkerRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TmallServicecenterWorkcardAssignworkerRequest) GetOuterId() string {
    return r.outerId
}
// StopOrderTypeCheckReason Setter
// 不检查订单类型的原因ID由运营提供，服务商不可自由传
func (r *TmallServicecenterWorkcardAssignworkerRequest) SetStopOrderTypeCheckReason(stopOrderTypeCheckReason int64) error {
    r.stopOrderTypeCheckReason = stopOrderTypeCheckReason
    r.Set("stop_order_type_check_reason", stopOrderTypeCheckReason)
    return nil
}

// StopOrderTypeCheckReason Getter
func (r TmallServicecenterWorkcardAssignworkerRequest) GetStopOrderTypeCheckReason() int64 {
    return r.stopOrderTypeCheckReason
}
// FulfilTaskId Setter
// 核销单id
func (r *TmallServicecenterWorkcardAssignworkerRequest) SetFulfilTaskId(fulfilTaskId int64) error {
    r.fulfilTaskId = fulfilTaskId
    r.Set("fulfil_task_id", fulfilTaskId)
    return nil
}

// FulfilTaskId Getter
func (r TmallServicecenterWorkcardAssignworkerRequest) GetFulfilTaskId() int64 {
    return r.fulfilTaskId
}
// ExtJson Setter
// 扩展信息
func (r *TmallServicecenterWorkcardAssignworkerRequest) SetExtJson(extJson string) error {
    r.extJson = extJson
    r.Set("ext_json", extJson)
    return nil
}

// ExtJson Getter
func (r TmallServicecenterWorkcardAssignworkerRequest) GetExtJson() string {
    return r.extJson
}
