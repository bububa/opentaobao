package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
查询推广账户等级 
alibaba.scbp.ad.account.level.get

查询推广账户等级
*/
func AlibabaScbpAdAccountLevelGet(clt *core.SDKClient, req *scbp.AlibabaScbpAdAccountLevelGetAPIRequest, session string) (*scbp.AlibabaScbpAdAccountLevelGetAPIResponse, error) {
    var resp scbp.AlibabaScbpAdAccountLevelGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
