package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaorefundrefundactionsjudgementAPIRequest 判断当前用户是否能对订单执行一些逆向操作，比如退货操作 API请求
// cainiao.refund.refundactions.judgement
//
// 判断当前用户是否能对订单执行一些逆向操作，比如退货操作
type CainiaorefundrefundactionsjudgementAPIRequest struct {
	model.Params
	// 操作请求
	_param0 *OrderRefundOperationJudgementReq
}

// NewCainiaorefundrefundactionsjudgementRequest 初始化CainiaorefundrefundactionsjudgementAPIRequest对象
func NewCainiaorefundrefundactionsjudgementRequest() *CainiaorefundrefundactionsjudgementAPIRequest {
	return &CainiaorefundrefundactionsjudgementAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaorefundrefundactionsjudgementAPIRequest) GetApiMethodName() string {
	return "cainiao.refund.refundactions.judgement"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaorefundrefundactionsjudgementAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaorefundrefundactionsjudgementAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 操作请求
func (r *CainiaorefundrefundactionsjudgementAPIRequest) SetParam0(_param0 *OrderRefundOperationJudgementReq) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r CainiaorefundrefundactionsjudgementAPIRequest) GetParam0() *OrderRefundOperationJudgementReq {
	return r._param0
}
