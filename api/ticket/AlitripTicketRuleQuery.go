package ticket

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ticket"
)

/* 
【门票API2.0】门票规则信息查询接口 
alitrip.ticket.rule.query

门票规则信息查询接口：返回商家上传的门票规则信息
*/
func AlitripTicketRuleQuery(clt *core.SDKClient, req *ticket.AlitripTicketRuleQueryRequest, session string) (*ticket.AlitripTicketRuleQueryResponse, error) {
    var resp ticket.AlitripTicketRuleQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
