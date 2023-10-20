package guoguo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoguoguocpbackupassigncourierAPIRequest CP兜底后指定接单的小件员 API请求
// cainiao.guoguo.cp.backup.assigncourier
//
// CP兜底后指定接单的小件员；CP改派小件员
type CainiaoguoguocpbackupassigncourierAPIRequest struct {
	model.Params
	// 小件员所在公司编号
	_cpCode string
	// 小件员员工编号
	_cpUserId string
	// LP订单号
	_lpCode string
	// 指派/改派原因编码
	_assignReasonCode string
	// 指派/改派原因
	_assignReason string
	// 小件员手机号
	_mobile string
	// 任务ID
	_taskId int64
}

// NewCainiaoguoguocpbackupassigncourierRequest 初始化CainiaoguoguocpbackupassigncourierAPIRequest对象
func NewCainiaoguoguocpbackupassigncourierRequest() *CainiaoguoguocpbackupassigncourierAPIRequest {
	return &CainiaoguoguocpbackupassigncourierAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoguoguocpbackupassigncourierAPIRequest) GetApiMethodName() string {
	return "cainiao.guoguo.cp.backup.assigncourier"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoguoguocpbackupassigncourierAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoguoguocpbackupassigncourierAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCpCode is CpCode Setter
// 小件员所在公司编号
func (r *CainiaoguoguocpbackupassigncourierAPIRequest) SetCpCode(_cpCode string) error {
	r._cpCode = _cpCode
	r.Set("cp_code", _cpCode)
	return nil
}

// GetCpCode CpCode Getter
func (r CainiaoguoguocpbackupassigncourierAPIRequest) GetCpCode() string {
	return r._cpCode
}

// SetCpUserId is CpUserId Setter
// 小件员员工编号
func (r *CainiaoguoguocpbackupassigncourierAPIRequest) SetCpUserId(_cpUserId string) error {
	r._cpUserId = _cpUserId
	r.Set("cp_user_id", _cpUserId)
	return nil
}

// GetCpUserId CpUserId Getter
func (r CainiaoguoguocpbackupassigncourierAPIRequest) GetCpUserId() string {
	return r._cpUserId
}

// SetLpCode is LpCode Setter
// LP订单号
func (r *CainiaoguoguocpbackupassigncourierAPIRequest) SetLpCode(_lpCode string) error {
	r._lpCode = _lpCode
	r.Set("lp_code", _lpCode)
	return nil
}

// GetLpCode LpCode Getter
func (r CainiaoguoguocpbackupassigncourierAPIRequest) GetLpCode() string {
	return r._lpCode
}

// SetAssignReasonCode is AssignReasonCode Setter
// 指派/改派原因编码
func (r *CainiaoguoguocpbackupassigncourierAPIRequest) SetAssignReasonCode(_assignReasonCode string) error {
	r._assignReasonCode = _assignReasonCode
	r.Set("assign_reason_code", _assignReasonCode)
	return nil
}

// GetAssignReasonCode AssignReasonCode Getter
func (r CainiaoguoguocpbackupassigncourierAPIRequest) GetAssignReasonCode() string {
	return r._assignReasonCode
}

// SetAssignReason is AssignReason Setter
// 指派/改派原因
func (r *CainiaoguoguocpbackupassigncourierAPIRequest) SetAssignReason(_assignReason string) error {
	r._assignReason = _assignReason
	r.Set("assign_reason", _assignReason)
	return nil
}

// GetAssignReason AssignReason Getter
func (r CainiaoguoguocpbackupassigncourierAPIRequest) GetAssignReason() string {
	return r._assignReason
}

// SetMobile is Mobile Setter
// 小件员手机号
func (r *CainiaoguoguocpbackupassigncourierAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r CainiaoguoguocpbackupassigncourierAPIRequest) GetMobile() string {
	return r._mobile
}

// SetTaskId is TaskId Setter
// 任务ID
func (r *CainiaoguoguocpbackupassigncourierAPIRequest) SetTaskId(_taskId int64) error {
	r._taskId = _taskId
	r.Set("task_id", _taskId)
	return nil
}

// GetTaskId TaskId Getter
func (r CainiaoguoguocpbackupassigncourierAPIRequest) GetTaskId() int64 {
	return r._taskId
}
