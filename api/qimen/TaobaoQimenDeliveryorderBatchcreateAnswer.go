package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
发货单创建结果通知接口(批量) 
taobao.qimen.deliveryorder.batchcreate.answer

WMS调用接口，用于异步化的批量发货单创建结果通知。（如菜鸟发货单批量创建结果的返回）
*/
func TaobaoQimenDeliveryorderBatchcreateAnswer(clt *core.SDKClient, req *qimen.TaobaoQimenDeliveryorderBatchcreateAnswerRequest, session string) (*qimen.TaobaoQimenDeliveryorderBatchcreateAnswerResponse, error) {
    var resp qimen.TaobaoQimenDeliveryorderBatchcreateAnswerAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
