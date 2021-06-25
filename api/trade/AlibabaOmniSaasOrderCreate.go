package trade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/trade"
)

/* 
订单创建接口 
alibaba.omni.saas.order.create

服务商利用现有的saas系统和阿里完成交易系统的对接
*/
func AlibabaOmniSaasOrderCreate(clt *core.SDKClient, req *trade.AlibabaOmniSaasOrderCreateRequest, session string) (*trade.AlibabaOmniSaasOrderCreateResponse, error) {
    var resp trade.AlibabaOmniSaasOrderCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
