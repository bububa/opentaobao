package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家汽车票车次更新通知接口 API请求
taobao.bus.busnumber.set

商家汽车票车次更新后，调用该接口通知平台。
*/
type TaobaoBusBusnumberSetRequest struct {
    model.Params
    // 车次更新通知参数
    _pushParam   *TopBusNumerPushRq
}

// 初始化TaobaoBusBusnumberSetRequest对象
func NewTaobaoBusBusnumberSetRequest() *TaobaoBusBusnumberSetRequest{
    return &TaobaoBusBusnumberSetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusBusnumberSetRequest) GetApiMethodName() string {
    return "taobao.bus.busnumber.set"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusBusnumberSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PushParam Setter
// 车次更新通知参数
func (r *TaobaoBusBusnumberSetRequest) SetPushParam(_pushParam *TopBusNumerPushRq) error {
    r._pushParam = _pushParam
    r.Set("push_param", _pushParam)
    return nil
}

// PushParam Getter
func (r TaobaoBusBusnumberSetRequest) GetPushParam() *TopBusNumerPushRq {
    return r._pushParam
}
