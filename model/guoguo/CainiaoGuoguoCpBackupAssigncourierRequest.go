package guoguo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
CP兜底后指定接单的小件员 APIRequest
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

func NewCainiaoGuoguoCpBackupAssigncourierRequest() *CainiaoGuoguoCpBackupAssigncourierRequest{
    return &CainiaoGuoguoCpBackupAssigncourierRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoGuoguoCpBackupAssigncourierRequest) GetApiMethodName() string {
    return "cainiao.guoguo.cp.backup.assigncourier"
}

func (r CainiaoGuoguoCpBackupAssigncourierRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoGuoguoCpBackupAssigncourierRequest) SetCpCode(cpCode string) error {
    r.cpCode = cpCode
    r.Set("cp_code", cpCode)
    return nil
}

func (r CainiaoGuoguoCpBackupAssigncourierRequest) GetCpCode() string {
    return r.cpCode
}

func (r *CainiaoGuoguoCpBackupAssigncourierRequest) SetCpUserId(cpUserId string) error {
    r.cpUserId = cpUserId
    r.Set("cp_user_id", cpUserId)
    return nil
}

func (r CainiaoGuoguoCpBackupAssigncourierRequest) GetCpUserId() string {
    return r.cpUserId
}

func (r *CainiaoGuoguoCpBackupAssigncourierRequest) SetLpCode(lpCode string) error {
    r.lpCode = lpCode
    r.Set("lp_code", lpCode)
    return nil
}

func (r CainiaoGuoguoCpBackupAssigncourierRequest) GetLpCode() string {
    return r.lpCode
}

func (r *CainiaoGuoguoCpBackupAssigncourierRequest) SetTaskId(taskId int64) error {
    r.taskId = taskId
    r.Set("task_id", taskId)
    return nil
}

func (r CainiaoGuoguoCpBackupAssigncourierRequest) GetTaskId() int64 {
    return r.taskId
}

func (r *CainiaoGuoguoCpBackupAssigncourierRequest) SetAssignReasonCode(assignReasonCode string) error {
    r.assignReasonCode = assignReasonCode
    r.Set("assign_reason_code", assignReasonCode)
    return nil
}

func (r CainiaoGuoguoCpBackupAssigncourierRequest) GetAssignReasonCode() string {
    return r.assignReasonCode
}

func (r *CainiaoGuoguoCpBackupAssigncourierRequest) SetAssignReason(assignReason string) error {
    r.assignReason = assignReason
    r.Set("assign_reason", assignReason)
    return nil
}

func (r CainiaoGuoguoCpBackupAssigncourierRequest) GetAssignReason() string {
    return r.assignReason
}

func (r *CainiaoGuoguoCpBackupAssigncourierRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

func (r CainiaoGuoguoCpBackupAssigncourierRequest) GetMobile() string {
    return r.mobile
}

