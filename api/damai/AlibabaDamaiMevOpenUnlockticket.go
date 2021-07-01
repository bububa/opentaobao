package damai

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/damai"
)

/* 
大麦换验平台-第三方对外开放-票单接口unlockTicket 
alibaba.damai.mev.open.unlockticket

开放接口 解锁票单
*/
func AlibabaDamaiMevOpenUnlockticket(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenUnlockticketAPIRequest, session string) (*damai.AlibabaDamaiMevOpenUnlockticketAPIResponse, error) {
    var resp damai.AlibabaDamaiMevOpenUnlockticketAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
