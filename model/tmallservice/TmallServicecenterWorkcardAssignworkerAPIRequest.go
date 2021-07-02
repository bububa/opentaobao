package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardAssignworkerAPIRequest 服务商分派工人 API请求
// tmall.servicecenter.workcard.assignworker
//
// 服务商调用该接口分派工人给具体的工单
type TmallServicecenterWorkcardAssignworkerAPIRequest struct {
	model.Params
	// 需要指派的工人id
	_targetWorkerId int64
	// 需要指派的工人手机
	_targetWorkerMobile string
	// 需要派工人的工单id
	_workcardId int64
	// 需要指派的工人姓名
	_targetWorkerName string
	// 核销单外部id
	_outerId string
	// 不检查订单类型的原因ID由运营提供，服务商不可自由传
	_stopOrderTypeCheckReason int64
	// 核销单id
	_fulfilTaskId int64
	// 扩展信息
	_extJson string
}

// NewTmallServicecenterWorkcardAssignworkerRequest 初始化TmallServicecenterWorkcardAssignworkerAPIRequest对象
func NewTmallServicecenterWorkcardAssignworkerRequest() *TmallServicecenterWorkcardAssignworkerAPIRequest {
	return &TmallServicecenterWorkcardAssignworkerAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardAssignworkerAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.assignworker"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardAssignworkerAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TargetWorkerId Setter
// 需要指派的工人id
func (r *TmallServicecenterWorkcardAssignworkerAPIRequest) SetTargetWorkerId(_targetWorkerId int64) error {
	r._targetWorkerId = _targetWorkerId
	r.Set("target_worker_id", _targetWorkerId)
	return nil
}

// Get TargetWorkerId Getter
func (r TmallServicecenterWorkcardAssignworkerAPIRequest) GetTargetWorkerId() int64 {
	return r._targetWorkerId
}

// Set is TargetWorkerMobile Setter
// 需要指派的工人手机
func (r *TmallServicecenterWorkcardAssignworkerAPIRequest) SetTargetWorkerMobile(_targetWorkerMobile string) error {
	r._targetWorkerMobile = _targetWorkerMobile
	r.Set("target_worker_mobile", _targetWorkerMobile)
	return nil
}

// Get TargetWorkerMobile Getter
func (r TmallServicecenterWorkcardAssignworkerAPIRequest) GetTargetWorkerMobile() string {
	return r._targetWorkerMobile
}

// Set is WorkcardId Setter
// 需要派工人的工单id
func (r *TmallServicecenterWorkcardAssignworkerAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// Get WorkcardId Getter
func (r TmallServicecenterWorkcardAssignworkerAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// Set is TargetWorkerName Setter
// 需要指派的工人姓名
func (r *TmallServicecenterWorkcardAssignworkerAPIRequest) SetTargetWorkerName(_targetWorkerName string) error {
	r._targetWorkerName = _targetWorkerName
	r.Set("target_worker_name", _targetWorkerName)
	return nil
}

// Get TargetWorkerName Getter
func (r TmallServicecenterWorkcardAssignworkerAPIRequest) GetTargetWorkerName() string {
	return r._targetWorkerName
}

// Set is OuterId Setter
// 核销单外部id
func (r *TmallServicecenterWorkcardAssignworkerAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// Get OuterId Getter
func (r TmallServicecenterWorkcardAssignworkerAPIRequest) GetOuterId() string {
	return r._outerId
}

// Set is StopOrderTypeCheckReason Setter
// 不检查订单类型的原因ID由运营提供，服务商不可自由传
func (r *TmallServicecenterWorkcardAssignworkerAPIRequest) SetStopOrderTypeCheckReason(_stopOrderTypeCheckReason int64) error {
	r._stopOrderTypeCheckReason = _stopOrderTypeCheckReason
	r.Set("stop_order_type_check_reason", _stopOrderTypeCheckReason)
	return nil
}

// Get StopOrderTypeCheckReason Getter
func (r TmallServicecenterWorkcardAssignworkerAPIRequest) GetStopOrderTypeCheckReason() int64 {
	return r._stopOrderTypeCheckReason
}

// Set is FulfilTaskId Setter
// 核销单id
func (r *TmallServicecenterWorkcardAssignworkerAPIRequest) SetFulfilTaskId(_fulfilTaskId int64) error {
	r._fulfilTaskId = _fulfilTaskId
	r.Set("fulfil_task_id", _fulfilTaskId)
	return nil
}

// Get FulfilTaskId Getter
func (r TmallServicecenterWorkcardAssignworkerAPIRequest) GetFulfilTaskId() int64 {
	return r._fulfilTaskId
}

// Set is ExtJson Setter
// 扩展信息
func (r *TmallServicecenterWorkcardAssignworkerAPIRequest) SetExtJson(_extJson string) error {
	r._extJson = _extJson
	r.Set("ext_json", _extJson)
	return nil
}

// Get ExtJson Getter
func (r TmallServicecenterWorkcardAssignworkerAPIRequest) GetExtJson() string {
	return r._extJson
}
