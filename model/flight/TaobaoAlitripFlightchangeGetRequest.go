package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取航变信息 API请求
taobao.alitrip.flightchange.get

查询航变是为了两个目的，阿里旅行抓取到未确认航变的航变信息源时可以由代理确认航变，同时对于确认航变的航变信息也共享给代理人做本体业务使用。
*/
type TaobaoAlitripFlightchangeGetRequest struct {
    model.Params
    // 查询信息封装
    searchOption   *FlightChangeDataQueryOption
}

// 初始化TaobaoAlitripFlightchangeGetRequest对象
func NewTaobaoAlitripFlightchangeGetRequest() *TaobaoAlitripFlightchangeGetRequest{
    return &TaobaoAlitripFlightchangeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripFlightchangeGetRequest) GetApiMethodName() string {
    return "taobao.alitrip.flightchange.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripFlightchangeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SearchOption Setter
// 查询信息封装
func (r *TaobaoAlitripFlightchangeGetRequest) SetSearchOption(searchOption *FlightChangeDataQueryOption) error {
    r.searchOption = searchOption
    r.Set("search_option", searchOption)
    return nil
}

// SearchOption Getter
func (r TaobaoAlitripFlightchangeGetRequest) GetSearchOption() *FlightChangeDataQueryOption {
    return r.searchOption
}
