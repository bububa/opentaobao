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
type CainiaoCbossWorkplatformOperationReplyAPIRequest struct {
    model.Params
    // 工单id
    _workOrderId   string
    // 工单任务id
    _taskId   string
    // 任务操作时间
    _actionTime   string
    // 任务操作类型
    _actionType   int64
    // 操作者userId
    _dealerUserId   string
    // 操作者联系方式
    _dealerContact   string
    // 商家工单操作回传备注
    _memo   string
    // 凭证照片地址拼接
    _attachPath   string
    // 扩展字段
    _features   string
}

// 初始化CainiaoCbossWorkplatformOperationReplyAPIRequest对象
func NewCainiaoCbossWorkplatformOperationReplyRequest() *CainiaoCbossWorkplatformOperationReplyAPIRequest{
    return &CainiaoCbossWorkplatformOperationReplyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoCbossWorkplatformOperationReplyAPIRequest) GetApiMethodName() string {
    return "cainiao.cboss.workplatform.operation.reply"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoCbossWorkplatformOperationReplyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkOrderId Setter
// 工单id
func (r *CainiaoCbossWorkplatformOperationReplyAPIRequest) SetWorkOrderId(_workOrderId string) error {
    r._workOrderId = _workOrderId
    r.Set("work_order_id", _workOrderId)
    return nil
}

// WorkOrderId Getter
func (r CainiaoCbossWorkplatformOperationReplyAPIRequest) GetWorkOrderId() string {
    return r._workOrderId
}
// TaskId Setter
// 工单任务id
func (r *CainiaoCbossWorkplatformOperationReplyAPIRequest) SetTaskId(_taskId string) error {
    r._taskId = _taskId
    r.Set("task_id", _taskId)
    return nil
}

// TaskId Getter
func (r CainiaoCbossWorkplatformOperationReplyAPIRequest) GetTaskId() string {
    return r._taskId
}
// ActionTime Setter
// 任务操作时间
func (r *CainiaoCbossWorkplatformOperationReplyAPIRequest) SetActionTime(_actionTime string) error {
    r._actionTime = _actionTime
    r.Set("action_time", _actionTime)
    return nil
}

// ActionTime Getter
func (r CainiaoCbossWorkplatformOperationReplyAPIRequest) GetActionTime() string {
    return r._actionTime
}
// ActionType Setter
// 任务操作类型
func (r *CainiaoCbossWorkplatformOperationReplyAPIRequest) SetActionType(_actionType int64) error {
    r._actionType = _actionType
    r.Set("action_type", _actionType)
    return nil
}

// ActionType Getter
func (r CainiaoCbossWorkplatformOperationReplyAPIRequest) GetActionType() int64 {
    return r._actionType
}
// DealerUserId Setter
// 操作者userId
func (r *CainiaoCbossWorkplatformOperationReplyAPIRequest) SetDealerUserId(_dealerUserId string) error {
    r._dealerUserId = _dealerUserId
    r.Set("dealer_user_id", _dealerUserId)
    return nil
}

// DealerUserId Getter
func (r CainiaoCbossWorkplatformOperationReplyAPIRequest) GetDealerUserId() string {
    return r._dealerUserId
}
// DealerContact Setter
// 操作者联系方式
func (r *CainiaoCbossWorkplatformOperationReplyAPIRequest) SetDealerContact(_dealerContact string) error {
    r._dealerContact = _dealerContact
    r.Set("dealer_contact", _dealerContact)
    return nil
}

// DealerContact Getter
func (r CainiaoCbossWorkplatformOperationReplyAPIRequest) GetDealerContact() string {
    return r._dealerContact
}
// Memo Setter
// 商家工单操作回传备注
func (r *CainiaoCbossWorkplatformOperationReplyAPIRequest) SetMemo(_memo string) error {
    r._memo = _memo
    r.Set("memo", _memo)
    return nil
}

// Memo Getter
func (r CainiaoCbossWorkplatformOperationReplyAPIRequest) GetMemo() string {
    return r._memo
}
// AttachPath Setter
// 凭证照片地址拼接
func (r *CainiaoCbossWorkplatformOperationReplyAPIRequest) SetAttachPath(_attachPath string) error {
    r._attachPath = _attachPath
    r.Set("attach_path", _attachPath)
    return nil
}

// AttachPath Getter
func (r CainiaoCbossWorkplatformOperationReplyAPIRequest) GetAttachPath() string {
    return r._attachPath
}
// Features Setter
// 扩展字段
func (r *CainiaoCbossWorkplatformOperationReplyAPIRequest) SetFeatures(_features string) error {
    r._features = _features
    r.Set("features", _features)
    return nil
}

// Features Getter
func (r CainiaoCbossWorkplatformOperationReplyAPIRequest) GetFeatures() string {
    return r._features
}
