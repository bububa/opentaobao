package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// Taobaoxhotelorderalipayfacesettle 信用住订单结账接口
// taobao.xhotel.order.alipayface.settle
//
// 用于离店付订单在客人离店后，发起结账以及扣款等后续动作
func Taobaoxhotelorderalipayfacesettle(clt *core.SDKClient, req *xhotelonlineorder.TaobaoxhotelorderalipayfacesettleAPIRequest, session string) (*xhotelonlineorder.TaobaoxhotelorderalipayfacesettleAPIResponse, error) {
	var resp xhotelonlineorder.TaobaoxhotelorderalipayfacesettleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
