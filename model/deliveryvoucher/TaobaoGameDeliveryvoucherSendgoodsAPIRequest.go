package deliveryvoucher

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaogamedeliveryvouchersendgoodsAPIRequest 提货券发货接口 API请求
// taobao.game.deliveryvoucher.sendgoods
//
// 提货券发券接口：同步券和订单的关联信息
type TaobaogamedeliveryvouchersendgoodsAPIRequest struct {
	model.Params
	// 发券参数
	_param0 *SendVoucherRequest
}

// NewTaobaogamedeliveryvouchersendgoodsRequest 初始化TaobaogamedeliveryvouchersendgoodsAPIRequest对象
func NewTaobaogamedeliveryvouchersendgoodsRequest() *TaobaogamedeliveryvouchersendgoodsAPIRequest {
	return &TaobaogamedeliveryvouchersendgoodsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaogamedeliveryvouchersendgoodsAPIRequest) GetApiMethodName() string {
	return "taobao.game.deliveryvoucher.sendgoods"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaogamedeliveryvouchersendgoodsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaogamedeliveryvouchersendgoodsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 发券参数
func (r *TaobaogamedeliveryvouchersendgoodsAPIRequest) SetParam0(_param0 *SendVoucherRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaogamedeliveryvouchersendgoodsAPIRequest) GetParam0() *SendVoucherRequest {
	return r._param0
}
