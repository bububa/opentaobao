package maitix

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/maitix"
)

/* 
计算渠道用户下单快递费 
alibaba.damai.maitix.distribution.delivery.calculate

计算渠道用户下单快递费
*/
func AlibabaDamaiMaitixDistributionDeliveryCalculate(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest, session string) (*maitix.AlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponse, error) {
    var resp maitix.AlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
