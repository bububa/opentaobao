package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// TaobaoXhotelOrderUpdate 酒店订单发货接口
// taobao.xhotel.order.update
//
// 卖家确认订单或者取消订单，适用于预付、面付、信用住订单
func TaobaoXhotelOrderUpdate(clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelOrderUpdateAPIRequest, resp *xhotelonlineorder.TaobaoXhotelOrderUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
