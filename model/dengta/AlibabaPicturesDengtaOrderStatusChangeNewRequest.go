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
    _remark   string
    // 新状态
    _status   int64
    // 变更时间
    _changeTime   string
    // ims订单编号
    _imsOrderId   string
    // task 编号
    _aliTaskId   string
    // 发布内容
    _taskContent   string
    // 发布图片url列表
    _taskPic   string
    // 扩展字段。json结构
    _extJson   string
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
func (r *AlibabaPicturesDengtaOrderStatusChangeNewRequest) SetRemark(_remark string) error {
    r._remark = _remark
    r.Set("remark", _remark)
    return nil
}

// Remark Getter
func (r AlibabaPicturesDengtaOrderStatusChangeNewRequest) GetRemark() string {
    return r._remark
}
// Status Setter
// 新状态
func (r *AlibabaPicturesDengtaOrderStatusChangeNewRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaPicturesDengtaOrderStatusChangeNewRequest) GetStatus() int64 {
    return r._status
}
// ChangeTime Setter
// 变更时间
func (r *AlibabaPicturesDengtaOrderStatusChangeNewRequest) SetChangeTime(_changeTime string) error {
    r._changeTime = _changeTime
    r.Set("change_time", _changeTime)
    return nil
}

// ChangeTime Getter
func (r AlibabaPicturesDengtaOrderStatusChangeNewRequest) GetChangeTime() string {
    return r._changeTime
}
// ImsOrderId Setter
// ims订单编号
func (r *AlibabaPicturesDengtaOrderStatusChangeNewRequest) SetImsOrderId(_imsOrderId string) error {
    r._imsOrderId = _imsOrderId
    r.Set("ims_order_id", _imsOrderId)
    return nil
}

// ImsOrderId Getter
func (r AlibabaPicturesDengtaOrderStatusChangeNewRequest) GetImsOrderId() string {
    return r._imsOrderId
}
// AliTaskId Setter
// task 编号
func (r *AlibabaPicturesDengtaOrderStatusChangeNewRequest) SetAliTaskId(_aliTaskId string) error {
    r._aliTaskId = _aliTaskId
    r.Set("ali_task_id", _aliTaskId)
    return nil
}

// AliTaskId Getter
func (r AlibabaPicturesDengtaOrderStatusChangeNewRequest) GetAliTaskId() string {
    return r._aliTaskId
}
// TaskContent Setter
// 发布内容
func (r *AlibabaPicturesDengtaOrderStatusChangeNewRequest) SetTaskContent(_taskContent string) error {
    r._taskContent = _taskContent
    r.Set("task_content", _taskContent)
    return nil
}

// TaskContent Getter
func (r AlibabaPicturesDengtaOrderStatusChangeNewRequest) GetTaskContent() string {
    return r._taskContent
}
// TaskPic Setter
// 发布图片url列表
func (r *AlibabaPicturesDengtaOrderStatusChangeNewRequest) SetTaskPic(_taskPic string) error {
    r._taskPic = _taskPic
    r.Set("task_pic", _taskPic)
    return nil
}

// TaskPic Getter
func (r AlibabaPicturesDengtaOrderStatusChangeNewRequest) GetTaskPic() string {
    return r._taskPic
}
// ExtJson Setter
// 扩展字段。json结构
func (r *AlibabaPicturesDengtaOrderStatusChangeNewRequest) SetExtJson(_extJson string) error {
    r._extJson = _extJson
    r.Set("ext_json", _extJson)
    return nil
}

// ExtJson Getter
func (r AlibabaPicturesDengtaOrderStatusChangeNewRequest) GetExtJson() string {
    return r._extJson
}
