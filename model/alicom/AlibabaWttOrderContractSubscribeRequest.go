package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销商合约生产 API请求
alibaba.wtt.order.contract.subscribe

分销商合约生产
*/
type AlibabaWttOrderContractSubscribeRequest struct {
    model.Params
    // 分销商合约生产
    distributionOrderModel   *DistributionOrderModel
}

// 初始化AlibabaWttOrderContractSubscribeRequest对象
func NewAlibabaWttOrderContractSubscribeRequest() *AlibabaWttOrderContractSubscribeRequest{
    return &AlibabaWttOrderContractSubscribeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWttOrderContractSubscribeRequest) GetApiMethodName() string {
    return "alibaba.wtt.order.contract.subscribe"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWttOrderContractSubscribeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DistributionOrderModel Setter
// 分销商合约生产
func (r *AlibabaWttOrderContractSubscribeRequest) SetDistributionOrderModel(distributionOrderModel *DistributionOrderModel) error {
    r.distributionOrderModel = distributionOrderModel
    r.Set("distribution_order_model", distributionOrderModel)
    return nil
}

// DistributionOrderModel Getter
func (r AlibabaWttOrderContractSubscribeRequest) GetDistributionOrderModel() *DistributionOrderModel {
    return r.distributionOrderModel
}
