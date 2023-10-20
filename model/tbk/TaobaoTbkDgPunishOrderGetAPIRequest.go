package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkdgpunishordergetAPIRequest 淘宝客-推广者-处罚订单查询 API请求
// taobao.tbk.dg.punish.order.get
//
// 新增处罚订单查询API，提供媒体调用查询能力。这个是给媒体自己用的
type TaobaotbkdgpunishordergetAPIRequest struct {
	model.Params
	// 入参的对象
	_afOrderOption *TopApiAfOrderOption
}

// NewTaobaotbkdgpunishordergetRequest 初始化TaobaotbkdgpunishordergetAPIRequest对象
func NewTaobaotbkdgpunishordergetRequest() *TaobaotbkdgpunishordergetAPIRequest {
	return &TaobaotbkdgpunishordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkdgpunishordergetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.punish.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkdgpunishordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkdgpunishordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAfOrderOption is AfOrderOption Setter
// 入参的对象
func (r *TaobaotbkdgpunishordergetAPIRequest) SetAfOrderOption(_afOrderOption *TopApiAfOrderOption) error {
	r._afOrderOption = _afOrderOption
	r.Set("af_order_option", _afOrderOption)
	return nil
}

// GetAfOrderOption AfOrderOption Getter
func (r TaobaotbkdgpunishordergetAPIRequest) GetAfOrderOption() *TopApiAfOrderOption {
	return r._afOrderOption
}
