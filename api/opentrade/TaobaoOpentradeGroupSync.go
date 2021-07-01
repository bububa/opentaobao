package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

/* TaobaoOpentradeGroupSync
组团购场景创建或更新组团活动
taobao.opentrade.group.sync

组团购场景中创建团购活动 */
func TaobaoOpentradeGroupSync(clt *core.SDKClient, req *opentrade.TaobaoOpentradeGroupSyncAPIRequest, session string) (*opentrade.TaobaoOpentradeGroupSyncAPIResponse, error) {
	var resp opentrade.TaobaoOpentradeGroupSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
