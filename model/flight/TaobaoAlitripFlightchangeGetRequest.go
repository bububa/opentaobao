package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取航变信息 APIRequest
taobao.alitrip.flightchange.get

查询航变是为了两个目的，阿里旅行抓取到未确认航变的航变信息源时可以由代理确认航变，同时对于确认航变的航变信息也共享给代理人做本体业务使用。
*/
type TaobaoAlitripFlightchangeGetRequest struct {
    model.Params

    // 查询信息封装
    searchOption   *FlightChangeDataQueryOption 

}

func NewTaobaoAlitripFlightchangeGetRequest() *TaobaoAlitripFlightchangeGetRequest{
    return &TaobaoAlitripFlightchangeGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripFlightchangeGetRequest) GetApiMethodName() string {
    return "taobao.alitrip.flightchange.get"
}

func (r TaobaoAlitripFlightchangeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripFlightchangeGetRequest) SetSearchOption(searchOption *FlightChangeDataQueryOption) error {
    r.searchOption = searchOption
    r.Set("search_option", searchOption)
    return nil
}

func (r TaobaoAlitripFlightchangeGetRequest) GetSearchOption() *FlightChangeDataQueryOption {
    return r.searchOption
}

