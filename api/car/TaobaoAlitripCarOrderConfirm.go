package car

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/car"
)

/* TaobaoAlitripCarOrderConfirm
司机应答API
taobao.alitrip.car.order.confirm

航旅事业群-度假事业部-旅行用车项目组对外部服务商提供的司机应答回调接口 */
func TaobaoAlitripCarOrderConfirm(clt *core.SDKClient, req *car.TaobaoAlitripCarOrderConfirmAPIRequest, session string) (*car.TaobaoAlitripCarOrderConfirmAPIResponse, error) {
	var resp car.TaobaoAlitripCarOrderConfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
