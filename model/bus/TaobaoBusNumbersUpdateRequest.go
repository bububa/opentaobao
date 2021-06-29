package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票车次更新服务 API请求
taobao.bus.numbers.update

用于汽车票车次信息的新增、更新和逻辑删除
*/
type TaobaoBusNumbersUpdateRequest struct {
    model.Params
    // 请求参数
    _paramTopBusNumberUpdateRQ   *TopBusNumberUpdateRq
}

// 初始化TaobaoBusNumbersUpdateRequest对象
func NewTaobaoBusNumbersUpdateRequest() *TaobaoBusNumbersUpdateRequest{
    return &TaobaoBusNumbersUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusNumbersUpdateRequest) GetApiMethodName() string {
    return "taobao.bus.numbers.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusNumbersUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamTopBusNumberUpdateRQ Setter
// 请求参数
func (r *TaobaoBusNumbersUpdateRequest) SetParamTopBusNumberUpdateRQ(_paramTopBusNumberUpdateRQ *TopBusNumberUpdateRq) error {
    r._paramTopBusNumberUpdateRQ = _paramTopBusNumberUpdateRQ
    r.Set("param_top_bus_number_update_r_q", _paramTopBusNumberUpdateRQ)
    return nil
}

// ParamTopBusNumberUpdateRQ Getter
func (r TaobaoBusNumbersUpdateRequest) GetParamTopBusNumberUpdateRQ() *TopBusNumberUpdateRq {
    return r._paramTopBusNumberUpdateRQ
}
