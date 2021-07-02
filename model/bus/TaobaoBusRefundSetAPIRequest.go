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
	_param0 *B2BRefundOrderRq
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
func (r TaobaoBusRefundSetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 入参
func (r *TaobaoBusRefundSetAPIRequest) SetParam0(_param0 *B2BRefundOrderRq) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r TaobaoBusRefundSetAPIRequest) GetParam0() *B2BRefundOrderRq {
	return r._param0
}
