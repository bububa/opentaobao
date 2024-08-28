package interact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// TaobaoMixnickPlaytowe 互动mixNick转微淘
// taobao.mixnick.playtowe
//
// 微淘应用的混淆nick转为互动类型混淆nick
func TaobaoMixnickPlaytowe(ctx context.Context, clt *core.SDKClient, req *interact.TaobaoMixnickPlaytoweAPIRequest, resp *interact.TaobaoMixnickPlaytoweAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
