package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

/* TaobaoBusLastplaceGet
获取目的地数据
taobao.bus.lastplace.get

传入城市 获取对应的目的地 */
func TaobaoBusLastplaceGet(clt *core.SDKClient, req *bus.TaobaoBusLastplaceGetAPIRequest, session string) (*bus.TaobaoBusLastplaceGetAPIResponse, error) {
	var resp bus.TaobaoBusLastplaceGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
