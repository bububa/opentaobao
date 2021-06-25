package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟工单操作回传 APIRequest
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

func NewCainiaoCbossWorkplatformOperationReplyRequest() *CainiaoCbossWorkplatformOperationReplyRequest{
    return &CainiaoCbossWorkplatformOperationReplyRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoCbossWorkplatformOperationReplyRequest) GetApiMethodName() string {
    return "cainiao.cboss.workplatform.operation.reply"
}

func (r CainiaoCbossWorkplatformOperationReplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoCbossWorkplatformOperationReplyRequest) SetWorkOrderId(workOrderId string) error {
    r.workOrderId = workOrderId
    r.Set("work_order_id", workOrderId)
    return nil
}

func (r CainiaoCbossWorkplatformOperationReplyRequest) GetWorkOrderId() string {
    return r.workOrderId
}

func (r *CainiaoCbossWorkplatformOperationReplyRequest) SetTaskId(taskId string) error {
    r.taskId = taskId
    r.Set("task_id", taskId)
    return nil
}

func (r CainiaoCbossWorkplatformOperationReplyRequest) GetTaskId() string {
    return r.taskId
}

func (r *CainiaoCbossWorkplatformOperationReplyRequest) SetActionTime(actionTime string) error {
    r.actionTime = actionTime
    r.Set("action_time", actionTime)
    return nil
}

func (r CainiaoCbossWorkplatformOperationReplyRequest) GetActionTime() string {
    return r.actionTime
}

func (r *CainiaoCbossWorkplatformOperationReplyRequest) SetActionType(actionType int64) error {
    r.actionType = actionType
    r.Set("action_type", actionType)
    return nil
}

func (r CainiaoCbossWorkplatformOperationReplyRequest) GetActionType() int64 {
    return r.actionType
}

func (r *CainiaoCbossWorkplatformOperationReplyRequest) SetDealerUserId(dealerUserId string) error {
    r.dealerUserId = dealerUserId
    r.Set("dealer_user_id", dealerUserId)
    return nil
}

func (r CainiaoCbossWorkplatformOperationReplyRequest) GetDealerUserId() string {
    return r.dealerUserId
}

func (r *CainiaoCbossWorkplatformOperationReplyRequest) SetDealerContact(dealerContact string) error {
    r.dealerContact = dealerContact
    r.Set("dealer_contact", dealerContact)
    return nil
}

func (r CainiaoCbossWorkplatformOperationReplyRequest) GetDealerContact() string {
    return r.dealerContact
}

func (r *CainiaoCbossWorkplatformOperationReplyRequest) SetMemo(memo string) error {
    r.memo = memo
    r.Set("memo", memo)
    return nil
}

func (r CainiaoCbossWorkplatformOperationReplyRequest) GetMemo() string {
    return r.memo
}

func (r *CainiaoCbossWorkplatformOperationReplyRequest) SetAttachPath(attachPath string) error {
    r.attachPath = attachPath
    r.Set("attach_path", attachPath)
    return nil
}

func (r CainiaoCbossWorkplatformOperationReplyRequest) GetAttachPath() string {
    return r.attachPath
}

func (r *CainiaoCbossWorkplatformOperationReplyRequest) SetFeatures(features string) error {
    r.features = features
    r.Set("features", features)
    return nil
}

func (r CainiaoCbossWorkplatformOperationReplyRequest) GetFeatures() string {
    return r.features
}

