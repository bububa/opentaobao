package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusRefundSetAPIRequest B2B退票申请接口 API请求
// taobao.bus.refund.set
//
// B2B业务支持退票
type TaobaoBusRefundSetAPIRequest struct {
	model.Params
	// 入参
	_param0 *B2brefundOrderRq
}

// NewTaobaoBusRefundSetRequest 初始化TaobaoBusRefundSetAPIRequest对象
func NewTaobaoBusRefundSetRequest() *TaobaoBusRefundSetAPIRequest {
	return &TaobaoBusRefundSetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusRefundSetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.refund.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusRefundSetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusRefundSetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参
func (r *TaobaoBusRefundSetAPIRequest) SetParam0(_param0 *B2brefundOrderRq) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoBusRefundSetAPIRequest) GetParam0() *B2brefundOrderRq {
	return r._param0
}
