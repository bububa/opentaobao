package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeWtverticalGetAPIRequest 网厅垂直信息查询接口 API请求
// taobao.trade.wtvertical.get
//
// 网厅订单垂直信息的查询
type TaobaoTradeWtverticalGetAPIRequest struct {
	model.Params
	// 主订单列表,用“，”分隔tid的字符串,最大列表长度为15
	_tids []string
}

// NewTaobaoTradeWtverticalGetRequest 初始化TaobaoTradeWtverticalGetAPIRequest对象
func NewTaobaoTradeWtverticalGetRequest() *TaobaoTradeWtverticalGetAPIRequest {
	return &TaobaoTradeWtverticalGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTradeWtverticalGetAPIRequest) GetApiMethodName() string {
	return "taobao.trade.wtvertical.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTradeWtverticalGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTradeWtverticalGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTids is Tids Setter
// 主订单列表,用“，”分隔tid的字符串,最大列表长度为15
func (r *TaobaoTradeWtverticalGetAPIRequest) SetTids(_tids []string) error {
	r._tids = _tids
	r.Set("tids", _tids)
	return nil
}

// GetTids Tids Getter
func (r TaobaoTradeWtverticalGetAPIRequest) GetTids() []string {
	return r._tids
}
