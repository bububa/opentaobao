package dengta

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天下秀订单状态变更通知 API请求
alibaba.pictures.dengta.order.status.change

天下秀订单状态变更通知
*/
type AlibabaPicturesDengtaOrderStatusChangeRequest struct {
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
}

// 初始化AlibabaPicturesDengtaOrderStatusChangeRequest对象
func NewAlibabaPicturesDengtaOrderStatusChangeRequest() *AlibabaPicturesDengtaOrderStatusChangeRequest{
    return &AlibabaPicturesDengtaOrderStatusChangeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaPicturesDengtaOrderStatusChangeRequest) GetApiMethodName() string {
    return "alibaba.pictures.dengta.order.status.change"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaPicturesDengtaOrderStatusChangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Remark Setter
// 拒绝原因
func (r *AlibabaPicturesDengtaOrderStatusChangeRequest) SetRemark(_remark string) error {
    r._remark = _remark
    r.Set("remark", _remark)
    return nil
}

// Remark Getter
func (r AlibabaPicturesDengtaOrderStatusChangeRequest) GetRemark() string {
    return r._remark
}
// Status Setter
// 新状态
func (r *AlibabaPicturesDengtaOrderStatusChangeRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaPicturesDengtaOrderStatusChangeRequest) GetStatus() int64 {
    return r._status
}
// ChangeTime Setter
// 变更时间
func (r *AlibabaPicturesDengtaOrderStatusChangeRequest) SetChangeTime(_changeTime string) error {
    r._changeTime = _changeTime
    r.Set("change_time", _changeTime)
    return nil
}

// ChangeTime Getter
func (r AlibabaPicturesDengtaOrderStatusChangeRequest) GetChangeTime() string {
    return r._changeTime
}
// ImsOrderId Setter
// ims订单编号
func (r *AlibabaPicturesDengtaOrderStatusChangeRequest) SetImsOrderId(_imsOrderId string) error {
    r._imsOrderId = _imsOrderId
    r.Set("ims_order_id", _imsOrderId)
    return nil
}

// ImsOrderId Getter
func (r AlibabaPicturesDengtaOrderStatusChangeRequest) GetImsOrderId() string {
    return r._imsOrderId
}
// AliTaskId Setter
// task 编号
func (r *AlibabaPicturesDengtaOrderStatusChangeRequest) SetAliTaskId(_aliTaskId string) error {
    r._aliTaskId = _aliTaskId
    r.Set("ali_task_id", _aliTaskId)
    return nil
}

// AliTaskId Getter
func (r AlibabaPicturesDengtaOrderStatusChangeRequest) GetAliTaskId() string {
    return r._aliTaskId
}
