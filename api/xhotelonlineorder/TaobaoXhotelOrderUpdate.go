package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// Taobaoxhotelorderupdate 酒店订单发货接口
// taobao.xhotel.order.update
//
// 卖家确认订单或者取消订单，适用于预付、面付、信用住订单
func Taobaoxhotelorderupdate(clt *core.SDKClient, req *xhotelonlineorder.TaobaoxhotelorderupdateAPIRequest, session string) (*xhotelonlineorder.TaobaoxhotelorderupdateAPIResponse, error) {
	var resp xhotelonlineorder.TaobaoxhotelorderupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
