package dengta

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPicturesDengtaOrderStatusChangeAPIRequest 天下秀订单状态变更通知 API请求
// alibaba.pictures.dengta.order.status.change
//
// 天下秀订单状态变更通知
type AlibabaPicturesDengtaOrderStatusChangeAPIRequest struct {
	model.Params
	// 拒绝原因
	_remark string
	// 新状态
	_status int64
	// 变更时间
	_changeTime string
	// ims订单编号
	_imsOrderId string
	// task 编号
	_aliTaskId string
}

// NewAlibabaPicturesDengtaOrderStatusChangeRequest 初始化AlibabaPicturesDengtaOrderStatusChangeAPIRequest对象
func NewAlibabaPicturesDengtaOrderStatusChangeRequest() *AlibabaPicturesDengtaOrderStatusChangeAPIRequest {
	return &AlibabaPicturesDengtaOrderStatusChangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPicturesDengtaOrderStatusChangeAPIRequest) GetApiMethodName() string {
	return "alibaba.pictures.dengta.order.status.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPicturesDengtaOrderStatusChangeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRemark is Remark Setter
// 拒绝原因
func (r *AlibabaPicturesDengtaOrderStatusChangeAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r AlibabaPicturesDengtaOrderStatusChangeAPIRequest) GetRemark() string {
	return r._remark
}

// SetStatus is Status Setter
// 新状态
func (r *AlibabaPicturesDengtaOrderStatusChangeAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaPicturesDengtaOrderStatusChangeAPIRequest) GetStatus() int64 {
	return r._status
}

// SetChangeTime is ChangeTime Setter
// 变更时间
func (r *AlibabaPicturesDengtaOrderStatusChangeAPIRequest) SetChangeTime(_changeTime string) error {
	r._changeTime = _changeTime
	r.Set("change_time", _changeTime)
	return nil
}

// GetChangeTime ChangeTime Getter
func (r AlibabaPicturesDengtaOrderStatusChangeAPIRequest) GetChangeTime() string {
	return r._changeTime
}

// SetImsOrderId is ImsOrderId Setter
// ims订单编号
func (r *AlibabaPicturesDengtaOrderStatusChangeAPIRequest) SetImsOrderId(_imsOrderId string) error {
	r._imsOrderId = _imsOrderId
	r.Set("ims_order_id", _imsOrderId)
	return nil
}

// GetImsOrderId ImsOrderId Getter
func (r AlibabaPicturesDengtaOrderStatusChangeAPIRequest) GetImsOrderId() string {
	return r._imsOrderId
}

// SetAliTaskId is AliTaskId Setter
// task 编号
func (r *AlibabaPicturesDengtaOrderStatusChangeAPIRequest) SetAliTaskId(_aliTaskId string) error {
	r._aliTaskId = _aliTaskId
	r.Set("ali_task_id", _aliTaskId)
	return nil
}

// GetAliTaskId AliTaskId Getter
func (r AlibabaPicturesDengtaOrderStatusChangeAPIRequest) GetAliTaskId() string {
	return r._aliTaskId
}
