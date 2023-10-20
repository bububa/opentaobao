package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderStoreSwitchstatusUpdate switchstatus.update
// taobao.omniorder.store.switchstatus.update
//
// 变更门店发货、门店自提状态
func TaobaoOmniorderStoreSwitchstatusUpdate(clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreSwitchstatusUpdateAPIRequest, resp *omniorder.TaobaoOmniorderStoreSwitchstatusUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
