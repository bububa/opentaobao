package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobustvmqueryorderget 线下自助机查询订单信息
// taobao.bus.tvmqueryorder.get
//
// 查询订单详情
func Taobaobustvmqueryorderget(clt *core.SDKClient, req *bus.TaobaobustvmqueryordergetAPIRequest, session string) (*bus.TaobaobustvmqueryordergetAPIResponse, error) {
	var resp bus.TaobaobustvmqueryordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
