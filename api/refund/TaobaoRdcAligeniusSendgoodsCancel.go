package refund

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/refund"
)

/* 
取消发货 
taobao.rdc.aligenius.sendgoods.cancel

提供商家在仅退款中发送取消发货状态
*/
func TaobaoRdcAligeniusSendgoodsCancel(clt *core.SDKClient, req *refund.TaobaoRdcAligeniusSendgoodsCancelAPIRequest, session string) (*refund.TaobaoRdcAligeniusSendgoodsCancelAPIResponse, error) {
    var resp refund.TaobaoRdcAligeniusSendgoodsCancelAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
