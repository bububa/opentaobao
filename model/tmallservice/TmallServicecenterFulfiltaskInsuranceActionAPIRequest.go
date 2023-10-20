package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterfulfiltaskinsuranceactionAPIRequest 供应链保险链路动作 API请求
// tmall.servicecenter.fulfiltask.insurance.action
//
// 服务供应链履约链路 保险类业务履约接口
type TmallservicecenterfulfiltaskinsuranceactionAPIRequest struct {
	model.Params
	// 履约动作，取值：send_goods, 上门取件;supplier_signed, 服务商签收;evaluate_report, 鉴定报告;insurance_claims, 保险理赔;send_back, 寄回;repair_finish, 维修完成;
	_taskAction string
	// 履约动作数据，不同动作取值见api文档；
	_taskContextData string
	// 外部单号id
	_outerId string
	// 工单id
	_workcardId int64
	// 履行单id
	_fulfilTaskId int64
}

// NewTmallservicecenterfulfiltaskinsuranceactionRequest 初始化TmallservicecenterfulfiltaskinsuranceactionAPIRequest对象
func NewTmallservicecenterfulfiltaskinsuranceactionRequest() *TmallservicecenterfulfiltaskinsuranceactionAPIRequest {
	return &TmallservicecenterfulfiltaskinsuranceactionAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterfulfiltaskinsuranceactionAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.fulfiltask.insurance.action"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterfulfiltaskinsuranceactionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterfulfiltaskinsuranceactionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTaskAction is TaskAction Setter
// 履约动作，取值：send_goods, 上门取件;supplier_signed, 服务商签收;evaluate_report, 鉴定报告;insurance_claims, 保险理赔;send_back, 寄回;repair_finish, 维修完成;
func (r *TmallservicecenterfulfiltaskinsuranceactionAPIRequest) SetTaskAction(_taskAction string) error {
	r._taskAction = _taskAction
	r.Set("task_action", _taskAction)
	return nil
}

// GetTaskAction TaskAction Getter
func (r TmallservicecenterfulfiltaskinsuranceactionAPIRequest) GetTaskAction() string {
	return r._taskAction
}

// SetTaskContextData is TaskContextData Setter
// 履约动作数据，不同动作取值见api文档；
func (r *TmallservicecenterfulfiltaskinsuranceactionAPIRequest) SetTaskContextData(_taskContextData string) error {
	r._taskContextData = _taskContextData
	r.Set("task_context_data", _taskContextData)
	return nil
}

// GetTaskContextData TaskContextData Getter
func (r TmallservicecenterfulfiltaskinsuranceactionAPIRequest) GetTaskContextData() string {
	return r._taskContextData
}

// SetOuterId is OuterId Setter
// 外部单号id
func (r *TmallservicecenterfulfiltaskinsuranceactionAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TmallservicecenterfulfiltaskinsuranceactionAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *TmallservicecenterfulfiltaskinsuranceactionAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallservicecenterfulfiltaskinsuranceactionAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// SetFulfilTaskId is FulfilTaskId Setter
// 履行单id
func (r *TmallservicecenterfulfiltaskinsuranceactionAPIRequest) SetFulfilTaskId(_fulfilTaskId int64) error {
	r._fulfilTaskId = _fulfilTaskId
	r.Set("fulfil_task_id", _fulfilTaskId)
	return nil
}

// GetFulfilTaskId FulfilTaskId Getter
func (r TmallservicecenterfulfiltaskinsuranceactionAPIRequest) GetFulfilTaskId() int64 {
	return r._fulfilTaskId
}
