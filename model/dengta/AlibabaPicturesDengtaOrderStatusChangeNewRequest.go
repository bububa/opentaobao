package dengta

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天下秀订单状态变更通知 API请求
alibaba.pictures.dengta.order.status.change.new

天下秀订单状态变更通知
*/
type AlibabaPicturesDengtaOrderStatusChangeNewRequest struct {
    model.Params
    // 拒绝原因
    remark   string
    // 新状态
    status   int64
    // 变更时间
    changeTime   string
    // ims订单编号
    imsOrderId   string
    // task 编号
    aliTaskId   string
    // 发布内容
    taskContent   string
    // 发布图片url列表
    taskPic   string
    // 扩展字段。json结构
    extJson   string
}

// 初始化AlibabaPicturesDengtaOrderStatusChangeNewRequest对象
func NewAlibabaPicturesDengtaOrderStatusChangeNewRequest() *AlibabaPicturesDengtaOrderStatusChangeNewRequest{
    return &AlibabaPicturesDengtaOrderStatusChangeNewRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaPicturesDengtaOrderStatusChangeNewRequest) GetApiMethodName() string {
    return "alibaba.pictures.dengta.order.status.change.new"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaPicturesDengtaOrderStatusChangeNewRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Remark Setter
// 拒绝原因
func (r *AlibabaPicturesDengtaOrderStatusChangeNewRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

// Remark Getter
func (r AlibabaPicturesDengtaOrderStatusChangeNewRequest) GetRemark() string {
    return r.remark
}
// Status Setter
// 新状态
func (r *AlibabaPicturesDengtaOrderStatusChangeNewRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r AlibabaPicturesDengtaOrderStatusChangeNewRequest) GetStatus() int64 {
    return r.status
}
// ChangeTime Setter
// 变更时间
func (r *AlibabaPicturesDengtaOrderStatusChangeNewRequest) SetChangeTime(changeTime string) error {
    r.changeTime = changeTime
    r.Set("change_time", changeTime)
    return nil
}

// ChangeTime Getter
func (r AlibabaPicturesDengtaOrderStatusChangeNewRequest) GetChangeTime() string {
    return r.changeTime
}
// ImsOrderId Setter
// ims订单编号
func (r *AlibabaPicturesDengtaOrderStatusChangeNewRequest) SetImsOrderId(imsOrderId string) error {
    r.imsOrderId = imsOrderId
    r.Set("ims_order_id", imsOrderId)
    return nil
}

// ImsOrderId Getter
func (r AlibabaPicturesDengtaOrderStatusChangeNewRequest) GetImsOrderId() string {
    return r.imsOrderId
}
// AliTaskId Setter
// task 编号
func (r *AlibabaPicturesDengtaOrderStatusChangeNewRequest) SetAliTaskId(aliTaskId string) error {
    r.aliTaskId = aliTaskId
    r.Set("ali_task_id", aliTaskId)
    return nil
}

// AliTaskId Getter
func (r AlibabaPicturesDengtaOrderStatusChangeNewRequest) GetAliTaskId() string {
    return r.aliTaskId
}
// TaskContent Setter
// 发布内容
func (r *AlibabaPicturesDengtaOrderStatusChangeNewRequest) SetTaskContent(taskContent string) error {
    r.taskContent = taskContent
    r.Set("task_content", taskContent)
    return nil
}

// TaskContent Getter
func (r AlibabaPicturesDengtaOrderStatusChangeNewRequest) GetTaskContent() string {
    return r.taskContent
}
// TaskPic Setter
// 发布图片url列表
func (r *AlibabaPicturesDengtaOrderStatusChangeNewRequest) SetTaskPic(taskPic string) error {
    r.taskPic = taskPic
    r.Set("task_pic", taskPic)
    return nil
}

// TaskPic Getter
func (r AlibabaPicturesDengtaOrderStatusChangeNewRequest) GetTaskPic() string {
    return r.taskPic
}
// ExtJson Setter
// 扩展字段。json结构
func (r *AlibabaPicturesDengtaOrderStatusChangeNewRequest) SetExtJson(extJson string) error {
    r.extJson = extJson
    r.Set("ext_json", extJson)
    return nil
}

// ExtJson Getter
func (r AlibabaPicturesDengtaOrderStatusChangeNewRequest) GetExtJson() string {
    return r.extJson
}
