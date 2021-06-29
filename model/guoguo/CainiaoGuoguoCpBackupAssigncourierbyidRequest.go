package guoguo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据菜鸟账号ID指派小件员 API请求
cainiao.guoguo.cp.backup.assigncourierbyid

根据菜鸟账号ID指派小件员
*/
type CainiaoGuoguoCpBackupAssigncourierbyidRequest struct {
    model.Params
    // 指派/改派原因
    _assignReason   string
    // 指派/改派原因编码
    _assignReasonCode   string
    // 任务编号
    _taskId   int64
    // 小件员菜鸟账号ID
    _accountId   int64
    // CP公司编号
    _cpCode   string
}

// 初始化CainiaoGuoguoCpBackupAssigncourierbyidRequest对象
func NewCainiaoGuoguoCpBackupAssigncourierbyidRequest() *CainiaoGuoguoCpBackupAssigncourierbyidRequest{
    return &CainiaoGuoguoCpBackupAssigncourierbyidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoGuoguoCpBackupAssigncourierbyidRequest) GetApiMethodName() string {
    return "cainiao.guoguo.cp.backup.assigncourierbyid"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoGuoguoCpBackupAssigncourierbyidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AssignReason Setter
// 指派/改派原因
func (r *CainiaoGuoguoCpBackupAssigncourierbyidRequest) SetAssignReason(_assignReason string) error {
    r._assignReason = _assignReason
    r.Set("assign_reason", _assignReason)
    return nil
}

// AssignReason Getter
func (r CainiaoGuoguoCpBackupAssigncourierbyidRequest) GetAssignReason() string {
    return r._assignReason
}
// AssignReasonCode Setter
// 指派/改派原因编码
func (r *CainiaoGuoguoCpBackupAssigncourierbyidRequest) SetAssignReasonCode(_assignReasonCode string) error {
    r._assignReasonCode = _assignReasonCode
    r.Set("assign_reason_code", _assignReasonCode)
    return nil
}

// AssignReasonCode Getter
func (r CainiaoGuoguoCpBackupAssigncourierbyidRequest) GetAssignReasonCode() string {
    return r._assignReasonCode
}
// TaskId Setter
// 任务编号
func (r *CainiaoGuoguoCpBackupAssigncourierbyidRequest) SetTaskId(_taskId int64) error {
    r._taskId = _taskId
    r.Set("task_id", _taskId)
    return nil
}

// TaskId Getter
func (r CainiaoGuoguoCpBackupAssigncourierbyidRequest) GetTaskId() int64 {
    return r._taskId
}
// AccountId Setter
// 小件员菜鸟账号ID
func (r *CainiaoGuoguoCpBackupAssigncourierbyidRequest) SetAccountId(_accountId int64) error {
    r._accountId = _accountId
    r.Set("account_id", _accountId)
    return nil
}

// AccountId Getter
func (r CainiaoGuoguoCpBackupAssigncourierbyidRequest) GetAccountId() int64 {
    return r._accountId
}
// CpCode Setter
// CP公司编号
func (r *CainiaoGuoguoCpBackupAssigncourierbyidRequest) SetCpCode(_cpCode string) error {
    r._cpCode = _cpCode
    r.Set("cp_code", _cpCode)
    return nil
}

// CpCode Getter
func (r CainiaoGuoguoCpBackupAssigncourierbyidRequest) GetCpCode() string {
    return r._cpCode
}
