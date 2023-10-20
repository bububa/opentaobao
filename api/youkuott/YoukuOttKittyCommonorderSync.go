package youkuott

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/youkuott"
)

// YoukuOttKittyCommonorderSync 运营商一般订单同步
// youku.ott.kitty.commonorder.sync
//
// 运营商一般订单同步
func YoukuOttKittyCommonorderSync(clt *core.SDKClient, req *youkuott.YoukuOttKittyCommonorderSyncAPIRequest, resp *youkuott.YoukuOttKittyCommonorderSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
