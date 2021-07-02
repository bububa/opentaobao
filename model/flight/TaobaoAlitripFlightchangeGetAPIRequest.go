package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripFlightchangeGetAPIRequest 获取航变信息 API请求
// taobao.alitrip.flightchange.get
//
// 查询航变是为了两个目的，阿里旅行抓取到未确认航变的航变信息源时可以由代理确认航变，同时对于确认航变的航变信息也共享给代理人做本体业务使用。
type TaobaoAlitripFlightchangeGetAPIRequest struct {
	model.Params
	// 查询信息封装
	_searchOption *FlightChangeDataQueryOption
}

// NewTaobaoAlitripFlightchangeGetRequest 初始化TaobaoAlitripFlightchangeGetAPIRequest对象
func NewTaobaoAlitripFlightchangeGetRequest() *TaobaoAlitripFlightchangeGetAPIRequest {
	return &TaobaoAlitripFlightchangeGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripFlightchangeGetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.flightchange.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripFlightchangeGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SearchOption Setter
// 查询信息封装
func (r *TaobaoAlitripFlightchangeGetAPIRequest) SetSearchOption(_searchOption *FlightChangeDataQueryOption) error {
	r._searchOption = _searchOption
	r.Set("search_option", _searchOption)
	return nil
}

// Get SearchOption Getter
func (r TaobaoAlitripFlightchangeGetAPIRequest) GetSearchOption() *FlightChangeDataQueryOption {
	return r._searchOption
}
