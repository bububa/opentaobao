package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgPunishOrderGetAPIRequest 淘宝客-推广者-处罚订单查询 API请求
// taobao.tbk.dg.punish.order.get
//
// 新增处罚订单查询API，提供媒体调用查询能力。这个是给媒体自己用的
type TaobaoTbkDgPunishOrderGetAPIRequest struct {
	model.Params
	// 入参的对象
	_afOrderOption *TopApiAfOrderOption
}

// NewTaobaoTbkDgPunishOrderGetRequest 初始化TaobaoTbkDgPunishOrderGetAPIRequest对象
func NewTaobaoTbkDgPunishOrderGetRequest() *TaobaoTbkDgPunishOrderGetAPIRequest {
	return &TaobaoTbkDgPunishOrderGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkDgPunishOrderGetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.punish.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkDgPunishOrderGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAfOrderOption is AfOrderOption Setter
// 入参的对象
func (r *TaobaoTbkDgPunishOrderGetAPIRequest) SetAfOrderOption(_afOrderOption *TopApiAfOrderOption) error {
	r._afOrderOption = _afOrderOption
	r.Set("af_order_option", _afOrderOption)
	return nil
}

// GetAfOrderOption AfOrderOption Getter
func (r TaobaoTbkDgPunishOrderGetAPIRequest) GetAfOrderOption() *TopApiAfOrderOption {
	return r._afOrderOption
}
