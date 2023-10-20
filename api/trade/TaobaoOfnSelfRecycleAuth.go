package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoOfnSelfRecycleAuth 自助回收鉴权
// taobao.ofn.self.recycle.auth
//
// 自助回收鉴权
func TaobaoOfnSelfRecycleAuth(clt *core.SDKClient, req *trade.TaobaoOfnSelfRecycleAuthAPIRequest, resp *trade.TaobaoOfnSelfRecycleAuthAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
