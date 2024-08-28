package store

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/store"
)

// TaobaoPlaceStoregroupDelete 删除门店库
// taobao.place.storegroup.delete
//
// 删除门店库
func TaobaoPlaceStoregroupDelete(ctx context.Context, clt *core.SDKClient, req *store.TaobaoPlaceStoregroupDeleteAPIRequest, resp *store.TaobaoPlaceStoregroupDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
