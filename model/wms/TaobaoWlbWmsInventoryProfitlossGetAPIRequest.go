package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbwmsinventoryprofitlossgetAPIRequest 通过订单列表批量获取库存损益单信息 API请求
// taobao.wlb.wms.inventory.profitloss.get
//
// 通过订单列表批量获取库存损益单信息
type TaobaowlbwmsinventoryprofitlossgetAPIRequest struct {
	model.Params
	// 菜鸟订单编码
	_cnOrderCode string
}

// NewTaobaowlbwmsinventoryprofitlossgetRequest 初始化TaobaowlbwmsinventoryprofitlossgetAPIRequest对象
func NewTaobaowlbwmsinventoryprofitlossgetRequest() *TaobaowlbwmsinventoryprofitlossgetAPIRequest {
	return &TaobaowlbwmsinventoryprofitlossgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbwmsinventoryprofitlossgetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.wms.inventory.profitloss.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbwmsinventoryprofitlossgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbwmsinventoryprofitlossgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCnOrderCode is CnOrderCode Setter
// 菜鸟订单编码
func (r *TaobaowlbwmsinventoryprofitlossgetAPIRequest) SetCnOrderCode(_cnOrderCode string) error {
	r._cnOrderCode = _cnOrderCode
	r.Set("cn_order_code", _cnOrderCode)
	return nil
}

// GetCnOrderCode CnOrderCode Getter
func (r TaobaowlbwmsinventoryprofitlossgetAPIRequest) GetCnOrderCode() string {
	return r._cnOrderCode
}
