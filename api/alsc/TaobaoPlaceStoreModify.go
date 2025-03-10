package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// TaobaoPlaceStoreModify 商家修改线下门店
// taobao.place.store.modify
//
// 用于商家修改线下门店信息
func TaobaoPlaceStoreModify(ctx context.Context, clt *core.SDKClient, req *alsc.TaobaoPlaceStoreModifyAPIRequest, resp *alsc.TaobaoPlaceStoreModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
