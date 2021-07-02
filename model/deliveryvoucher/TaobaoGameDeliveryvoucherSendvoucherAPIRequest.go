package deliveryvoucher

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoGameDeliveryvoucherSendvoucherAPIRequest 提货券发券接口 API请求
// taobao.game.deliveryvoucher.sendvoucher
//
// 提货券发券接口：同步券和订单的关联信息
type TaobaoGameDeliveryvoucherSendvoucherAPIRequest struct {
	model.Params
	// 发券参数
	_param0 *SendVoucherRequest
}

// NewTaobaoGameDeliveryvoucherSendvoucherRequest 初始化TaobaoGameDeliveryvoucherSendvoucherAPIRequest对象
func NewTaobaoGameDeliveryvoucherSendvoucherRequest() *TaobaoGameDeliveryvoucherSendvoucherAPIRequest {
	return &TaobaoGameDeliveryvoucherSendvoucherAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoGameDeliveryvoucherSendvoucherAPIRequest) GetApiMethodName() string {
	return "taobao.game.deliveryvoucher.sendvoucher"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoGameDeliveryvoucherSendvoucherAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam0 is Param0 Setter
// 发券参数
func (r *TaobaoGameDeliveryvoucherSendvoucherAPIRequest) SetParam0(_param0 *SendVoucherRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoGameDeliveryvoucherSendvoucherAPIRequest) GetParam0() *SendVoucherRequest {
	return r._param0
}
