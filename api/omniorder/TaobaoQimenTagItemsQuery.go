package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoQimenTagItemsQuery 打标结果查询-标维度
// taobao.qimen.tag.items.query
//
// 调用该接口，查询打了某个标的商品列表。说明：该接口调用后，返回值的时间较长，建议不要经常调用。
func TaobaoQimenTagItemsQuery(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoQimenTagItemsQueryAPIRequest, resp *omniorder.TaobaoQimenTagItemsQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
