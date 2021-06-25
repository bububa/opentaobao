package car

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/car"
)

/* 
航旅国内租车订单状态更新 
taobao.alitrip.domestic.rent.car.status.update

航旅国内租车订单状态更新
*/
func TaobaoAlitripDomesticRentCarStatusUpdate(clt *core.SDKClient, req *car.TaobaoAlitripDomesticRentCarStatusUpdateRequest, session string) (*car.TaobaoAlitripDomesticRentCarStatusUpdateResponse, error) {
    var resp car.TaobaoAlitripDomesticRentCarStatusUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
