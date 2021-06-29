package opentrade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/opentrade"
)

/* 
交易开放获取商品绑定信息 
taobao.opentrade.tools.items.query

交易开放获取商品绑定信息
*/
func TaobaoOpentradeToolsItemsQuery(clt *core.SDKClient, req *opentrade.TaobaoOpentradeToolsItemsQueryRequest, session string) (*opentrade.TaobaoOpentradeToolsItemsQueryAPIResponse, error) {
    var resp opentrade.TaobaoOpentradeToolsItemsQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
