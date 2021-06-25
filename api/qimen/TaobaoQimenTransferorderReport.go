package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
调拨单通知 
taobao.qimen.transferorder.report

调拨单通知
*/
func TaobaoQimenTransferorderReport(clt *core.SDKClient, req *qimen.TaobaoQimenTransferorderReportRequest, session string) (*qimen.TaobaoQimenTransferorderReportResponse, error) {
    var resp qimen.TaobaoQimenTransferorderReportAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
