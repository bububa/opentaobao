package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Taobaoofnselfrecycleauth 自助回收鉴权
// taobao.ofn.self.recycle.auth
//
// 自助回收鉴权
func Taobaoofnselfrecycleauth(clt *core.SDKClient, req *trade.TaobaoofnselfrecycleauthAPIRequest, session string) (*trade.TaobaoofnselfrecycleauthAPIResponse, error) {
	var resp trade.TaobaoofnselfrecycleauthAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
