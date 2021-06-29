package retail

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/retail"
)

/* 
贩卖机货道解锁 
alibaba.retail.device.road.status.reset

贩卖机货道解锁
*/
func AlibabaRetailDeviceRoadStatusReset(clt *core.SDKClient, req *retail.AlibabaRetailDeviceRoadStatusResetRequest, session string) (*retail.AlibabaRetailDeviceRoadStatusResetAPIResponse, error) {
    var resp retail.AlibabaRetailDeviceRoadStatusResetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
