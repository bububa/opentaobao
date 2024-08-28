package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderStoreDeliverconfigUpdate 修改门店发货配置内容
// taobao.omniorder.store.deliverconfig.update
//
// 修改门店发货配置内容
func TaobaoOmniorderStoreDeliverconfigUpdate(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreDeliverconfigUpdateAPIRequest, resp *omniorder.TaobaoOmniorderStoreDeliverconfigUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
