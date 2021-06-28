package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
获取最近报表生成时间 
alibaba.scbp.effect.account.date.get

获取最近报表生成时间,格式为yyyy-MM-dd
*/
func AlibabaScbpEffectAccountDateGet(clt *core.SDKClient, req *scbp.AlibabaScbpEffectAccountDateGetRequest, session string) (*scbp.AlibabaScbpEffectAccountDateGetAPIResponse, error) {
    var resp scbp.AlibabaScbpEffectAccountDateGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
