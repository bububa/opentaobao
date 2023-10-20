package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomevideochangestatusAPIRequest 视频草稿状态更新 API请求
// alibaba.alihouse.newhome.video.changestatus
//
// 视频草稿状态更新
type AlibabaalihousenewhomevideochangestatusAPIRequest struct {
	model.Params
	// 外部视频id
	_outerId string
	// 外部楼盘id（不可多值）
	_outerProjectId string
	// 外部门店id
	_outerStoreId string
	// 0 失效 1 有效
	_status int64
}

// NewAlibabaalihousenewhomevideochangestatusRequest 初始化AlibabaalihousenewhomevideochangestatusAPIRequest对象
func NewAlibabaalihousenewhomevideochangestatusRequest() *AlibabaalihousenewhomevideochangestatusAPIRequest {
	return &AlibabaalihousenewhomevideochangestatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomevideochangestatusAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.video.changestatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomevideochangestatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomevideochangestatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 外部视频id
func (r *AlibabaalihousenewhomevideochangestatusAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaalihousenewhomevideochangestatusAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetOuterProjectId is OuterProjectId Setter
// 外部楼盘id（不可多值）
func (r *AlibabaalihousenewhomevideochangestatusAPIRequest) SetOuterProjectId(_outerProjectId string) error {
	r._outerProjectId = _outerProjectId
	r.Set("outer_project_id", _outerProjectId)
	return nil
}

// GetOuterProjectId OuterProjectId Getter
func (r AlibabaalihousenewhomevideochangestatusAPIRequest) GetOuterProjectId() string {
	return r._outerProjectId
}

// SetOuterStoreId is OuterStoreId Setter
// 外部门店id
func (r *AlibabaalihousenewhomevideochangestatusAPIRequest) SetOuterStoreId(_outerStoreId string) error {
	r._outerStoreId = _outerStoreId
	r.Set("outer_store_id", _outerStoreId)
	return nil
}

// GetOuterStoreId OuterStoreId Getter
func (r AlibabaalihousenewhomevideochangestatusAPIRequest) GetOuterStoreId() string {
	return r._outerStoreId
}

// SetStatus is Status Setter
// 0 失效 1 有效
func (r *AlibabaalihousenewhomevideochangestatusAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaalihousenewhomevideochangestatusAPIRequest) GetStatus() int64 {
	return r._status
}
