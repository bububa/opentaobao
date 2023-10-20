package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaorefundrefundactionsdisplayAPIRequest 退货退款操作的展示信息(展现给买家) API请求
// cainiao.refund.refundactions.display
//
// 退货退款操作的展示信息(展现给买家)
type CainiaorefundrefundactionsdisplayAPIRequest struct {
	model.Params
	// 请求入参
	_param0 *OrderRefundOperationReq
}

// NewCainiaorefundrefundactionsdisplayRequest 初始化CainiaorefundrefundactionsdisplayAPIRequest对象
func NewCainiaorefundrefundactionsdisplayRequest() *CainiaorefundrefundactionsdisplayAPIRequest {
	return &CainiaorefundrefundactionsdisplayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaorefundrefundactionsdisplayAPIRequest) GetApiMethodName() string {
	return "cainiao.refund.refundactions.display"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaorefundrefundactionsdisplayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaorefundrefundactionsdisplayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 请求入参
func (r *CainiaorefundrefundactionsdisplayAPIRequest) SetParam0(_param0 *OrderRefundOperationReq) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r CainiaorefundrefundactionsdisplayAPIRequest) GetParam0() *OrderRefundOperationReq {
	return r._param0
}
