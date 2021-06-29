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
    assignReason   string
    // 指派/改派原因编码
    assignReasonCode   string
    // 任务编号
    taskId   int64
    // 小件员菜鸟账号ID
    accountId   int64
    // CP公司编号
    cpCode   string
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
func (r *CainiaoGuoguoCpBackupAssigncourierbyidRequest) SetAssignReason(assignReason string) error {
    r.assignReason = assignReason
    r.Set("assign_reason", assignReason)
    return nil
}

// AssignReason Getter
func (r CainiaoGuoguoCpBackupAssigncourierbyidRequest) GetAssignReason() string {
    return r.assignReason
}
// AssignReasonCode Setter
// 指派/改派原因编码
func (r *CainiaoGuoguoCpBackupAssigncourierbyidRequest) SetAssignReasonCode(assignReasonCode string) error {
    r.assignReasonCode = assignReasonCode
    r.Set("assign_reason_code", assignReasonCode)
    return nil
}

// AssignReasonCode Getter
func (r CainiaoGuoguoCpBackupAssigncourierbyidRequest) GetAssignReasonCode() string {
    return r.assignReasonCode
}
// TaskId Setter
// 任务编号
func (r *CainiaoGuoguoCpBackupAssigncourierbyidRequest) SetTaskId(taskId int64) error {
    r.taskId = taskId
    r.Set("task_id", taskId)
    return nil
}

// TaskId Getter
func (r CainiaoGuoguoCpBackupAssigncourierbyidRequest) GetTaskId() int64 {
    return r.taskId
}
// AccountId Setter
// 小件员菜鸟账号ID
func (r *CainiaoGuoguoCpBackupAssigncourierbyidRequest) SetAccountId(accountId int64) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

// AccountId Getter
func (r CainiaoGuoguoCpBackupAssigncourierbyidRequest) GetAccountId() int64 {
    return r.accountId
}
// CpCode Setter
// CP公司编号
func (r *CainiaoGuoguoCpBackupAssigncourierbyidRequest) SetCpCode(cpCode string) error {
    r.cpCode = cpCode
    r.Set("cp_code", cpCode)
    return nil
}

// CpCode Getter
func (r CainiaoGuoguoCpBackupAssigncourierbyidRequest) GetCpCode() string {
    return r.cpCode
}
