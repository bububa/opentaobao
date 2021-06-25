package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
查询品牌下的入会菜品规则 
alibaba.alsc.crm.rule.querydishrule

查询品牌下的入会菜品规则
*/
func AlibabaAlscCrmRuleQuerydishrule(clt *core.SDKClient, req *alsc.AlibabaAlscCrmRuleQuerydishruleRequest, session string) (*alsc.AlibabaAlscCrmRuleQuerydishruleResponse, error) {
    var resp alsc.AlibabaAlscCrmRuleQuerydishruleAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
