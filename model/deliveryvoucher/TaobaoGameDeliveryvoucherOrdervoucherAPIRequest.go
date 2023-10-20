package deliveryvoucher

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaogamedeliveryvoucherordervoucherAPIRequest 预约接口 API请求
// taobao.game.deliveryvoucher.ordervoucher
//
// 提货券发券接口：同步券和订单的关联信息
type TaobaogamedeliveryvoucherordervoucherAPIRequest struct {
	model.Params
	// 发券参数
	_param0 *OrderVoucherRequest
}

// NewTaobaogamedeliveryvoucherordervoucherRequest 初始化TaobaogamedeliveryvoucherordervoucherAPIRequest对象
func NewTaobaogamedeliveryvoucherordervoucherRequest() *TaobaogamedeliveryvoucherordervoucherAPIRequest {
	return &TaobaogamedeliveryvoucherordervoucherAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaogamedeliveryvoucherordervoucherAPIRequest) GetApiMethodName() string {
	return "taobao.game.deliveryvoucher.ordervoucher"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaogamedeliveryvoucherordervoucherAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaogamedeliveryvoucherordervoucherAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 发券参数
func (r *TaobaogamedeliveryvoucherordervoucherAPIRequest) SetParam0(_param0 *OrderVoucherRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaogamedeliveryvoucherordervoucherAPIRequest) GetParam0() *OrderVoucherRequest {
	return r._param0
}
