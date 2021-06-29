package icbuassurance

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbuassurance"
)

/* 
icbu信保账户信息 
alibaba.icbu.trade.assurance.account.get

icbu交易信用保障开通状态&额度信息查询
*/
func AlibabaIcbuTradeAssuranceAccountGet(clt *core.SDKClient, req *icbuassurance.AlibabaIcbuTradeAssuranceAccountGetRequest, session string) (*icbuassurance.AlibabaIcbuTradeAssuranceAccountGetAPIResponse, error) {
    var resp icbuassurance.AlibabaIcbuTradeAssuranceAccountGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
