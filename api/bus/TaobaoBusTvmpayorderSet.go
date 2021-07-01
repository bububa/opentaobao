package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

/* TaobaoBusTvmpayorderSet
自助机条形码被动支付
taobao.bus.tvmpayorder.set

汽车票线下自助机条形码支付 */
func TaobaoBusTvmpayorderSet(clt *core.SDKClient, req *bus.TaobaoBusTvmpayorderSetAPIRequest, session string) (*bus.TaobaoBusTvmpayorderSetAPIResponse, error) {
	var resp bus.TaobaoBusTvmpayorderSetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
