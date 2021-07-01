package openmall

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openmall"
)

/* 
渲染订单价格 
taobao.openmall.trade.render

请求渲染订单价格
*/
func TaobaoOpenmallTradeRender(clt *core.SDKClient, req *openmall.TaobaoOpenmallTradeRenderAPIRequest, session string) (*openmall.TaobaoOpenmallTradeRenderAPIResponse, error) {
    var resp openmall.TaobaoOpenmallTradeRenderAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
