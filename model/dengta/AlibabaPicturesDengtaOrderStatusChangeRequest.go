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
    remark   string
    // 新状态
    status   int64
    // 变更时间
    changeTime   string
    // ims订单编号
    imsOrderId   string
    // task 编号
    aliTaskId   string
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
func (r *AlibabaPicturesDengtaOrderStatusChangeRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

// Remark Getter
func (r AlibabaPicturesDengtaOrderStatusChangeRequest) GetRemark() string {
    return r.remark
}
// Status Setter
// 新状态
func (r *AlibabaPicturesDengtaOrderStatusChangeRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r AlibabaPicturesDengtaOrderStatusChangeRequest) GetStatus() int64 {
    return r.status
}
// ChangeTime Setter
// 变更时间
func (r *AlibabaPicturesDengtaOrderStatusChangeRequest) SetChangeTime(changeTime string) error {
    r.changeTime = changeTime
    r.Set("change_time", changeTime)
    return nil
}

// ChangeTime Getter
func (r AlibabaPicturesDengtaOrderStatusChangeRequest) GetChangeTime() string {
    return r.changeTime
}
// ImsOrderId Setter
// ims订单编号
func (r *AlibabaPicturesDengtaOrderStatusChangeRequest) SetImsOrderId(imsOrderId string) error {
    r.imsOrderId = imsOrderId
    r.Set("ims_order_id", imsOrderId)
    return nil
}

// ImsOrderId Getter
func (r AlibabaPicturesDengtaOrderStatusChangeRequest) GetImsOrderId() string {
    return r.imsOrderId
}
// AliTaskId Setter
// task 编号
func (r *AlibabaPicturesDengtaOrderStatusChangeRequest) SetAliTaskId(aliTaskId string) error {
    r.aliTaskId = aliTaskId
    r.Set("ali_task_id", aliTaskId)
    return nil
}

// AliTaskId Getter
func (r AlibabaPicturesDengtaOrderStatusChangeRequest) GetAliTaskId() string {
    return r.aliTaskId
}
