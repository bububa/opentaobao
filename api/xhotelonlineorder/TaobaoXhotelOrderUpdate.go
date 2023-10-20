package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// TaobaoXhotelOrderUpdate 酒店订单发货接口
// taobao.xhotel.order.update
//
// 卖家确认订单或者取消订单，适用于预付、面付、信用住订单
func TaobaoXhotelOrderUpdate(clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelOrderUpdateAPIRequest, session string) (*xhotelonlineorder.TaobaoXhotelOrderUpdateAPIResponse, error) {
	var resp xhotelonlineorder.TaobaoXhotelOrderUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
