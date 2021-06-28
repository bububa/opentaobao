package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
查询品牌下的会员日规则 
alibaba.alsc.crm.rule.querymdayerule

查询品牌下的会员日规则
*/
func AlibabaAlscCrmRuleQuerymdayerule(clt *core.SDKClient, req *alsc.AlibabaAlscCrmRuleQuerymdayeruleRequest, session string) (*alsc.AlibabaAlscCrmRuleQuerymdayeruleAPIResponse, error) {
    var resp alsc.AlibabaAlscCrmRuleQuerymdayeruleAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
