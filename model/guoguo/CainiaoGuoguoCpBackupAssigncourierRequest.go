package guoguo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
CP兜底后指定接单的小件员 API请求
cainiao.guoguo.cp.backup.assigncourier

CP兜底后指定接单的小件员；CP改派小件员
*/
type CainiaoGuoguoCpBackupAssigncourierRequest struct {
    model.Params
    // 小件员所在公司编号
    cpCode   string
    // 小件员员工编号
    cpUserId   string
    // LP订单号
    lpCode   string
    // 任务ID
    taskId   int64
    // 指派/改派原因编码
    assignReasonCode   string
    // 指派/改派原因
    assignReason   string
    // 小件员手机号
    mobile   string
}

// 初始化CainiaoGuoguoCpBackupAssigncourierRequest对象
func NewCainiaoGuoguoCpBackupAssigncourierRequest() *CainiaoGuoguoCpBackupAssigncourierRequest{
    return &CainiaoGuoguoCpBackupAssigncourierRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoGuoguoCpBackupAssigncourierRequest) GetApiMethodName() string {
    return "cainiao.guoguo.cp.backup.assigncourier"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoGuoguoCpBackupAssigncourierRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CpCode Setter
// 小件员所在公司编号
func (r *CainiaoGuoguoCpBackupAssigncourierRequest) SetCpCode(cpCode string) error {
    r.cpCode = cpCode
    r.Set("cp_code", cpCode)
    return nil
}

// CpCode Getter
func (r CainiaoGuoguoCpBackupAssigncourierRequest) GetCpCode() string {
    return r.cpCode
}
// CpUserId Setter
// 小件员员工编号
func (r *CainiaoGuoguoCpBackupAssigncourierRequest) SetCpUserId(cpUserId string) error {
    r.cpUserId = cpUserId
    r.Set("cp_user_id", cpUserId)
    return nil
}

// CpUserId Getter
func (r CainiaoGuoguoCpBackupAssigncourierRequest) GetCpUserId() string {
    return r.cpUserId
}
// LpCode Setter
// LP订单号
func (r *CainiaoGuoguoCpBackupAssigncourierRequest) SetLpCode(lpCode string) error {
    r.lpCode = lpCode
    r.Set("lp_code", lpCode)
    return nil
}

// LpCode Getter
func (r CainiaoGuoguoCpBackupAssigncourierRequest) GetLpCode() string {
    return r.lpCode
}
// TaskId Setter
// 任务ID
func (r *CainiaoGuoguoCpBackupAssigncourierRequest) SetTaskId(taskId int64) error {
    r.taskId = taskId
    r.Set("task_id", taskId)
    return nil
}

// TaskId Getter
func (r CainiaoGuoguoCpBackupAssigncourierRequest) GetTaskId() int64 {
    return r.taskId
}
// AssignReasonCode Setter
// 指派/改派原因编码
func (r *CainiaoGuoguoCpBackupAssigncourierRequest) SetAssignReasonCode(assignReasonCode string) error {
    r.assignReasonCode = assignReasonCode
    r.Set("assign_reason_code", assignReasonCode)
    return nil
}

// AssignReasonCode Getter
func (r CainiaoGuoguoCpBackupAssigncourierRequest) GetAssignReasonCode() string {
    return r.assignReasonCode
}
// AssignReason Setter
// 指派/改派原因
func (r *CainiaoGuoguoCpBackupAssigncourierRequest) SetAssignReason(assignReason string) error {
    r.assignReason = assignReason
    r.Set("assign_reason", assignReason)
    return nil
}

// AssignReason Getter
func (r CainiaoGuoguoCpBackupAssigncourierRequest) GetAssignReason() string {
    return r.assignReason
}
// Mobile Setter
// 小件员手机号
func (r *CainiaoGuoguoCpBackupAssigncourierRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

// Mobile Getter
func (r CainiaoGuoguoCpBackupAssigncourierRequest) GetMobile() string {
    return r.mobile
}
