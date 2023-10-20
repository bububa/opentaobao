package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterFulfiltaskInsuranceActionAPIRequest 供应链保险链路动作 API请求
// tmall.servicecenter.fulfiltask.insurance.action
//
// 服务供应链履约链路 保险类业务履约接口
type TmallServicecenterFulfiltaskInsuranceActionAPIRequest struct {
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

// NewTmallServicecenterFulfiltaskInsuranceActionRequest 初始化TmallServicecenterFulfiltaskInsuranceActionAPIRequest对象
func NewTmallServicecenterFulfiltaskInsuranceActionRequest() *TmallServicecenterFulfiltaskInsuranceActionAPIRequest {
	return &TmallServicecenterFulfiltaskInsuranceActionAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterFulfiltaskInsuranceActionAPIRequest) Reset() {
	r._taskAction = ""
	r._taskContextData = ""
	r._outerId = ""
	r._workcardId = 0
	r._fulfilTaskId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterFulfiltaskInsuranceActionAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.fulfiltask.insurance.action"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterFulfiltaskInsuranceActionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterFulfiltaskInsuranceActionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTaskAction is TaskAction Setter
// 履约动作，取值：send_goods, 上门取件;supplier_signed, 服务商签收;evaluate_report, 鉴定报告;insurance_claims, 保险理赔;send_back, 寄回;repair_finish, 维修完成;
func (r *TmallServicecenterFulfiltaskInsuranceActionAPIRequest) SetTaskAction(_taskAction string) error {
	r._taskAction = _taskAction
	r.Set("task_action", _taskAction)
	return nil
}

// GetTaskAction TaskAction Getter
func (r TmallServicecenterFulfiltaskInsuranceActionAPIRequest) GetTaskAction() string {
	return r._taskAction
}

// SetTaskContextData is TaskContextData Setter
// 履约动作数据，不同动作取值见api文档；
func (r *TmallServicecenterFulfiltaskInsuranceActionAPIRequest) SetTaskContextData(_taskContextData string) error {
	r._taskContextData = _taskContextData
	r.Set("task_context_data", _taskContextData)
	return nil
}

// GetTaskContextData TaskContextData Getter
func (r TmallServicecenterFulfiltaskInsuranceActionAPIRequest) GetTaskContextData() string {
	return r._taskContextData
}

// SetOuterId is OuterId Setter
// 外部单号id
func (r *TmallServicecenterFulfiltaskInsuranceActionAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TmallServicecenterFulfiltaskInsuranceActionAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *TmallServicecenterFulfiltaskInsuranceActionAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallServicecenterFulfiltaskInsuranceActionAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// SetFulfilTaskId is FulfilTaskId Setter
// 履行单id
func (r *TmallServicecenterFulfiltaskInsuranceActionAPIRequest) SetFulfilTaskId(_fulfilTaskId int64) error {
	r._fulfilTaskId = _fulfilTaskId
	r.Set("fulfil_task_id", _fulfilTaskId)
	return nil
}

// GetFulfilTaskId FulfilTaskId Getter
func (r TmallServicecenterFulfiltaskInsuranceActionAPIRequest) GetFulfilTaskId() int64 {
	return r._fulfilTaskId
}

var poolTmallServicecenterFulfiltaskInsuranceActionAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterFulfiltaskInsuranceActionRequest()
	},
}

// GetTmallServicecenterFulfiltaskInsuranceActionRequest 从 sync.Pool 获取 TmallServicecenterFulfiltaskInsuranceActionAPIRequest
func GetTmallServicecenterFulfiltaskInsuranceActionAPIRequest() *TmallServicecenterFulfiltaskInsuranceActionAPIRequest {
	return poolTmallServicecenterFulfiltaskInsuranceActionAPIRequest.Get().(*TmallServicecenterFulfiltaskInsuranceActionAPIRequest)
}

// ReleaseTmallServicecenterFulfiltaskInsuranceActionAPIRequest 将 TmallServicecenterFulfiltaskInsuranceActionAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterFulfiltaskInsuranceActionAPIRequest(v *TmallServicecenterFulfiltaskInsuranceActionAPIRequest) {
	v.Reset()
	poolTmallServicecenterFulfiltaskInsuranceActionAPIRequest.Put(v)
}
