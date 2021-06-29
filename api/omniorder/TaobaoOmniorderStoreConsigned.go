package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
Pos端门店发货 
taobao.omniorder.store.consigned

ISV Pos端门店发货，通知星盘
*/
func TaobaoOmniorderStoreConsigned(clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreConsignedRequest, session string) (*omniorder.TaobaoOmniorderStoreConsignedAPIResponse, error) {
    var resp omniorder.TaobaoOmniorderStoreConsignedAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
