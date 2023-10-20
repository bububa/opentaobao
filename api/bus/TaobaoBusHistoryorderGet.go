package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobushistoryorderget 历史订单查询（对账）
// taobao.bus.historyorder.get
//
// 历史订单查询，对账接口
func Taobaobushistoryorderget(clt *core.SDKClient, req *bus.TaobaobushistoryordergetAPIRequest, session string) (*bus.TaobaobushistoryordergetAPIResponse, error) {
	var resp bus.TaobaobushistoryordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
