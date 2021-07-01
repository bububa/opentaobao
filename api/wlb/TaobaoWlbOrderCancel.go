package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

/* TaobaoWlbOrderCancel
取消物流宝订单
taobao.wlb.order.cancel

取消物流宝订单 */
func TaobaoWlbOrderCancel(clt *core.SDKClient, req *wlb.TaobaoWlbOrderCancelAPIRequest, session string) (*wlb.TaobaoWlbOrderCancelAPIResponse, error) {
	var resp wlb.TaobaoWlbOrderCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
