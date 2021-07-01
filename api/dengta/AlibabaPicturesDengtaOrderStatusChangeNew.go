package dengta

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/dengta"
)

/* 
天下秀订单状态变更通知 
alibaba.pictures.dengta.order.status.change.new

天下秀订单状态变更通知
*/
func AlibabaPicturesDengtaOrderStatusChangeNew(clt *core.SDKClient, req *dengta.AlibabaPicturesDengtaOrderStatusChangeNewAPIRequest, session string) (*dengta.AlibabaPicturesDengtaOrderStatusChangeNewAPIResponse, error) {
    var resp dengta.AlibabaPicturesDengtaOrderStatusChangeNewAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
