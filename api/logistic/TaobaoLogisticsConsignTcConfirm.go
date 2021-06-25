package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
通知交易确认发货接口 
taobao.logistics.consign.tc.confirm

下述业务场景可以使用此接口通知相关的交易订单发货：
1、发货过程分为多段操作，在确定发货之前，不需要通知交易，当货确认已发出之后，才通知交易发货。
2、发货过程涉及到多个订单，其中一个订单是跟实操的发货操作同步的，剩下的订单，需要在实操的订单发货之后，一并通知交易发货。
*/
func TaobaoLogisticsConsignTcConfirm(clt *core.SDKClient, req *logistic.TaobaoLogisticsConsignTcConfirmRequest, session string) (*logistic.TaobaoLogisticsConsignTcConfirmResponse, error) {
    var resp logistic.TaobaoLogisticsConsignTcConfirmAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
