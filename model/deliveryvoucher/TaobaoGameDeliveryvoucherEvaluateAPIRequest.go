package deliveryvoucher

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaogamedeliveryvoucherevaluateAPIRequest 卡券评价回传 API请求
// taobao.game.deliveryvoucher.evaluate
//
// 卡券ISV回传商品评价
type TaobaogamedeliveryvoucherevaluateAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *VoucherEvaluateRequest
}

// NewTaobaogamedeliveryvoucherevaluateRequest 初始化TaobaogamedeliveryvoucherevaluateAPIRequest对象
func NewTaobaogamedeliveryvoucherevaluateRequest() *TaobaogamedeliveryvoucherevaluateAPIRequest {
	return &TaobaogamedeliveryvoucherevaluateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaogamedeliveryvoucherevaluateAPIRequest) GetApiMethodName() string {
	return "taobao.game.deliveryvoucher.evaluate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaogamedeliveryvoucherevaluateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaogamedeliveryvoucherevaluateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 系统自动生成
func (r *TaobaogamedeliveryvoucherevaluateAPIRequest) SetParam0(_param0 *VoucherEvaluateRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaogamedeliveryvoucherevaluateAPIRequest) GetParam0() *VoucherEvaluateRequest {
	return r._param0
}
