package deliveryvoucher

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoGameDeliveryvoucherEvaluateAPIRequest 卡券评价回传 API请求
// taobao.game.deliveryvoucher.evaluate
//
// 卡券ISV回传商品评价
type TaobaoGameDeliveryvoucherEvaluateAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *VoucherEvaluateRequest
}

// NewTaobaoGameDeliveryvoucherEvaluateRequest 初始化TaobaoGameDeliveryvoucherEvaluateAPIRequest对象
func NewTaobaoGameDeliveryvoucherEvaluateRequest() *TaobaoGameDeliveryvoucherEvaluateAPIRequest {
	return &TaobaoGameDeliveryvoucherEvaluateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoGameDeliveryvoucherEvaluateAPIRequest) GetApiMethodName() string {
	return "taobao.game.deliveryvoucher.evaluate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoGameDeliveryvoucherEvaluateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam0 is Param0 Setter
// 系统自动生成
func (r *TaobaoGameDeliveryvoucherEvaluateAPIRequest) SetParam0(_param0 *VoucherEvaluateRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoGameDeliveryvoucherEvaluateAPIRequest) GetParam0() *VoucherEvaluateRequest {
	return r._param0
}
