package openmall

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openmall"
)

/* 
提交退款单留言 
taobao.openmall.refund.message.submit

OpenMall业务提交退款单留言
*/
func TaobaoOpenmallRefundMessageSubmit(clt *core.SDKClient, req *openmall.TaobaoOpenmallRefundMessageSubmitAPIRequest, session string) (*openmall.TaobaoOpenmallRefundMessageSubmitAPIResponse, error) {
    var resp openmall.TaobaoOpenmallRefundMessageSubmitAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
