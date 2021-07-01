package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
门店自送根据核销码查订单 
taobao.omniorder.dtd.query

门店自送根据核销码码查询订单信息
*/
func TaobaoOmniorderDtdQuery(clt *core.SDKClient, req *omniorder.TaobaoOmniorderDtdQueryAPIRequest, session string) (*omniorder.TaobaoOmniorderDtdQueryAPIResponse, error) {
    var resp omniorder.TaobaoOmniorderDtdQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
