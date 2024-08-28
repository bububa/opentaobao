package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderStoreSwitchstatusUpdate switchstatus.update
// taobao.omniorder.store.switchstatus.update
//
// 变更门店发货、门店自提状态
func TaobaoOmniorderStoreSwitchstatusUpdate(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreSwitchstatusUpdateAPIRequest, resp *omniorder.TaobaoOmniorderStoreSwitchstatusUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
