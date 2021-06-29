package mtopopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
手淘下单能力开放 API请求
alibaba.interact.sensor.trade.buy

交易流程鉴权
*/
type AlibabaInteractSensorTradeBuyRequest struct {
    model.Params
    // 系统自动生成
    _id   string
}

// 初始化AlibabaInteractSensorTradeBuyRequest对象
func NewAlibabaInteractSensorTradeBuyRequest() *AlibabaInteractSensorTradeBuyRequest{
    return &AlibabaInteractSensorTradeBuyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorTradeBuyRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.trade.buy"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorTradeBuyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 系统自动生成
func (r *AlibabaInteractSensorTradeBuyRequest) SetId(_id string) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r AlibabaInteractSensorTradeBuyRequest) GetId() string {
    return r._id
}
