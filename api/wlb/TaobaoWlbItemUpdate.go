package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbItemUpdate 物流宝商品修改
// taobao.wlb.item.update
//
// 修改物流宝商品信息
func TaobaoWlbItemUpdate(clt *core.SDKClient, req *wlb.TaobaoWlbItemUpdateAPIRequest, resp *wlb.TaobaoWlbItemUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
