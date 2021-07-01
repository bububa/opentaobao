package alitripmerchant

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alitripmerchant"
)

/* 
星河-订单列表查询 
alitrip.merchant.galaxy.order.list.query

为C端用户提供酒店预订订单列表查询服务，包括订单支付状态、订单日期
*/
func AlitripMerchantGalaxyOrderListQuery(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyOrderListQueryAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyOrderListQueryAPIResponse, error) {
    var resp alitripmerchant.AlitripMerchantGalaxyOrderListQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
