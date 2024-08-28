package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoQimenItemsMarking 商品通自动打标
// taobao.qimen.items.marking
//
// 调用该接口，对商品进行XXXX标的打标、去标的动作。
func TaobaoQimenItemsMarking(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoQimenItemsMarkingAPIRequest, resp *omniorder.TaobaoQimenItemsMarkingAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
