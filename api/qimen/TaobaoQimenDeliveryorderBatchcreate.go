package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
发货单创建批量接口 
taobao.qimen.deliveryorder.batchcreate

ERP调用接口，将发货信息批量推送给WMS
*/
func TaobaoQimenDeliveryorderBatchcreate(clt *core.SDKClient, req *qimen.TaobaoQimenDeliveryorderBatchcreateRequest, session string) (*qimen.TaobaoQimenDeliveryorderBatchcreateAPIResponse, error) {
    var resp qimen.TaobaoQimenDeliveryorderBatchcreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
