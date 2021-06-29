package icbudropshipping

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里巴巴订单物流轨迹查询 API请求
alibaba.order.logistics.tracking.get

阿里巴巴订单物流轨迹查询
*/
type AlibabaOrderLogisticsTrackingGetRequest struct {
    model.Params
    // order id
    _tradeId   int64
}

// 初始化AlibabaOrderLogisticsTrackingGetRequest对象
func NewAlibabaOrderLogisticsTrackingGetRequest() *AlibabaOrderLogisticsTrackingGetRequest{
    return &AlibabaOrderLogisticsTrackingGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaOrderLogisticsTrackingGetRequest) GetApiMethodName() string {
    return "alibaba.order.logistics.tracking.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaOrderLogisticsTrackingGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TradeId Setter
// order id
func (r *AlibabaOrderLogisticsTrackingGetRequest) SetTradeId(_tradeId int64) error {
    r._tradeId = _tradeId
    r.Set("trade_id", _tradeId)
    return nil
}

// TradeId Getter
func (r AlibabaOrderLogisticsTrackingGetRequest) GetTradeId() int64 {
    return r._tradeId
}
