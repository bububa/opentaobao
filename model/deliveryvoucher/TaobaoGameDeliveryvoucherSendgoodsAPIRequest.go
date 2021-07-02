package deliveryvoucher

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoGameDeliveryvoucherSendgoodsAPIRequest 提货券发货接口 API请求
// taobao.game.deliveryvoucher.sendgoods
//
// 提货券发券接口：同步券和订单的关联信息
type TaobaoGameDeliveryvoucherSendgoodsAPIRequest struct {
	model.Params
	// 发券参数
	_param0 *SendVoucherRequest
}

// NewTaobaoGameDeliveryvoucherSendgoodsRequest 初始化TaobaoGameDeliveryvoucherSendgoodsAPIRequest对象
func NewTaobaoGameDeliveryvoucherSendgoodsRequest() *TaobaoGameDeliveryvoucherSendgoodsAPIRequest {
	return &TaobaoGameDeliveryvoucherSendgoodsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoGameDeliveryvoucherSendgoodsAPIRequest) GetApiMethodName() string {
	return "taobao.game.deliveryvoucher.sendgoods"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoGameDeliveryvoucherSendgoodsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 发券参数
func (r *TaobaoGameDeliveryvoucherSendgoodsAPIRequest) SetParam0(_param0 *SendVoucherRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r TaobaoGameDeliveryvoucherSendgoodsAPIRequest) GetParam0() *SendVoucherRequest {
	return r._param0
}
