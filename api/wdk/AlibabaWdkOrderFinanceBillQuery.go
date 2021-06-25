package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
资金合规商家账单 
alibaba.wdk.order.finance.bill.query

拉取资金合规商家账单
*/
func AlibabaWdkOrderFinanceBillQuery(clt *core.SDKClient, req *wdk.AlibabaWdkOrderFinanceBillQueryRequest, session string) (*wdk.AlibabaWdkOrderFinanceBillQueryResponse, error) {
    var resp wdk.AlibabaWdkOrderFinanceBillQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
