package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询分销物流单 API请求
alibaba.damai.maitix.distribution.delivery.query

渠道查询物流订单
*/
type AlibabaDamaiMaitixDistributionDeliveryQueryRequest struct {
    model.Params
    // 主订单号
    mainOrderId   string
}

// 初始化AlibabaDamaiMaitixDistributionDeliveryQueryRequest对象
func NewAlibabaDamaiMaitixDistributionDeliveryQueryRequest() *AlibabaDamaiMaitixDistributionDeliveryQueryRequest{
    return &AlibabaDamaiMaitixDistributionDeliveryQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixDistributionDeliveryQueryRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.distribution.delivery.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixDistributionDeliveryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainOrderId Setter
// 主订单号
func (r *AlibabaDamaiMaitixDistributionDeliveryQueryRequest) SetMainOrderId(mainOrderId string) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

// MainOrderId Getter
func (r AlibabaDamaiMaitixDistributionDeliveryQueryRequest) GetMainOrderId() string {
    return r.mainOrderId
}
