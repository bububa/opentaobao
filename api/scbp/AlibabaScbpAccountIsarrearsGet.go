package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
查询关键词推广账户是否欠款 
alibaba.scbp.account.isarrears.get

查询关键词推广账户是否欠款
*/
func AlibabaScbpAccountIsarrearsGet(clt *core.SDKClient, req *scbp.AlibabaScbpAccountIsarrearsGetRequest, session string) (*scbp.AlibabaScbpAccountIsarrearsGetAPIResponse, error) {
    var resp scbp.AlibabaScbpAccountIsarrearsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
