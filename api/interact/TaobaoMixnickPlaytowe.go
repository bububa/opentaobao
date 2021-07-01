package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

/* TaobaoMixnickPlaytowe
互动mixNick转微淘
taobao.mixnick.playtowe

微淘应用的混淆nick转为互动类型混淆nick */
func TaobaoMixnickPlaytowe(clt *core.SDKClient, req *interact.TaobaoMixnickPlaytoweAPIRequest, session string) (*interact.TaobaoMixnickPlaytoweAPIResponse, error) {
	var resp interact.TaobaoMixnickPlaytoweAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
