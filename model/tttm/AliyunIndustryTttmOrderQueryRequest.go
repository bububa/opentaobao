package tttm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天天特卖数字工厂订单获取 API请求
aliyun.industry.tttm.order.query

获取阿里云数字工厂内天天特卖业务的订单
*/
type AliyunIndustryTttmOrderQueryRequest struct {
    model.Params
    // 订单号
    _orderId   string
    // 外部采购单号
    _externalId   string
}

// 初始化AliyunIndustryTttmOrderQueryRequest对象
func NewAliyunIndustryTttmOrderQueryRequest() *AliyunIndustryTttmOrderQueryRequest{
    return &AliyunIndustryTttmOrderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunIndustryTttmOrderQueryRequest) GetApiMethodName() string {
    return "aliyun.industry.tttm.order.query"
}

// IRequest interface 方法, 获取API参数
func (r AliyunIndustryTttmOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单号
func (r *AliyunIndustryTttmOrderQueryRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AliyunIndustryTttmOrderQueryRequest) GetOrderId() string {
    return r._orderId
}
// ExternalId Setter
// 外部采购单号
func (r *AliyunIndustryTttmOrderQueryRequest) SetExternalId(_externalId string) error {
    r._externalId = _externalId
    r.Set("external_id", _externalId)
    return nil
}

// ExternalId Getter
func (r AliyunIndustryTttmOrderQueryRequest) GetExternalId() string {
    return r._externalId
}
