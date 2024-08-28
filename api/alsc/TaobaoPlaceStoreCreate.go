package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// TaobaoPlaceStoreCreate 商户门店创建接口
// taobao.place.store.create
//
// 用于商家创建线下门店
func TaobaoPlaceStoreCreate(ctx context.Context, clt *core.SDKClient, req *alsc.TaobaoPlaceStoreCreateAPIRequest, resp *alsc.TaobaoPlaceStoreCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
