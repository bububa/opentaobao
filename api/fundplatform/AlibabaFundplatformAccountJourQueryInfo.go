package fundplatform

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fundplatform"
)

/* 
查询账户流水信息 
alibaba.fundplatform.account.jour.query.info

外部查询账户流水信息
*/
func AlibabaFundplatformAccountJourQueryInfo(clt *core.SDKClient, req *fundplatform.AlibabaFundplatformAccountJourQueryInfoRequest, session string) (*fundplatform.AlibabaFundplatformAccountJourQueryInfoAPIResponse, error) {
    var resp fundplatform.AlibabaFundplatformAccountJourQueryInfoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
