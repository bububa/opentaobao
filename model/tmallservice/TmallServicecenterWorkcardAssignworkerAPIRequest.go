package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardassignworkerAPIRequest 服务商分派工人 API请求
// tmall.servicecenter.workcard.assignworker
//
// 服务商调用该接口分派工人给具体的工单
type TmallservicecenterworkcardassignworkerAPIRequest struct {
	model.Params
	// 核销单外部id
	_outerId string
	// 需要指派的工人手机
	_targetWorkerMobile string
	// 需要指派的工人姓名
	_targetWorkerName string
	// 扩展信息
	_extJson string
	// 需要指派的工人id
	_targetWorkerId int64
	// 需要派工人的工单id
	_workcardId int64
	// 不检查订单类型的原因ID由运营提供，服务商不可自由传
	_stopOrderTypeCheckReason int64
	// 核销单id
	_fulfilTaskId int64
}

// NewTmallservicecenterworkcardassignworkerRequest 初始化TmallservicecenterworkcardassignworkerAPIRequest对象
func NewTmallservicecenterworkcardassignworkerRequest() *TmallservicecenterworkcardassignworkerAPIRequest {
	return &TmallservicecenterworkcardassignworkerAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterworkcardassignworkerAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.assignworker"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterworkcardassignworkerAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterworkcardassignworkerAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 核销单外部id
func (r *TmallservicecenterworkcardassignworkerAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TmallservicecenterworkcardassignworkerAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetTargetWorkerMobile is TargetWorkerMobile Setter
// 需要指派的工人手机
func (r *TmallservicecenterworkcardassignworkerAPIRequest) SetTargetWorkerMobile(_targetWorkerMobile string) error {
	r._targetWorkerMobile = _targetWorkerMobile
	r.Set("target_worker_mobile", _targetWorkerMobile)
	return nil
}

// GetTargetWorkerMobile TargetWorkerMobile Getter
func (r TmallservicecenterworkcardassignworkerAPIRequest) GetTargetWorkerMobile() string {
	return r._targetWorkerMobile
}

// SetTargetWorkerName is TargetWorkerName Setter
// 需要指派的工人姓名
func (r *TmallservicecenterworkcardassignworkerAPIRequest) SetTargetWorkerName(_targetWorkerName string) error {
	r._targetWorkerName = _targetWorkerName
	r.Set("target_worker_name", _targetWorkerName)
	return nil
}

// GetTargetWorkerName TargetWorkerName Getter
func (r TmallservicecenterworkcardassignworkerAPIRequest) GetTargetWorkerName() string {
	return r._targetWorkerName
}

// SetExtJson is ExtJson Setter
// 扩展信息
func (r *TmallservicecenterworkcardassignworkerAPIRequest) SetExtJson(_extJson string) error {
	r._extJson = _extJson
	r.Set("ext_json", _extJson)
	return nil
}

// GetExtJson ExtJson Getter
func (r TmallservicecenterworkcardassignworkerAPIRequest) GetExtJson() string {
	return r._extJson
}

// SetTargetWorkerId is TargetWorkerId Setter
// 需要指派的工人id
func (r *TmallservicecenterworkcardassignworkerAPIRequest) SetTargetWorkerId(_targetWorkerId int64) error {
	r._targetWorkerId = _targetWorkerId
	r.Set("target_worker_id", _targetWorkerId)
	return nil
}

// GetTargetWorkerId TargetWorkerId Getter
func (r TmallservicecenterworkcardassignworkerAPIRequest) GetTargetWorkerId() int64 {
	return r._targetWorkerId
}

// SetWorkcardId is WorkcardId Setter
// 需要派工人的工单id
func (r *TmallservicecenterworkcardassignworkerAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallservicecenterworkcardassignworkerAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// SetStopOrderTypeCheckReason is StopOrderTypeCheckReason Setter
// 不检查订单类型的原因ID由运营提供，服务商不可自由传
func (r *TmallservicecenterworkcardassignworkerAPIRequest) SetStopOrderTypeCheckReason(_stopOrderTypeCheckReason int64) error {
	r._stopOrderTypeCheckReason = _stopOrderTypeCheckReason
	r.Set("stop_order_type_check_reason", _stopOrderTypeCheckReason)
	return nil
}

// GetStopOrderTypeCheckReason StopOrderTypeCheckReason Getter
func (r TmallservicecenterworkcardassignworkerAPIRequest) GetStopOrderTypeCheckReason() int64 {
	return r._stopOrderTypeCheckReason
}

// SetFulfilTaskId is FulfilTaskId Setter
// 核销单id
func (r *TmallservicecenterworkcardassignworkerAPIRequest) SetFulfilTaskId(_fulfilTaskId int64) error {
	r._fulfilTaskId = _fulfilTaskId
	r.Set("fulfil_task_id", _fulfilTaskId)
	return nil
}

// GetFulfilTaskId FulfilTaskId Getter
func (r TmallservicecenterworkcardassignworkerAPIRequest) GetFulfilTaskId() int64 {
	return r._fulfilTaskId
}
