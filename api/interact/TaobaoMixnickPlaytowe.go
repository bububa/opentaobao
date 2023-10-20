package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Taobaomixnickplaytowe 互动mixNick转微淘
// taobao.mixnick.playtowe
//
// 微淘应用的混淆nick转为互动类型混淆nick
func Taobaomixnickplaytowe(clt *core.SDKClient, req *interact.TaobaomixnickplaytoweAPIRequest, session string) (*interact.TaobaomixnickplaytoweAPIResponse, error) {
	var resp interact.TaobaomixnickplaytoweAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
