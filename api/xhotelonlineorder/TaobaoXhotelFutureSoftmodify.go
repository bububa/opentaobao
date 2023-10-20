package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// Taobaoxhotelfuturesoftmodify 未来酒店信息下发
// taobao.xhotel.future.softmodify
//
// 未来酒店信息下发，包含PMS订单查询和自助入住
func Taobaoxhotelfuturesoftmodify(clt *core.SDKClient, req *xhotelonlineorder.TaobaoxhotelfuturesoftmodifyAPIRequest, session string) (*xhotelonlineorder.TaobaoxhotelfuturesoftmodifyAPIResponse, error) {
	var resp xhotelonlineorder.TaobaoxhotelfuturesoftmodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
