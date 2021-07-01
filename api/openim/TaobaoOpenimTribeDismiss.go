package openim

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openim"
)

/* 
OPENIM群解散 
taobao.openim.tribe.dismiss

OPENIM群解散
*/
func TaobaoOpenimTribeDismiss(clt *core.SDKClient, req *openim.TaobaoOpenimTribeDismissAPIRequest, session string) (*openim.TaobaoOpenimTribeDismissAPIResponse, error) {
    var resp openim.TaobaoOpenimTribeDismissAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
