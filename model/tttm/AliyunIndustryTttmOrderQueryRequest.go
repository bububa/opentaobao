package tttm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天天特卖数字工厂订单获取 APIRequest
aliyun.industry.tttm.order.query

获取阿里云数字工厂内天天特卖业务的订单
*/
type AliyunIndustryTttmOrderQueryRequest struct {
    model.Params

    // 订单号
    orderId   string 

    // 外部采购单号
    externalId   string 

}

func NewAliyunIndustryTttmOrderQueryRequest() *AliyunIndustryTttmOrderQueryRequest{
    return &AliyunIndustryTttmOrderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AliyunIndustryTttmOrderQueryRequest) GetApiMethodName() string {
    return "aliyun.industry.tttm.order.query"
}

func (r AliyunIndustryTttmOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliyunIndustryTttmOrderQueryRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AliyunIndustryTttmOrderQueryRequest) GetOrderId() string {
    return r.orderId
}

func (r *AliyunIndustryTttmOrderQueryRequest) SetExternalId(externalId string) error {
    r.externalId = externalId
    r.Set("external_id", externalId)
    return nil
}

func (r AliyunIndustryTttmOrderQueryRequest) GetExternalId() string {
    return r.externalId
}

