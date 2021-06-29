package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟工单操作回传 API请求
cainiao.cboss.workplatform.operation.reply

菜鸟工单进度下发接口，目前调用者ISV
*/
type CainiaoCbossWorkplatformOperationReplyRequest struct {
    model.Params
    // 工单id
    workOrderId   string
    // 工单任务id
    taskId   string
    // 任务操作时间
    actionTime   string
    // 任务操作类型
    actionType   int64
    // 操作者userId
    dealerUserId   string
    // 操作者联系方式
    dealerContact   string
    // 商家工单操作回传备注
    memo   string
    // 凭证照片地址拼接
    attachPath   string
    // 扩展字段
    features   string
}

// 初始化CainiaoCbossWorkplatformOperationReplyRequest对象
func NewCainiaoCbossWorkplatformOperationReplyRequest() *CainiaoCbossWorkplatformOperationReplyRequest{
    return &CainiaoCbossWorkplatformOperationReplyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoCbossWorkplatformOperationReplyRequest) GetApiMethodName() string {
    return "cainiao.cboss.workplatform.operation.reply"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoCbossWorkplatformOperationReplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkOrderId Setter
// 工单id
func (r *CainiaoCbossWorkplatformOperationReplyRequest) SetWorkOrderId(workOrderId string) error {
    r.workOrderId = workOrderId
    r.Set("work_order_id", workOrderId)
    return nil
}

// WorkOrderId Getter
func (r CainiaoCbossWorkplatformOperationReplyRequest) GetWorkOrderId() string {
    return r.workOrderId
}
// TaskId Setter
// 工单任务id
func (r *CainiaoCbossWorkplatformOperationReplyRequest) SetTaskId(taskId string) error {
    r.taskId = taskId
    r.Set("task_id", taskId)
    return nil
}

// TaskId Getter
func (r CainiaoCbossWorkplatformOperationReplyRequest) GetTaskId() string {
    return r.taskId
}
// ActionTime Setter
// 任务操作时间
func (r *CainiaoCbossWorkplatformOperationReplyRequest) SetActionTime(actionTime string) error {
    r.actionTime = actionTime
    r.Set("action_time", actionTime)
    return nil
}

// ActionTime Getter
func (r CainiaoCbossWorkplatformOperationReplyRequest) GetActionTime() string {
    return r.actionTime
}
// ActionType Setter
// 任务操作类型
func (r *CainiaoCbossWorkplatformOperationReplyRequest) SetActionType(actionType int64) error {
    r.actionType = actionType
    r.Set("action_type", actionType)
    return nil
}

// ActionType Getter
func (r CainiaoCbossWorkplatformOperationReplyRequest) GetActionType() int64 {
    return r.actionType
}
// DealerUserId Setter
// 操作者userId
func (r *CainiaoCbossWorkplatformOperationReplyRequest) SetDealerUserId(dealerUserId string) error {
    r.dealerUserId = dealerUserId
    r.Set("dealer_user_id", dealerUserId)
    return nil
}

// DealerUserId Getter
func (r CainiaoCbossWorkplatformOperationReplyRequest) GetDealerUserId() string {
    return r.dealerUserId
}
// DealerContact Setter
// 操作者联系方式
func (r *CainiaoCbossWorkplatformOperationReplyRequest) SetDealerContact(dealerContact string) error {
    r.dealerContact = dealerContact
    r.Set("dealer_contact", dealerContact)
    return nil
}

// DealerContact Getter
func (r CainiaoCbossWorkplatformOperationReplyRequest) GetDealerContact() string {
    return r.dealerContact
}
// Memo Setter
// 商家工单操作回传备注
func (r *CainiaoCbossWorkplatformOperationReplyRequest) SetMemo(memo string) error {
    r.memo = memo
    r.Set("memo", memo)
    return nil
}

// Memo Getter
func (r CainiaoCbossWorkplatformOperationReplyRequest) GetMemo() string {
    return r.memo
}
// AttachPath Setter
// 凭证照片地址拼接
func (r *CainiaoCbossWorkplatformOperationReplyRequest) SetAttachPath(attachPath string) error {
    r.attachPath = attachPath
    r.Set("attach_path", attachPath)
    return nil
}

// AttachPath Getter
func (r CainiaoCbossWorkplatformOperationReplyRequest) GetAttachPath() string {
    return r.attachPath
}
// Features Setter
// 扩展字段
func (r *CainiaoCbossWorkplatformOperationReplyRequest) SetFeatures(features string) error {
    r.features = features
    r.Set("features", features)
    return nil
}

// Features Getter
func (r CainiaoCbossWorkplatformOperationReplyRequest) GetFeatures() string {
    return r.features
}
