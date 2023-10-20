package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomercchangestatusAPIRequest 图文草稿状态更新 API请求
// alibaba.alihouse.newhome.rc.changestatus
//
// 图文草稿状态更新
type AlibabaalihousenewhomercchangestatusAPIRequest struct {
	model.Params
	// 外部图文id
	_outerId string
	// 外部楼盘id
	_outerProjectId string
	// 外部门店id
	_outerStoreId string
	// 0 失效 1 有效
	_status int64
}

// NewAlibabaalihousenewhomercchangestatusRequest 初始化AlibabaalihousenewhomercchangestatusAPIRequest对象
func NewAlibabaalihousenewhomercchangestatusRequest() *AlibabaalihousenewhomercchangestatusAPIRequest {
	return &AlibabaalihousenewhomercchangestatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomercchangestatusAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.rc.changestatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomercchangestatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomercchangestatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 外部图文id
func (r *AlibabaalihousenewhomercchangestatusAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaalihousenewhomercchangestatusAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetOuterProjectId is OuterProjectId Setter
// 外部楼盘id
func (r *AlibabaalihousenewhomercchangestatusAPIRequest) SetOuterProjectId(_outerProjectId string) error {
	r._outerProjectId = _outerProjectId
	r.Set("outer_project_id", _outerProjectId)
	return nil
}

// GetOuterProjectId OuterProjectId Getter
func (r AlibabaalihousenewhomercchangestatusAPIRequest) GetOuterProjectId() string {
	return r._outerProjectId
}

// SetOuterStoreId is OuterStoreId Setter
// 外部门店id
func (r *AlibabaalihousenewhomercchangestatusAPIRequest) SetOuterStoreId(_outerStoreId string) error {
	r._outerStoreId = _outerStoreId
	r.Set("outer_store_id", _outerStoreId)
	return nil
}

// GetOuterStoreId OuterStoreId Getter
func (r AlibabaalihousenewhomercchangestatusAPIRequest) GetOuterStoreId() string {
	return r._outerStoreId
}

// SetStatus is Status Setter
// 0 失效 1 有效
func (r *AlibabaalihousenewhomercchangestatusAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaalihousenewhomercchangestatusAPIRequest) GetStatus() int64 {
	return r._status
}
