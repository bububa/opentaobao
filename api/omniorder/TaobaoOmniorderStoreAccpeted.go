package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
Pos端门店接单接口 
taobao.omniorder.store.accpeted

ISV Pos端门店接单，通知星盘
*/
func TaobaoOmniorderStoreAccpeted(clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreAccpetedAPIRequest, session string) (*omniorder.TaobaoOmniorderStoreAccpetedAPIResponse, error) {
    var resp omniorder.TaobaoOmniorderStoreAccpetedAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
