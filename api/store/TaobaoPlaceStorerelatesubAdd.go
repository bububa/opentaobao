package store

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/store"
)

// TaobaoPlaceStorerelatesubAdd 门店和子门店关系新增
// taobao.place.storerelatesub.add
//
// 门店和子门店关系新增
func TaobaoPlaceStorerelatesubAdd(ctx context.Context, clt *core.SDKClient, req *store.TaobaoPlaceStorerelatesubAddAPIRequest, resp *store.TaobaoPlaceStorerelatesubAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
