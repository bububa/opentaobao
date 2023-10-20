package deliveryvoucher

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaogamedeliveryvouchersendvoucherAPIRequest 提货券发券接口 API请求
// taobao.game.deliveryvoucher.sendvoucher
//
// 提货券发券接口：同步券和订单的关联信息
type TaobaogamedeliveryvouchersendvoucherAPIRequest struct {
	model.Params
	// 发券参数
	_param0 *SendVoucherRequest
}

// NewTaobaogamedeliveryvouchersendvoucherRequest 初始化TaobaogamedeliveryvouchersendvoucherAPIRequest对象
func NewTaobaogamedeliveryvouchersendvoucherRequest() *TaobaogamedeliveryvouchersendvoucherAPIRequest {
	return &TaobaogamedeliveryvouchersendvoucherAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaogamedeliveryvouchersendvoucherAPIRequest) GetApiMethodName() string {
	return "taobao.game.deliveryvoucher.sendvoucher"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaogamedeliveryvouchersendvoucherAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaogamedeliveryvouchersendvoucherAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 发券参数
func (r *TaobaogamedeliveryvouchersendvoucherAPIRequest) SetParam0(_param0 *SendVoucherRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaogamedeliveryvouchersendvoucherAPIRequest) GetParam0() *SendVoucherRequest {
	return r._param0
}
