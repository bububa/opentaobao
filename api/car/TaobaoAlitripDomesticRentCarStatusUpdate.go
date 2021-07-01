package car

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/car"
)

/* TaobaoAlitripDomesticRentCarStatusUpdate
航旅国内租车订单状态更新
taobao.alitrip.domestic.rent.car.status.update

航旅国内租车订单状态更新 */
func TaobaoAlitripDomesticRentCarStatusUpdate(clt *core.SDKClient, req *car.TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest, session string) (*car.TaobaoAlitripDomesticRentCarStatusUpdateAPIResponse, error) {
	var resp car.TaobaoAlitripDomesticRentCarStatusUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
