package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
查询今日消耗 
alibaba.scbp.account.daycost.get

查询今日消耗
*/
func AlibabaScbpAccountDaycostGet(clt *core.SDKClient, req *scbp.AlibabaScbpAccountDaycostGetRequest, session string) (*scbp.AlibabaScbpAccountDaycostGetResponse, error) {
    var resp scbp.AlibabaScbpAccountDaycostGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
