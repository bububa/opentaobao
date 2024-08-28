package store

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/store"
)

// TaobaoPlaceStoregroupUpdate 门店库修改基本信息
// taobao.place.storegroup.update
//
// 门店库修改基本信息
func TaobaoPlaceStoregroupUpdate(ctx context.Context, clt *core.SDKClient, req *store.TaobaoPlaceStoregroupUpdateAPIRequest, resp *store.TaobaoPlaceStoregroupUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
