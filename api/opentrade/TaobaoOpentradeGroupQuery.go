package opentrade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/opentrade"
)

/* 
组团购场景查询商品组团详情 
taobao.opentrade.group.query

组团购场景下，查询商品开团详情
*/
func TaobaoOpentradeGroupQuery(clt *core.SDKClient, req *opentrade.TaobaoOpentradeGroupQueryRequest, session string) (*opentrade.TaobaoOpentradeGroupQueryAPIResponse, error) {
    var resp opentrade.TaobaoOpentradeGroupQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
