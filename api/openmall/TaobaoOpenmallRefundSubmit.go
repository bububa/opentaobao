package openmall

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openmall"
)

/* 
提交OpenMall退款单物流 
taobao.openmall.refund.submit

提交OpenMall退款单物流
*/
func TaobaoOpenmallRefundSubmit(clt *core.SDKClient, req *openmall.TaobaoOpenmallRefundSubmitRequest, session string) (*openmall.TaobaoOpenmallRefundSubmitAPIResponse, error) {
    var resp openmall.TaobaoOpenmallRefundSubmitAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
