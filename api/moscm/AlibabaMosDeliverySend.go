package moscm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/moscm"
)

/* 
发货 
alibaba.mos.delivery.send

订单发货填写快递单
*/
func AlibabaMosDeliverySend(clt *core.SDKClient, req *moscm.AlibabaMosDeliverySendAPIRequest, session string) (*moscm.AlibabaMosDeliverySendAPIResponse, error) {
    var resp moscm.AlibabaMosDeliverySendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
