package aedata

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aedata"
)

/* 
AE流量订单详情获取API 
aliexpress.affiliate.order.get

联盟推广订单效果获取API
*/
func AliexpressAffiliateOrderGet(clt *core.SDKClient, req *aedata.AliexpressAffiliateOrderGetAPIRequest, session string) (*aedata.AliexpressAffiliateOrderGetAPIResponse, error) {
    var resp aedata.AliexpressAffiliateOrderGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
