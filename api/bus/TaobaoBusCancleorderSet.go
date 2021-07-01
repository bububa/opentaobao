package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

/* TaobaoBusCancleorderSet
取消订单
taobao.bus.cancleorder.set

取消订单 */
func TaobaoBusCancleorderSet(clt *core.SDKClient, req *bus.TaobaoBusCancleorderSetAPIRequest, session string) (*bus.TaobaoBusCancleorderSetAPIResponse, error) {
	var resp bus.TaobaoBusCancleorderSetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
