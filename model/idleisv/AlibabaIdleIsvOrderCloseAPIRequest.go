package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleisvordercloseAPIRequest 服务商闲鱼卖家主动关闭订单 API请求
// alibaba.idle.isv.order.close
//
// 供外部服务商 isv 提供卖家主动关闭交易订单的功能
type AlibabaidleisvordercloseAPIRequest struct {
	model.Params
	// 输入参数
	_isvAppraiseIsvOrderCloseDto *AppraiseIsvOrderCloseDto
}

// NewAlibabaidleisvordercloseRequest 初始化AlibabaidleisvordercloseAPIRequest对象
func NewAlibabaidleisvordercloseRequest() *AlibabaidleisvordercloseAPIRequest {
	return &AlibabaidleisvordercloseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleisvordercloseAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.order.close"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleisvordercloseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleisvordercloseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIsvAppraiseIsvOrderCloseDto is IsvAppraiseIsvOrderCloseDto Setter
// 输入参数
func (r *AlibabaidleisvordercloseAPIRequest) SetIsvAppraiseIsvOrderCloseDto(_isvAppraiseIsvOrderCloseDto *AppraiseIsvOrderCloseDto) error {
	r._isvAppraiseIsvOrderCloseDto = _isvAppraiseIsvOrderCloseDto
	r.Set("isv_appraise_isv_order_close_dto", _isvAppraiseIsvOrderCloseDto)
	return nil
}

// GetIsvAppraiseIsvOrderCloseDto IsvAppraiseIsvOrderCloseDto Getter
func (r AlibabaidleisvordercloseAPIRequest) GetIsvAppraiseIsvOrderCloseDto() *AppraiseIsvOrderCloseDto {
	return r._isvAppraiseIsvOrderCloseDto
}
