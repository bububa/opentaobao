package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销商合约生产 APIRequest
alibaba.wtt.order.contract.subscribe

分销商合约生产
*/
type AlibabaWttOrderContractSubscribeRequest struct {
    model.Params

    // 分销商合约生产
    distributionOrderModel   *DistributionOrderModel 

}

func NewAlibabaWttOrderContractSubscribeRequest() *AlibabaWttOrderContractSubscribeRequest{
    return &AlibabaWttOrderContractSubscribeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWttOrderContractSubscribeRequest) GetApiMethodName() string {
    return "alibaba.wtt.order.contract.subscribe"
}

func (r AlibabaWttOrderContractSubscribeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWttOrderContractSubscribeRequest) SetDistributionOrderModel(distributionOrderModel *DistributionOrderModel) error {
    r.distributionOrderModel = distributionOrderModel
    r.Set("distribution_order_model", distributionOrderModel)
    return nil
}

func (r AlibabaWttOrderContractSubscribeRequest) GetDistributionOrderModel() *DistributionOrderModel {
    return r.distributionOrderModel
}

