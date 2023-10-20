package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// TaobaoXhotelOrderUpdateConfirmcode 推送及更新订单确认号
// taobao.xhotel.order.update.confirmcode
//
// 商家拿到订单确认号后，异步推送给飞猪或更新给飞猪。订单确认号用于到店查无时的紧急查单依据。
func TaobaoXhotelOrderUpdateConfirmcode(clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelOrderUpdateConfirmcodeAPIRequest, session string) (*xhotelonlineorder.TaobaoXhotelOrderUpdateConfirmcodeAPIResponse, error) {
	var resp xhotelonlineorder.TaobaoXhotelOrderUpdateConfirmcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
