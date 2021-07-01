package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

/* TaobaoOmniorderStoreSwitchstatusUpdate
switchstatus.update
taobao.omniorder.store.switchstatus.update

变更门店发货、门店自提状态 */
func TaobaoOmniorderStoreSwitchstatusUpdate(clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreSwitchstatusUpdateAPIRequest, session string) (*omniorder.TaobaoOmniorderStoreSwitchstatusUpdateAPIResponse, error) {
	var resp omniorder.TaobaoOmniorderStoreSwitchstatusUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
