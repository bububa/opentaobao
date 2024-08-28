package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderStoreSwitchstatusGet switchstatus.get
// taobao.omniorder.store.switchstatus.get
//
// 查询门店发货、门店自提状态
func TaobaoOmniorderStoreSwitchstatusGet(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreSwitchstatusGetAPIRequest, resp *omniorder.TaobaoOmniorderStoreSwitchstatusGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
