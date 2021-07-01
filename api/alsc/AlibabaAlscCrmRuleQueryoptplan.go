package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
查询运营计划 
alibaba.alsc.crm.rule.queryoptplan

查询运营计划
*/
func AlibabaAlscCrmRuleQueryoptplan(clt *core.SDKClient, req *alsc.AlibabaAlscCrmRuleQueryoptplanAPIRequest, session string) (*alsc.AlibabaAlscCrmRuleQueryoptplanAPIResponse, error) {
    var resp alsc.AlibabaAlscCrmRuleQueryoptplanAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
