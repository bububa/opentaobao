package store

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/store"
)

// TaobaoPlaceStorerelatesubGet 门店和子门店关系查找
// taobao.place.storerelatesub.get
//
// 门店和子门店关系查找
func TaobaoPlaceStorerelatesubGet(ctx context.Context, clt *core.SDKClient, req *store.TaobaoPlaceStorerelatesubGetAPIRequest, resp *store.TaobaoPlaceStorerelatesubGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
