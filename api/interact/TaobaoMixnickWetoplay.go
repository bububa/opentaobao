package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// TaobaoMixnickWetoplay 微淘混淆nick转互动混淆nick
// taobao.mixnick.wetoplay
//
// 微淘应用的混淆nick转为互动类型混淆nick
func TaobaoMixnickWetoplay(clt *core.SDKClient, req *interact.TaobaoMixnickWetoplayAPIRequest, resp *interact.TaobaoMixnickWetoplayAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
