package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderStoreDeliverconfigGet 查询门店发货配置内容
// taobao.omniorder.store.deliverconfig.get
//
// 查询门店发货配置内容
func TaobaoOmniorderStoreDeliverconfigGet(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreDeliverconfigGetAPIRequest, resp *omniorder.TaobaoOmniorderStoreDeliverconfigGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
