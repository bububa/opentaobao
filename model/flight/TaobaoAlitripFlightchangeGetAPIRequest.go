package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripflightchangegetAPIRequest 获取航变信息 API请求
// taobao.alitrip.flightchange.get
//
// 查询航变是为了两个目的，阿里旅行抓取到未确认航变的航变信息源时可以由代理确认航变，同时对于确认航变的航变信息也共享给代理人做本体业务使用。
type TaobaoalitripflightchangegetAPIRequest struct {
	model.Params
	// 查询信息封装
	_searchOption *FlightChangeDataQueryOption
}

// NewTaobaoalitripflightchangegetRequest 初始化TaobaoalitripflightchangegetAPIRequest对象
func NewTaobaoalitripflightchangegetRequest() *TaobaoalitripflightchangegetAPIRequest {
	return &TaobaoalitripflightchangegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripflightchangegetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.flightchange.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripflightchangegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripflightchangegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSearchOption is SearchOption Setter
// 查询信息封装
func (r *TaobaoalitripflightchangegetAPIRequest) SetSearchOption(_searchOption *FlightChangeDataQueryOption) error {
	r._searchOption = _searchOption
	r.Set("search_option", _searchOption)
	return nil
}

// GetSearchOption SearchOption Getter
func (r TaobaoalitripflightchangegetAPIRequest) GetSearchOption() *FlightChangeDataQueryOption {
	return r._searchOption
}
