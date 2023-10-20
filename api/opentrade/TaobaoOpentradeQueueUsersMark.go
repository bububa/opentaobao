package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// Taobaoopentradequeueusersmark 尖货交易可购买用户标记
// taobao.opentrade.queue.users.mark
//
// 尖货交易用户标记信息回传，根据openId标记用户可购买商品
func Taobaoopentradequeueusersmark(clt *core.SDKClient, req *opentrade.TaobaoopentradequeueusersmarkAPIRequest, session string) (*opentrade.TaobaoopentradequeueusersmarkAPIResponse, error) {
	var resp opentrade.TaobaoopentradequeueusersmarkAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
