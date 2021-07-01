package icbulogistics

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbulogistics"
)

/* 
订单详细信息(面单及仓库信息) 
alibaba.onetouch.logistics.express.order.detail.get

订单详细信息(面单及仓库信息)
*/
func AlibabaOnetouchLogisticsExpressOrderDetailGet(clt *core.SDKClient, req *icbulogistics.AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest, session string) (*icbulogistics.AlibabaOnetouchLogisticsExpressOrderDetailGetAPIResponse, error) {
    var resp icbulogistics.AlibabaOnetouchLogisticsExpressOrderDetailGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
