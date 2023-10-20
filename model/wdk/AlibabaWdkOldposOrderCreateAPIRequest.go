package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkoldposordercreateAPIRequest 淘鲜达外部商户老pos机产生的订单同步进淘鲜达 API请求
// alibaba.wdk.oldpos.order.create
//
// 淘鲜达外部商户老pos机产生的订单同步进淘鲜达
type AlibabawdkoldposordercreateAPIRequest struct {
	model.Params
	// 入参
	_posOrderCreateRequest *PosOrderCreateRequest
}

// NewAlibabawdkoldposordercreateRequest 初始化AlibabawdkoldposordercreateAPIRequest对象
func NewAlibabawdkoldposordercreateRequest() *AlibabawdkoldposordercreateAPIRequest {
	return &AlibabawdkoldposordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkoldposordercreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.oldpos.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkoldposordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkoldposordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPosOrderCreateRequest is PosOrderCreateRequest Setter
// 入参
func (r *AlibabawdkoldposordercreateAPIRequest) SetPosOrderCreateRequest(_posOrderCreateRequest *PosOrderCreateRequest) error {
	r._posOrderCreateRequest = _posOrderCreateRequest
	r.Set("pos_order_create_request", _posOrderCreateRequest)
	return nil
}

// GetPosOrderCreateRequest PosOrderCreateRequest Getter
func (r AlibabawdkoldposordercreateAPIRequest) GetPosOrderCreateRequest() *PosOrderCreateRequest {
	return r._posOrderCreateRequest
}
