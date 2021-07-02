package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// TaobaoXhotelFutureSoftmodify 未来酒店信息下发
// taobao.xhotel.future.softmodify
//
// 未来酒店信息下发，包含PMS订单查询和自助入住
func TaobaoXhotelFutureSoftmodify(clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelFutureSoftmodifyAPIRequest, session string) (*xhotelonlineorder.TaobaoXhotelFutureSoftmodifyAPIResponse, error) {
	var resp xhotelonlineorder.TaobaoXhotelFutureSoftmodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
