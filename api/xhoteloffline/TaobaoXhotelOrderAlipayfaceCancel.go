package xhoteloffline

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhoteloffline"
)

// Taobaoxhotelorderalipayfacecancel 线下信用住订单取消
// taobao.xhotel.order.alipayface.cancel
//
// 提供给卖家进行线下信用住的订单取消。此接口仅仅支持线下信用住订单的取消
func Taobaoxhotelorderalipayfacecancel(clt *core.SDKClient, req *xhoteloffline.TaobaoxhotelorderalipayfacecancelAPIRequest, session string) (*xhoteloffline.TaobaoxhotelorderalipayfacecancelAPIResponse, error) {
	var resp xhoteloffline.TaobaoxhotelorderalipayfacecancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
