package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamjocpayAPIRequest POS收银成功后订单同步 API请求
// alibaba.mj.oc.pay
//
// 此API用于在银泰商场中，消费者在收银台收银/退款时， POS系统在收银或退款成功后，调用此接口进行订单同步
type AlibabamjocpayAPIRequest struct {
	model.Params
	// 订单数据
	_posOrder *PosOrderDto
}

// NewAlibabamjocpayRequest 初始化AlibabamjocpayAPIRequest对象
func NewAlibabamjocpayRequest() *AlibabamjocpayAPIRequest {
	return &AlibabamjocpayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamjocpayAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.oc.pay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamjocpayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamjocpayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPosOrder is PosOrder Setter
// 订单数据
func (r *AlibabamjocpayAPIRequest) SetPosOrder(_posOrder *PosOrderDto) error {
	r._posOrder = _posOrder
	r.Set("pos_order", _posOrder)
	return nil
}

// GetPosOrder PosOrder Getter
func (r AlibabamjocpayAPIRequest) GetPosOrder() *PosOrderDto {
	return r._posOrder
}
