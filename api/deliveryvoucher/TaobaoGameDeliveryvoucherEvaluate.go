package deliveryvoucher

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/deliveryvoucher"
)

// Taobaogamedeliveryvoucherevaluate 卡券评价回传
// taobao.game.deliveryvoucher.evaluate
//
// 卡券ISV回传商品评价
func Taobaogamedeliveryvoucherevaluate(clt *core.SDKClient, req *deliveryvoucher.TaobaogamedeliveryvoucherevaluateAPIRequest, session string) (*deliveryvoucher.TaobaogamedeliveryvoucherevaluateAPIResponse, error) {
	var resp deliveryvoucher.TaobaogamedeliveryvoucherevaluateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
