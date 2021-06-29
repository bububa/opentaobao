package icbulogistics

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbulogistics"
)

/* 
快递下单 
alibaba.onetouch.logistics.express.logistics.order.create

快递下单
*/
func AlibabaOnetouchLogisticsExpressLogisticsOrderCreate(clt *core.SDKClient, req *icbulogistics.AlibabaOnetouchLogisticsExpressLogisticsOrderCreateRequest, session string) (*icbulogistics.AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIResponse, error) {
    var resp icbulogistics.AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
