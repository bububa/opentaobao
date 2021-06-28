package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
库存异动通知接口 
taobao.qimen.stockchange.report

WMS调用奇门的接口,将库存异动信息信息回传给ERP
*/
func TaobaoQimenStockchangeReport(clt *core.SDKClient, req *qimen.TaobaoQimenStockchangeReportRequest, session string) (*qimen.TaobaoQimenStockchangeReportAPIResponse, error) {
    var resp qimen.TaobaoQimenStockchangeReportAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
