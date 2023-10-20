package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// Alibabahappytriptaxiordernotify 状态通知
// alibaba.happytrip.taxi.order.notify
//
// 当订单发生变化是供应商通过状态通知API通知欢行，欢行获取最新的订单详情和状态进行业务处理。
func Alibabahappytriptaxiordernotify(clt *core.SDKClient, req *happytrip.AlibabahappytriptaxiordernotifyAPIRequest, session string) (*happytrip.AlibabahappytriptaxiordernotifyAPIResponse, error) {
	var resp happytrip.AlibabahappytriptaxiordernotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
