package dengta

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/dengta"
)

/* 
天下秀订单状态变更通知 
alibaba.pictures.dengta.order.status.change

天下秀订单状态变更通知
*/
func AlibabaPicturesDengtaOrderStatusChange(clt *core.SDKClient, req *dengta.AlibabaPicturesDengtaOrderStatusChangeAPIRequest, session string) (*dengta.AlibabaPicturesDengtaOrderStatusChangeAPIResponse, error) {
    var resp dengta.AlibabaPicturesDengtaOrderStatusChangeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
