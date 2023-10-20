package guoguo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoguoguocpbackupassigncourierbyidAPIRequest 根据菜鸟账号ID指派小件员 API请求
// cainiao.guoguo.cp.backup.assigncourierbyid
//
// 根据菜鸟账号ID指派小件员
type CainiaoguoguocpbackupassigncourierbyidAPIRequest struct {
	model.Params
	// 指派/改派原因
	_assignReason string
	// 指派/改派原因编码
	_assignReasonCode string
	// CP公司编号
	_cpCode string
	// 任务编号
	_taskId int64
	// 小件员菜鸟账号ID
	_accountId int64
}

// NewCainiaoguoguocpbackupassigncourierbyidRequest 初始化CainiaoguoguocpbackupassigncourierbyidAPIRequest对象
func NewCainiaoguoguocpbackupassigncourierbyidRequest() *CainiaoguoguocpbackupassigncourierbyidAPIRequest {
	return &CainiaoguoguocpbackupassigncourierbyidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoguoguocpbackupassigncourierbyidAPIRequest) GetApiMethodName() string {
	return "cainiao.guoguo.cp.backup.assigncourierbyid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoguoguocpbackupassigncourierbyidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoguoguocpbackupassigncourierbyidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAssignReason is AssignReason Setter
// 指派/改派原因
func (r *CainiaoguoguocpbackupassigncourierbyidAPIRequest) SetAssignReason(_assignReason string) error {
	r._assignReason = _assignReason
	r.Set("assign_reason", _assignReason)
	return nil
}

// GetAssignReason AssignReason Getter
func (r CainiaoguoguocpbackupassigncourierbyidAPIRequest) GetAssignReason() string {
	return r._assignReason
}

// SetAssignReasonCode is AssignReasonCode Setter
// 指派/改派原因编码
func (r *CainiaoguoguocpbackupassigncourierbyidAPIRequest) SetAssignReasonCode(_assignReasonCode string) error {
	r._assignReasonCode = _assignReasonCode
	r.Set("assign_reason_code", _assignReasonCode)
	return nil
}

// GetAssignReasonCode AssignReasonCode Getter
func (r CainiaoguoguocpbackupassigncourierbyidAPIRequest) GetAssignReasonCode() string {
	return r._assignReasonCode
}

// SetCpCode is CpCode Setter
// CP公司编号
func (r *CainiaoguoguocpbackupassigncourierbyidAPIRequest) SetCpCode(_cpCode string) error {
	r._cpCode = _cpCode
	r.Set("cp_code", _cpCode)
	return nil
}

// GetCpCode CpCode Getter
func (r CainiaoguoguocpbackupassigncourierbyidAPIRequest) GetCpCode() string {
	return r._cpCode
}

// SetTaskId is TaskId Setter
// 任务编号
func (r *CainiaoguoguocpbackupassigncourierbyidAPIRequest) SetTaskId(_taskId int64) error {
	r._taskId = _taskId
	r.Set("task_id", _taskId)
	return nil
}

// GetTaskId TaskId Getter
func (r CainiaoguoguocpbackupassigncourierbyidAPIRequest) GetTaskId() int64 {
	return r._taskId
}

// SetAccountId is AccountId Setter
// 小件员菜鸟账号ID
func (r *CainiaoguoguocpbackupassigncourierbyidAPIRequest) SetAccountId(_accountId int64) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r CainiaoguoguocpbackupassigncourierbyidAPIRequest) GetAccountId() int64 {
	return r._accountId
}
