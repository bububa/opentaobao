package store

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/store"
)

// TaobaoPlaceStoregroupCreate 商户门店库创建接口
// taobao.place.storegroup.create
//
// 用于商家创建线下门店库
func TaobaoPlaceStoregroupCreate(ctx context.Context, clt *core.SDKClient, req *store.TaobaoPlaceStoregroupCreateAPIRequest, resp *store.TaobaoPlaceStoregroupCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
