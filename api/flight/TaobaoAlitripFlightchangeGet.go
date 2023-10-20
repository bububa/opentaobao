package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Taobaoalitripflightchangeget 获取航变信息
// taobao.alitrip.flightchange.get
//
// 查询航变是为了两个目的，阿里旅行抓取到未确认航变的航变信息源时可以由代理确认航变，同时对于确认航变的航变信息也共享给代理人做本体业务使用。
func Taobaoalitripflightchangeget(clt *core.SDKClient, req *flight.TaobaoalitripflightchangegetAPIRequest, session string) (*flight.TaobaoalitripflightchangegetAPIResponse, error) {
	var resp flight.TaobaoalitripflightchangegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
