package store

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/store"
)

// TaobaoPlaceStorerelatesubDelete 门店和子门店关系删除
// taobao.place.storerelatesub.delete
//
// 门店和子门店关系删除
func TaobaoPlaceStorerelatesubDelete(ctx context.Context, clt *core.SDKClient, req *store.TaobaoPlaceStorerelatesubDeleteAPIRequest, resp *store.TaobaoPlaceStorerelatesubDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
