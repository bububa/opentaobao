package bus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/bus"
)

/* 
B2B退票申请接口 
taobao.bus.refund.set

B2B业务支持退票
*/
func TaobaoBusRefundSet(clt *core.SDKClient, req *bus.TaobaoBusRefundSetAPIRequest, session string) (*bus.TaobaoBusRefundSetAPIResponse, error) {
    var resp bus.TaobaoBusRefundSetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
