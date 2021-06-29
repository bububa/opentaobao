package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票车次查询 API请求
taobao.bus.busnumber.get

提供汽车票车次查询服务
*/
type TaobaoBusBusnumberGetRequest struct {
    model.Params
    // 车次查询入参
    _paramBusNumberSearchRQ   *BusNumberSearchRq
}

// 初始化TaobaoBusBusnumberGetRequest对象
func NewTaobaoBusBusnumberGetRequest() *TaobaoBusBusnumberGetRequest{
    return &TaobaoBusBusnumberGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusBusnumberGetRequest) GetApiMethodName() string {
    return "taobao.bus.busnumber.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusBusnumberGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamBusNumberSearchRQ Setter
// 车次查询入参
func (r *TaobaoBusBusnumberGetRequest) SetParamBusNumberSearchRQ(_paramBusNumberSearchRQ *BusNumberSearchRq) error {
    r._paramBusNumberSearchRQ = _paramBusNumberSearchRQ
    r.Set("param_bus_number_search_r_q", _paramBusNumberSearchRQ)
    return nil
}

// ParamBusNumberSearchRQ Getter
func (r TaobaoBusBusnumberGetRequest) GetParamBusNumberSearchRQ() *BusNumberSearchRq {
    return r._paramBusNumberSearchRQ
}
