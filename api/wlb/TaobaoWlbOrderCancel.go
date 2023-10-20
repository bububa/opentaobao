package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// Taobaowlbordercancel 取消物流宝订单
// taobao.wlb.order.cancel
//
// 取消物流宝订单
func Taobaowlbordercancel(clt *core.SDKClient, req *wlb.TaobaowlbordercancelAPIRequest, session string) (*wlb.TaobaowlbordercancelAPIResponse, error) {
	var resp wlb.TaobaowlbordercancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
