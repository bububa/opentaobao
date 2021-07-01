package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
发货单缺货通知接口 
taobao.qimen.itemlack.report

WMS调用奇门的接口,将商家在库某商品缺货的信息回传给ERP
*/
func TaobaoQimenItemlackReport(clt *core.SDKClient, req *qimen.TaobaoQimenItemlackReportAPIRequest, session string) (*qimen.TaobaoQimenItemlackReportAPIResponse, error) {
    var resp qimen.TaobaoQimenItemlackReportAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
