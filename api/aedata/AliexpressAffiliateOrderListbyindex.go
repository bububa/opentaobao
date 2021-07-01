package aedata

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aedata"
)

/* 
AE联盟推广者订单查询接口-按游标索引查询 
aliexpress.affiliate.order.listbyindex

AE联盟推广者订单按游标查询接口
*/
func AliexpressAffiliateOrderListbyindex(clt *core.SDKClient, req *aedata.AliexpressAffiliateOrderListbyindexAPIRequest, session string) (*aedata.AliexpressAffiliateOrderListbyindexAPIResponse, error) {
    var resp aedata.AliexpressAffiliateOrderListbyindexAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
