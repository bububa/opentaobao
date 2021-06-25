package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
订单流水通知接口 
taobao.qimen.orderprocess.report

WMS调用奇门的接口,将订单在仓库的状态回传给ERP；场景说明：仓库仓内操作状态回传给ERP, 比如打包操作完成时, 回传一个打 包完成的状态给到ERP, ERP自行决定如何处理。
*/
func TaobaoQimenOrderprocessReport(clt *core.SDKClient, req *qimen.TaobaoQimenOrderprocessReportRequest, session string) (*qimen.TaobaoQimenOrderprocessReportResponse, error) {
    var resp qimen.TaobaoQimenOrderprocessReportAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
