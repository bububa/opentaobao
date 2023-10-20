package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkoldposrefundcreateAPIRequest 五道口外部商户老pos机产生的退款单同步进盒马 API请求
// alibaba.wdk.oldpos.refund.create
//
// 淘鲜达外部商户老pos机产生的退款单同步进淘鲜达
type AlibabawdkoldposrefundcreateAPIRequest struct {
	model.Params
	// 入参
	_posRefundCreateRequest *PosRefundCreateRequest
}

// NewAlibabawdkoldposrefundcreateRequest 初始化AlibabawdkoldposrefundcreateAPIRequest对象
func NewAlibabawdkoldposrefundcreateRequest() *AlibabawdkoldposrefundcreateAPIRequest {
	return &AlibabawdkoldposrefundcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkoldposrefundcreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.oldpos.refund.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkoldposrefundcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkoldposrefundcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPosRefundCreateRequest is PosRefundCreateRequest Setter
// 入参
func (r *AlibabawdkoldposrefundcreateAPIRequest) SetPosRefundCreateRequest(_posRefundCreateRequest *PosRefundCreateRequest) error {
	r._posRefundCreateRequest = _posRefundCreateRequest
	r.Set("pos_refund_create_request", _posRefundCreateRequest)
	return nil
}

// GetPosRefundCreateRequest PosRefundCreateRequest Getter
func (r AlibabawdkoldposrefundcreateAPIRequest) GetPosRefundCreateRequest() *PosRefundCreateRequest {
	return r._posRefundCreateRequest
}
