package xhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotel"
)

// TaobaoXhotelOrderHotelsignQuery 获取直连酒店（客栈）签约上线进度信息
// taobao.xhotel.order.hotelsign.query
//
// 获取直连酒店（客栈）签约上线进度信息
func TaobaoXhotelOrderHotelsignQuery(clt *core.SDKClient, req *xhotel.TaobaoXhotelOrderHotelsignQueryAPIRequest, resp *xhotel.TaobaoXhotelOrderHotelsignQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
