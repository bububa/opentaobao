package store

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/store"
)

// TaobaoPlaceStoreUpdateLabel 商户门店标签更新接口
// taobao.place.store.update.label
//
// 更新商户门店标签（服务、权益、标签）接口
func TaobaoPlaceStoreUpdateLabel(ctx context.Context, clt *core.SDKClient, req *store.TaobaoPlaceStoreUpdateLabelAPIRequest, resp *store.TaobaoPlaceStoreUpdateLabelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
