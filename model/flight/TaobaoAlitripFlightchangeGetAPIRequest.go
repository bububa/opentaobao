package flight

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripFlightchangeGetAPIRequest) Reset() {
	r._searchOption = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripFlightchangeGetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.flightchange.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripFlightchangeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripFlightchangeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSearchOption is SearchOption Setter
// 查询信息封装
func (r *TaobaoAlitripFlightchangeGetAPIRequest) SetSearchOption(_searchOption *FlightChangeDataQueryOption) error {
	r._searchOption = _searchOption
	r.Set("search_option", _searchOption)
	return nil
}

// GetSearchOption SearchOption Getter
func (r TaobaoAlitripFlightchangeGetAPIRequest) GetSearchOption() *FlightChangeDataQueryOption {
	return r._searchOption
}

var poolTaobaoAlitripFlightchangeGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripFlightchangeGetRequest()
	},
}

// GetTaobaoAlitripFlightchangeGetRequest 从 sync.Pool 获取 TaobaoAlitripFlightchangeGetAPIRequest
func GetTaobaoAlitripFlightchangeGetAPIRequest() *TaobaoAlitripFlightchangeGetAPIRequest {
	return poolTaobaoAlitripFlightchangeGetAPIRequest.Get().(*TaobaoAlitripFlightchangeGetAPIRequest)
}

// ReleaseTaobaoAlitripFlightchangeGetAPIRequest 将 TaobaoAlitripFlightchangeGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripFlightchangeGetAPIRequest(v *TaobaoAlitripFlightchangeGetAPIRequest) {
	v.Reset()
	poolTaobaoAlitripFlightchangeGetAPIRequest.Put(v)
}
