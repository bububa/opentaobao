package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderStoreSwitchstatusGet switchstatus.get
// taobao.omniorder.store.switchstatus.get
//
// 查询门店发货、门店自提状态
func TaobaoOmniorderStoreSwitchstatusGet(clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreSwitchstatusGetAPIRequest, resp *omniorder.TaobaoOmniorderStoreSwitchstatusGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
