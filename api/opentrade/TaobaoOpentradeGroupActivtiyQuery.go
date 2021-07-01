package opentrade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/opentrade"
)

/* 
组团活动信息查询 
taobao.opentrade.group.activtiy.query

组团购场景下，团购活动信息查询
*/
func TaobaoOpentradeGroupActivtiyQuery(clt *core.SDKClient, req *opentrade.TaobaoOpentradeGroupActivtiyQueryAPIRequest, session string) (*opentrade.TaobaoOpentradeGroupActivtiyQueryAPIResponse, error) {
    var resp opentrade.TaobaoOpentradeGroupActivtiyQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
