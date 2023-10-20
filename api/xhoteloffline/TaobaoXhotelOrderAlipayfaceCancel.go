package xhoteloffline

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhoteloffline"
)

// TaobaoXhotelOrderAlipayfaceCancel 线下信用住订单取消
// taobao.xhotel.order.alipayface.cancel
//
// 提供给卖家进行线下信用住的订单取消。此接口仅仅支持线下信用住订单的取消
func TaobaoXhotelOrderAlipayfaceCancel(clt *core.SDKClient, req *xhoteloffline.TaobaoXhotelOrderAlipayfaceCancelAPIRequest, resp *xhoteloffline.TaobaoXhotelOrderAlipayfaceCancelAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
