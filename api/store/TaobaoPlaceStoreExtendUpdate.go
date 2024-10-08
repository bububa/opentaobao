package store

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/store"
)

// TaobaoPlaceStoreExtendUpdate 商户门店拓展信息更新接口
// taobao.place.store.extend.update
//
// 更新商户门店拓展信息（tags、attribute、bizAtrribute）更新接口
func TaobaoPlaceStoreExtendUpdate(ctx context.Context, clt *core.SDKClient, req *store.TaobaoPlaceStoreExtendUpdateAPIRequest, resp *store.TaobaoPlaceStoreExtendUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
