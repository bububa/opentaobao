package deliveryvoucher

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoGameDeliveryvoucherOrdervoucherAPIRequest 预约接口 API请求
// taobao.game.deliveryvoucher.ordervoucher
//
// 提货券发券接口：同步券和订单的关联信息
type TaobaoGameDeliveryvoucherOrdervoucherAPIRequest struct {
	model.Params
	// 发券参数
	_param0 *OrderVoucherRequest
}

// NewTaobaoGameDeliveryvoucherOrdervoucherRequest 初始化TaobaoGameDeliveryvoucherOrdervoucherAPIRequest对象
func NewTaobaoGameDeliveryvoucherOrdervoucherRequest() *TaobaoGameDeliveryvoucherOrdervoucherAPIRequest {
	return &TaobaoGameDeliveryvoucherOrdervoucherAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoGameDeliveryvoucherOrdervoucherAPIRequest) GetApiMethodName() string {
	return "taobao.game.deliveryvoucher.ordervoucher"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoGameDeliveryvoucherOrdervoucherAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoGameDeliveryvoucherOrdervoucherAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 发券参数
func (r *TaobaoGameDeliveryvoucherOrdervoucherAPIRequest) SetParam0(_param0 *OrderVoucherRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoGameDeliveryvoucherOrdervoucherAPIRequest) GetParam0() *OrderVoucherRequest {
	return r._param0
}
