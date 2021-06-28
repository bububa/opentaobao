package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
获取删除的创意ID 
taobao.simba.creativeids.deleted.get

获取删除的创意ID
*/
func TaobaoSimbaCreativeidsDeletedGet(clt *core.SDKClient, req *simba.TaobaoSimbaCreativeidsDeletedGetRequest, session string) (*simba.TaobaoSimbaCreativeidsDeletedGetAPIResponse, error) {
    var resp simba.TaobaoSimbaCreativeidsDeletedGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
