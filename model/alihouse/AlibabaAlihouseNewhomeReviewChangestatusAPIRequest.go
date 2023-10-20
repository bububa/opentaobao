package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomereviewchangestatusAPIRequest 楼盘测评草稿状态同步 API请求
// alibaba.alihouse.newhome.review.changestatus
//
// 楼盘测评草稿状态更新
type AlibabaalihousenewhomereviewchangestatusAPIRequest struct {
	model.Params
	// 外部测评id
	_outerId string
	// 0 失效 1 有效
	_status int64
}

// NewAlibabaalihousenewhomereviewchangestatusRequest 初始化AlibabaalihousenewhomereviewchangestatusAPIRequest对象
func NewAlibabaalihousenewhomereviewchangestatusRequest() *AlibabaalihousenewhomereviewchangestatusAPIRequest {
	return &AlibabaalihousenewhomereviewchangestatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomereviewchangestatusAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.review.changestatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomereviewchangestatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomereviewchangestatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 外部测评id
func (r *AlibabaalihousenewhomereviewchangestatusAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaalihousenewhomereviewchangestatusAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetStatus is Status Setter
// 0 失效 1 有效
func (r *AlibabaalihousenewhomereviewchangestatusAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaalihousenewhomereviewchangestatusAPIRequest) GetStatus() int64 {
	return r._status
}
