package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
单据取消接口 
taobao.qimen.order.cancel

ERP调用奇门的接口,取消创建单据操作。场景介绍：ERP主动发起取消某些创建的单据。如入库单、出库单、退货单等；所有的场景
*/
func TaobaoQimenOrderCancel(clt *core.SDKClient, req *qimen.TaobaoQimenOrderCancelAPIRequest, session string) (*qimen.TaobaoQimenOrderCancelAPIResponse, error) {
    var resp qimen.TaobaoQimenOrderCancelAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
