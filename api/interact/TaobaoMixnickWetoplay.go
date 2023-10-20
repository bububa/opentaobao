package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Taobaomixnickwetoplay 微淘混淆nick转互动混淆nick
// taobao.mixnick.wetoplay
//
// 微淘应用的混淆nick转为互动类型混淆nick
func Taobaomixnickwetoplay(clt *core.SDKClient, req *interact.TaobaomixnickwetoplayAPIRequest, session string) (*interact.TaobaomixnickwetoplayAPIResponse, error) {
	var resp interact.TaobaomixnickwetoplayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
