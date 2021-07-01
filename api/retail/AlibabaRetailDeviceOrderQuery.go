package retail

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/retail"
)

/* 
贩卖机订单查询 
alibaba.retail.device.order.query

贩卖机订单查询
*/
func AlibabaRetailDeviceOrderQuery(clt *core.SDKClient, req *retail.AlibabaRetailDeviceOrderQueryAPIRequest, session string) (*retail.AlibabaRetailDeviceOrderQueryAPIResponse, error) {
    var resp retail.AlibabaRetailDeviceOrderQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
