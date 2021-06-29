package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询分销物流单 APIRequest
alibaba.damai.maitix.distribution.delivery.query

渠道查询物流订单
*/
type AlibabaDamaiMaitixDistributionDeliveryQueryRequest struct {
    model.Params

    // 主订单号
    mainOrderId   string 

}

func NewAlibabaDamaiMaitixDistributionDeliveryQueryRequest() *AlibabaDamaiMaitixDistributionDeliveryQueryRequest{
    return &AlibabaDamaiMaitixDistributionDeliveryQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMaitixDistributionDeliveryQueryRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.distribution.delivery.query"
}

func (r AlibabaDamaiMaitixDistributionDeliveryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMaitixDistributionDeliveryQueryRequest) SetMainOrderId(mainOrderId string) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

func (r AlibabaDamaiMaitixDistributionDeliveryQueryRequest) GetMainOrderId() string {
    return r.mainOrderId
}

