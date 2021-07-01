package ship

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ship"
)

/* 
船票退票退款回填接口 
alitrip.ship.return.notify

此接口为接入商调用飞猪接口回填退票状态，飞猪平台给用户进行退票退款。飞猪平台保证数据幂等。
*/
func AlitripShipReturnNotify(clt *core.SDKClient, req *ship.AlitripShipReturnNotifyAPIRequest, session string) (*ship.AlitripShipReturnNotifyAPIResponse, error) {
    var resp ship.AlitripShipReturnNotifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
