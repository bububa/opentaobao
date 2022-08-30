package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScPunishOrderGetAPIRequest 淘宝客-服务商-处罚订单查询 API请求
// taobao.tbk.sc.punish.order.get
//
// 服务商使用。可通过该接口查询推广者账号下在处罚管理后台，可直接下载的处罚订单明细。请注意：若是基于账号(member)、媒体id(site)、推广位(adzone)、渠道方(relationid)维度的完整处罚，对应订单明细请根据处罚后台对应的处罚订单时间说明，在“推广订单明细”中筛选对应订单。
type TaobaoTbkScPunishOrderGetAPIRequest struct {
	model.Params
	// 入参的对象
	_afOrderOption *TopApiAfOrderOption
}

// NewTaobaoTbkScPunishOrderGetRequest 初始化TaobaoTbkScPunishOrderGetAPIRequest对象
func NewTaobaoTbkScPunishOrderGetRequest() *TaobaoTbkScPunishOrderGetAPIRequest {
	return &TaobaoTbkScPunishOrderGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkScPunishOrderGetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.punish.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkScPunishOrderGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAfOrderOption is AfOrderOption Setter
// 入参的对象
func (r *TaobaoTbkScPunishOrderGetAPIRequest) SetAfOrderOption(_afOrderOption *TopApiAfOrderOption) error {
	r._afOrderOption = _afOrderOption
	r.Set("af_order_option", _afOrderOption)
	return nil
}

// GetAfOrderOption AfOrderOption Getter
func (r TaobaoTbkScPunishOrderGetAPIRequest) GetAfOrderOption() *TopApiAfOrderOption {
	return r._afOrderOption
}
