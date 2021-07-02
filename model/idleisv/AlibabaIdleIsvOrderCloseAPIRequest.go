package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvOrderCloseAPIRequest 服务商闲鱼卖家主动关闭订单 API请求
// alibaba.idle.isv.order.close
//
// 供外部服务商 isv 提供卖家主动关闭交易订单的功能
type AlibabaIdleIsvOrderCloseAPIRequest struct {
	model.Params
	// 输入参数
	_isvAppraiseIsvOrderCloseDto *AppraiseIsvOrderCloseDto
}

// NewAlibabaIdleIsvOrderCloseRequest 初始化AlibabaIdleIsvOrderCloseAPIRequest对象
func NewAlibabaIdleIsvOrderCloseRequest() *AlibabaIdleIsvOrderCloseAPIRequest {
	return &AlibabaIdleIsvOrderCloseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvOrderCloseAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.order.close"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvOrderCloseAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is IsvAppraiseIsvOrderCloseDto Setter
// 输入参数
func (r *AlibabaIdleIsvOrderCloseAPIRequest) SetIsvAppraiseIsvOrderCloseDto(_isvAppraiseIsvOrderCloseDto *AppraiseIsvOrderCloseDto) error {
	r._isvAppraiseIsvOrderCloseDto = _isvAppraiseIsvOrderCloseDto
	r.Set("isv_appraise_isv_order_close_dto", _isvAppraiseIsvOrderCloseDto)
	return nil
}

// Get IsvAppraiseIsvOrderCloseDto Getter
func (r AlibabaIdleIsvOrderCloseAPIRequest) GetIsvAppraiseIsvOrderCloseDto() *AppraiseIsvOrderCloseDto {
	return r._isvAppraiseIsvOrderCloseDto
}
