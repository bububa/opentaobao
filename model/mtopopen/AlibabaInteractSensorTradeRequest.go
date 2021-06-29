package mtopopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
交易组件 API请求
alibaba.interact.sensor.trade

交易流程
*/
type AlibabaInteractSensorTradeRequest struct {
    model.Params
    // 系统自动生成
    _id   string
}

// 初始化AlibabaInteractSensorTradeRequest对象
func NewAlibabaInteractSensorTradeRequest() *AlibabaInteractSensorTradeRequest{
    return &AlibabaInteractSensorTradeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorTradeRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.trade"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorTradeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 系统自动生成
func (r *AlibabaInteractSensorTradeRequest) SetId(_id string) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r AlibabaInteractSensorTradeRequest) GetId() string {
    return r._id
}
