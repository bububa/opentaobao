package xhotelofficial

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelofficial"
)

// Taobaoxhotelorderofficialsettlecancel 官网信用住取消结账
// taobao.xhotel.order.official.settle.cancel
//
// 用于官网信用住取消结账
func Taobaoxhotelorderofficialsettlecancel(clt *core.SDKClient, req *xhotelofficial.TaobaoxhotelorderofficialsettlecancelAPIRequest, session string) (*xhotelofficial.TaobaoxhotelorderofficialsettlecancelAPIResponse, error) {
	var resp xhotelofficial.TaobaoxhotelorderofficialsettlecancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
