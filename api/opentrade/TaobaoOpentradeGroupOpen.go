package opentrade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/opentrade"
)

/* 
组团购场景开团 
taobao.opentrade.group.open

组团购场景下，团长开团
*/
func TaobaoOpentradeGroupOpen(clt *core.SDKClient, req *opentrade.TaobaoOpentradeGroupOpenAPIRequest, session string) (*opentrade.TaobaoOpentradeGroupOpenAPIResponse, error) {
    var resp opentrade.TaobaoOpentradeGroupOpenAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
