package damai

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/damai"
)

/* 
大麦换验平台-第三方对外开放-票单接口withdrawTicket 
alibaba.damai.mev.open.withdrawticket

开放接口退票
*/
func AlibabaDamaiMevOpenWithdrawticket(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenWithdrawticketRequest, session string) (*damai.AlibabaDamaiMevOpenWithdrawticketAPIResponse, error) {
    var resp damai.AlibabaDamaiMevOpenWithdrawticketAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
