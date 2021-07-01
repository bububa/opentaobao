package openmall

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openmall"
)

/* 
关闭OpenMall退款单 
taobao.openmall.refund.close

关闭OpenMall退款单
*/
func TaobaoOpenmallRefundClose(clt *core.SDKClient, req *openmall.TaobaoOpenmallRefundCloseAPIRequest, session string) (*openmall.TaobaoOpenmallRefundCloseAPIResponse, error) {
    var resp openmall.TaobaoOpenmallRefundCloseAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
