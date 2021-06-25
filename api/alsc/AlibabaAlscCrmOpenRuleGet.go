package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
查询规则 
alibaba.alsc.crm.open.rule.get

查询会员规则
*/
func AlibabaAlscCrmOpenRuleGet(clt *core.SDKClient, req *alsc.AlibabaAlscCrmOpenRuleGetRequest, session string) (*alsc.AlibabaAlscCrmOpenRuleGetResponse, error) {
    var resp alsc.AlibabaAlscCrmOpenRuleGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
