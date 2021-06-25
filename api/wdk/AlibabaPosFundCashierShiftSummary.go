package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
收银换班数据同步 
alibaba.pos.fund.cashier.shift.summary

收银换班数据同步，将每天收银换班的数据回流给商家。
*/
func AlibabaPosFundCashierShiftSummary(clt *core.SDKClient, req *wdk.AlibabaPosFundCashierShiftSummaryRequest, session string) (*wdk.AlibabaPosFundCashierShiftSummaryResponse, error) {
    var resp wdk.AlibabaPosFundCashierShiftSummaryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
