package healthnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康新零售物流详情接口 API请求
alibaba.health.nr.logistics.query

对阿里健康o2o对接的商户提供查询物流单详情的能力
*/
type AlibabaHealthNrLogisticsQueryAPIRequest struct {
    model.Params
    // 订单id
    _orderId   int64
}

// 初始化AlibabaHealthNrLogisticsQueryAPIRequest对象
func NewAlibabaHealthNrLogisticsQueryRequest() *AlibabaHealthNrLogisticsQueryAPIRequest{
    return &AlibabaHealthNrLogisticsQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHealthNrLogisticsQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.health.nr.logistics.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHealthNrLogisticsQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单id
func (r *AlibabaHealthNrLogisticsQueryAPIRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHealthNrLogisticsQueryAPIRequest) GetOrderId() int64 {
    return r._orderId
}
