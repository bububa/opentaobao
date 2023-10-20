package car

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/car"
)

// Taobaoalitripdomesticrentcarstatusupdate 航旅国内租车订单状态更新
// taobao.alitrip.domestic.rent.car.status.update
//
// 航旅国内租车订单状态更新
func Taobaoalitripdomesticrentcarstatusupdate(clt *core.SDKClient, req *car.TaobaoalitripdomesticrentcarstatusupdateAPIRequest, session string) (*car.TaobaoalitripdomesticrentcarstatusupdateAPIResponse, error) {
	var resp car.TaobaoalitripdomesticrentcarstatusupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
