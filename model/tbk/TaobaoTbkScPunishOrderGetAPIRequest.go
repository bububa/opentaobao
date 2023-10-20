package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkscpunishordergetAPIRequest 淘宝客-服务商-处罚订单查询 API请求
// taobao.tbk.sc.punish.order.get
//
// 服务商使用。可通过该接口查询推广者账号下在处罚管理后台，可直接下载的处罚订单明细。请注意：若是基于账号(member)、媒体id(site)、推广位(adzone)、渠道方(relationid)维度的完整处罚，对应订单明细请根据处罚后台对应的处罚订单时间说明，在“推广订单明细”中筛选对应订单。
type TaobaotbkscpunishordergetAPIRequest struct {
	model.Params
	// 入参的对象
	_aforderoption *TopApiAfOrderOption
}

// NewTaobaotbkscpunishordergetRequest 初始化TaobaotbkscpunishordergetAPIRequest对象
func NewTaobaotbkscpunishordergetRequest() *TaobaotbkscpunishordergetAPIRequest {
	return &TaobaotbkscpunishordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkscpunishordergetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.punish.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkscpunishordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkscpunishordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAforderoption is Aforderoption Setter
// 入参的对象
func (r *TaobaotbkscpunishordergetAPIRequest) SetAforderoption(_aforderoption *TopApiAfOrderOption) error {
	r._aforderoption = _aforderoption
	r.Set("af_order_option", _aforderoption)
	return nil
}

// GetAforderoption Aforderoption Getter
func (r TaobaotbkscpunishordergetAPIRequest) GetAforderoption() *TopApiAfOrderOption {
	return r._aforderoption
}
