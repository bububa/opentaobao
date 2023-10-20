package dengta

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabapicturesdengtaorderstatuschangeAPIRequest 天下秀订单状态变更通知 API请求
// alibaba.pictures.dengta.order.status.change
//
// 天下秀订单状态变更通知
type AlibabapicturesdengtaorderstatuschangeAPIRequest struct {
	model.Params
	// 拒绝原因
	_remark string
	// 变更时间
	_changeTime string
	// ims订单编号
	_imsOrderId string
	// task 编号
	_aliTaskId string
	// 新状态
	_status int64
}

// NewAlibabapicturesdengtaorderstatuschangeRequest 初始化AlibabapicturesdengtaorderstatuschangeAPIRequest对象
func NewAlibabapicturesdengtaorderstatuschangeRequest() *AlibabapicturesdengtaorderstatuschangeAPIRequest {
	return &AlibabapicturesdengtaorderstatuschangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabapicturesdengtaorderstatuschangeAPIRequest) GetApiMethodName() string {
	return "alibaba.pictures.dengta.order.status.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabapicturesdengtaorderstatuschangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabapicturesdengtaorderstatuschangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRemark is Remark Setter
// 拒绝原因
func (r *AlibabapicturesdengtaorderstatuschangeAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r AlibabapicturesdengtaorderstatuschangeAPIRequest) GetRemark() string {
	return r._remark
}

// SetChangeTime is ChangeTime Setter
// 变更时间
func (r *AlibabapicturesdengtaorderstatuschangeAPIRequest) SetChangeTime(_changeTime string) error {
	r._changeTime = _changeTime
	r.Set("change_time", _changeTime)
	return nil
}

// GetChangeTime ChangeTime Getter
func (r AlibabapicturesdengtaorderstatuschangeAPIRequest) GetChangeTime() string {
	return r._changeTime
}

// SetImsOrderId is ImsOrderId Setter
// ims订单编号
func (r *AlibabapicturesdengtaorderstatuschangeAPIRequest) SetImsOrderId(_imsOrderId string) error {
	r._imsOrderId = _imsOrderId
	r.Set("ims_order_id", _imsOrderId)
	return nil
}

// GetImsOrderId ImsOrderId Getter
func (r AlibabapicturesdengtaorderstatuschangeAPIRequest) GetImsOrderId() string {
	return r._imsOrderId
}

// SetAliTaskId is AliTaskId Setter
// task 编号
func (r *AlibabapicturesdengtaorderstatuschangeAPIRequest) SetAliTaskId(_aliTaskId string) error {
	r._aliTaskId = _aliTaskId
	r.Set("ali_task_id", _aliTaskId)
	return nil
}

// GetAliTaskId AliTaskId Getter
func (r AlibabapicturesdengtaorderstatuschangeAPIRequest) GetAliTaskId() string {
	return r._aliTaskId
}

// SetStatus is Status Setter
// 新状态
func (r *AlibabapicturesdengtaorderstatuschangeAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabapicturesdengtaorderstatuschangeAPIRequest) GetStatus() int64 {
	return r._status
}
