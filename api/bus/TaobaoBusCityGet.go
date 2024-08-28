package bus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusCityGet 城市接口
// taobao.bus.city.get
//
// 汽车票出发城市获取接口，获取所有出发城市
func TaobaoBusCityGet(ctx context.Context, clt *core.SDKClient, req *bus.TaobaoBusCityGetAPIRequest, resp *bus.TaobaoBusCityGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
