package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
单据挂起（恢复）接口 
taobao.qimen.order.pending

ERP调用奇门的接口,挂起某些创建的单据;场景介绍：ERP主动发起挂起（恢复）某些创建的单据，如入库单、出库单、退货单等
*/
func TaobaoQimenOrderPending(clt *core.SDKClient, req *qimen.TaobaoQimenOrderPendingRequest, session string) (*qimen.TaobaoQimenOrderPendingResponse, error) {
    var resp qimen.TaobaoQimenOrderPendingAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
