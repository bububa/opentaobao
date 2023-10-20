package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// Taobaowlbwmsinventoryprofitlossget 通过订单列表批量获取库存损益单信息
// taobao.wlb.wms.inventory.profitloss.get
//
// 通过订单列表批量获取库存损益单信息
func Taobaowlbwmsinventoryprofitlossget(clt *core.SDKClient, req *wms.TaobaowlbwmsinventoryprofitlossgetAPIRequest, session string) (*wms.TaobaowlbwmsinventoryprofitlossgetAPIResponse, error) {
	var resp wms.TaobaowlbwmsinventoryprofitlossgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
