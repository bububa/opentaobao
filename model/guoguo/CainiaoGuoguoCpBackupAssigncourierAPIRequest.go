package guoguo

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGuoguoCpBackupAssigncourierAPIRequest CP兜底后指定接单的小件员 API请求
// cainiao.guoguo.cp.backup.assigncourier
//
// CP兜底后指定接单的小件员；CP改派小件员
type CainiaoGuoguoCpBackupAssigncourierAPIRequest struct {
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

// NewCainiaoGuoguoCpBackupAssigncourierRequest 初始化CainiaoGuoguoCpBackupAssigncourierAPIRequest对象
func NewCainiaoGuoguoCpBackupAssigncourierRequest() *CainiaoGuoguoCpBackupAssigncourierAPIRequest {
	return &CainiaoGuoguoCpBackupAssigncourierAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoGuoguoCpBackupAssigncourierAPIRequest) Reset() {
	r._cpCode = ""
	r._cpUserId = ""
	r._lpCode = ""
	r._assignReasonCode = ""
	r._assignReason = ""
	r._mobile = ""
	r._taskId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGuoguoCpBackupAssigncourierAPIRequest) GetApiMethodName() string {
	return "cainiao.guoguo.cp.backup.assigncourier"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGuoguoCpBackupAssigncourierAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoGuoguoCpBackupAssigncourierAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCpCode is CpCode Setter
// 小件员所在公司编号
func (r *CainiaoGuoguoCpBackupAssigncourierAPIRequest) SetCpCode(_cpCode string) error {
	r._cpCode = _cpCode
	r.Set("cp_code", _cpCode)
	return nil
}

// GetCpCode CpCode Getter
func (r CainiaoGuoguoCpBackupAssigncourierAPIRequest) GetCpCode() string {
	return r._cpCode
}

// SetCpUserId is CpUserId Setter
// 小件员员工编号
func (r *CainiaoGuoguoCpBackupAssigncourierAPIRequest) SetCpUserId(_cpUserId string) error {
	r._cpUserId = _cpUserId
	r.Set("cp_user_id", _cpUserId)
	return nil
}

// GetCpUserId CpUserId Getter
func (r CainiaoGuoguoCpBackupAssigncourierAPIRequest) GetCpUserId() string {
	return r._cpUserId
}

// SetLpCode is LpCode Setter
// LP订单号
func (r *CainiaoGuoguoCpBackupAssigncourierAPIRequest) SetLpCode(_lpCode string) error {
	r._lpCode = _lpCode
	r.Set("lp_code", _lpCode)
	return nil
}

// GetLpCode LpCode Getter
func (r CainiaoGuoguoCpBackupAssigncourierAPIRequest) GetLpCode() string {
	return r._lpCode
}

// SetAssignReasonCode is AssignReasonCode Setter
// 指派/改派原因编码
func (r *CainiaoGuoguoCpBackupAssigncourierAPIRequest) SetAssignReasonCode(_assignReasonCode string) error {
	r._assignReasonCode = _assignReasonCode
	r.Set("assign_reason_code", _assignReasonCode)
	return nil
}

// GetAssignReasonCode AssignReasonCode Getter
func (r CainiaoGuoguoCpBackupAssigncourierAPIRequest) GetAssignReasonCode() string {
	return r._assignReasonCode
}

// SetAssignReason is AssignReason Setter
// 指派/改派原因
func (r *CainiaoGuoguoCpBackupAssigncourierAPIRequest) SetAssignReason(_assignReason string) error {
	r._assignReason = _assignReason
	r.Set("assign_reason", _assignReason)
	return nil
}

// GetAssignReason AssignReason Getter
func (r CainiaoGuoguoCpBackupAssigncourierAPIRequest) GetAssignReason() string {
	return r._assignReason
}

// SetMobile is Mobile Setter
// 小件员手机号
func (r *CainiaoGuoguoCpBackupAssigncourierAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r CainiaoGuoguoCpBackupAssigncourierAPIRequest) GetMobile() string {
	return r._mobile
}

// SetTaskId is TaskId Setter
// 任务ID
func (r *CainiaoGuoguoCpBackupAssigncourierAPIRequest) SetTaskId(_taskId int64) error {
	r._taskId = _taskId
	r.Set("task_id", _taskId)
	return nil
}

// GetTaskId TaskId Getter
func (r CainiaoGuoguoCpBackupAssigncourierAPIRequest) GetTaskId() int64 {
	return r._taskId
}

var poolCainiaoGuoguoCpBackupAssigncourierAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoGuoguoCpBackupAssigncourierRequest()
	},
}

// GetCainiaoGuoguoCpBackupAssigncourierRequest 从 sync.Pool 获取 CainiaoGuoguoCpBackupAssigncourierAPIRequest
func GetCainiaoGuoguoCpBackupAssigncourierAPIRequest() *CainiaoGuoguoCpBackupAssigncourierAPIRequest {
	return poolCainiaoGuoguoCpBackupAssigncourierAPIRequest.Get().(*CainiaoGuoguoCpBackupAssigncourierAPIRequest)
}

// ReleaseCainiaoGuoguoCpBackupAssigncourierAPIRequest 将 CainiaoGuoguoCpBackupAssigncourierAPIRequest 放入 sync.Pool
func ReleaseCainiaoGuoguoCpBackupAssigncourierAPIRequest(v *CainiaoGuoguoCpBackupAssigncourierAPIRequest) {
	v.Reset()
	poolCainiaoGuoguoCpBackupAssigncourierAPIRequest.Put(v)
}
