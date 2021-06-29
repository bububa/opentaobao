package opentrade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/opentrade"
)

/* 
组团购场景查询团详情 
taobao.opentrade.group.detail

组团购场景下，查询团详情
*/
func TaobaoOpentradeGroupDetail(clt *core.SDKClient, req *opentrade.TaobaoOpentradeGroupDetailRequest, session string) (*opentrade.TaobaoOpentradeGroupDetailAPIResponse, error) {
    var resp opentrade.TaobaoOpentradeGroupDetailAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
