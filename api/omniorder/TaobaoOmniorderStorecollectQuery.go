package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
全渠道门店自提根据核销码查询订单 
taobao.omniorder.storecollect.query

全渠道门店自提根据核销码查询订单
*/
func TaobaoOmniorderStorecollectQuery(clt *core.SDKClient, req *omniorder.TaobaoOmniorderStorecollectQueryAPIRequest, session string) (*omniorder.TaobaoOmniorderStorecollectQueryAPIResponse, error) {
    var resp omniorder.TaobaoOmniorderStorecollectQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
