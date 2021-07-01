package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
Pos端门店拒单 
taobao.omniorder.store.refused

ISV Pos端门店拒单，通知星盘
*/
func TaobaoOmniorderStoreRefused(clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreRefusedAPIRequest, session string) (*omniorder.TaobaoOmniorderStoreRefusedAPIResponse, error) {
    var resp omniorder.TaobaoOmniorderStoreRefusedAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
